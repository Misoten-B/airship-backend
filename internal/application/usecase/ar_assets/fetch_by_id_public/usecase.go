package fetchbyidpublic

import (
	"errors"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	ID string
}

type Output struct {
	ID                   string
	SpeakingDescription  string
	SpeakingAudioPath    string
	ThreeDimentionalPath string
	IsCompleted          bool
}

// Interactor はARアセットID検索（Public）ユースケースの実装です。
type Interactor struct {
	arAssetsRepository           service.ARAssetsRepository
	speakingAudioStorage         vservice.SpeakingAudioStorage
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage
}

func NewInteractor(
	arAssetsRepository service.ARAssetsRepository,
	speakingAudioStorage vservice.SpeakingAudioStorage,
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage,
) *Interactor {
	return &Interactor{
		arAssetsRepository:           arAssetsRepository,
		speakingAudioStorage:         speakingAudioStorage,
		threeDimentionalModelStorage: threeDimentionalModelStorage,
	}
}

// Execute はARアセットID検索（Public）ユースケースを実行します。
func (i *Interactor) Execute(
	input Input,
) (
	Output,
	error,
) {
	var output Output

	// バリデーション & オブジェクト生成
	id := shared.ReconstructID(input.ID)

	// リポジトリから取得
	model, err := i.arAssetsRepository.FetchByID(id)
	if err != nil {
		if errors.Is(err, service.ErrArAssetsNotFound) {
			return output, customerror.NewApplicationErrorWithoutDetails(
				"AR assets not found",
				http.StatusNotFound,
			)
		}
		msg := "failed to fetch AR assets"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// アクセスカウントをインクリメント

	// URL生成
	speakingAudioPath, err := i.speakingAudioStorage.GetAudioURL(model.SpeakingAudioPath())
	if err != nil {
		msg := "failed to get speaking audio URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	threeDimentionalPath, err := i.threeDimentionalModelStorage.GetModelURL(model.ThreeDimentionalPath())
	if err != nil {
		msg := "failed to get 3D model URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return Output{
		ID:                   model.ID(),
		SpeakingDescription:  model.SpeakingDescription(),
		SpeakingAudioPath:    speakingAudioPath,
		ThreeDimentionalPath: threeDimentionalPath,
		IsCompleted:          model.IsCreated(),
	}, nil
}
