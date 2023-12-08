package fetchbyidpublic

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type FetchByPublicIDUsecase interface {
	Execute(input FetchByPublicIDInput) (FetchByPublicIDOutput, error)
}

type FetchByPublicIDInput struct {
	ID string
}

type FetchByPublicIDOutput struct {
	ID                   string
	SpeakingDescription  string
	SpeakingAudioPath    string
	ThreeDimentionalPath string
}

type FetchByPublicIDInteractor struct {
	arAssetsRepository           service.ARAssetsRepository
	speakingAudioStorage         vservice.SpeakingAudioStorage
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage
}

func NewFetchByPublicIDInteractor(
	arAssetsRepository service.ARAssetsRepository,
	speakingAudioStorage vservice.SpeakingAudioStorage,
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage,
) *FetchByPublicIDInteractor {
	return &FetchByPublicIDInteractor{
		arAssetsRepository:           arAssetsRepository,
		speakingAudioStorage:         speakingAudioStorage,
		threeDimentionalModelStorage: threeDimentionalModelStorage,
	}
}

func (i *FetchByPublicIDInteractor) Execute(
	input FetchByPublicIDInput,
) (
	FetchByPublicIDOutput,
	error,
) {
	var output FetchByPublicIDOutput

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

	if !model.IsCreated() {
		return output, customerror.NewApplicationErrorWithoutDetails(
			"AR assets has not been created",
			http.StatusBadRequest,
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

	return FetchByPublicIDOutput{
		ID:                   model.ID(),
		SpeakingDescription:  model.SpeakingDescription(),
		SpeakingAudioPath:    speakingAudioPath,
		ThreeDimentionalPath: threeDimentionalPath,
	}, nil
}
