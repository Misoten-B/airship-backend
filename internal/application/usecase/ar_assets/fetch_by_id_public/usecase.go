package fetchbyidpublic

import (
	"fmt"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/id"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
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
	id := id.ReconstructID(input.ID)

	// リポジトリから取得
	model, err := i.arAssetsRepository.FetchByID(id)
	if err != nil {
		msg := "failed to fetch AR assets"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// アクセスカウントをインクリメント

	// URL生成
	adapter := voice.NewVallEXAdapter()
	relativePath := adapter.GetAudioRelativePath()
	speakingAudioPath := fmt.Sprintf("%s/%s", relativePath, model.SpeakingAudioPath())

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
