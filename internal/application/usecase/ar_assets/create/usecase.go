package usecase

import (
	"mime/multipart"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ARAssetsUsecase interface {
	Create(input ARAssetsCreateInput) (ARAssetsCreateOutput, error)
}

type ARAssetsUsecaseImpl struct {
	arAssetsRepository           service.ARAssetsRepository
	qrCodeImageStorage           service.QRCodeImageStorage
	voiceModelAdapter            voiceservice.VoiceModelAdapter
	voiceService                 voiceservice.VoiceService
	threeDimentionalModelService threeservice.ThreeDimentionalModelService
}

func NewARAssetsUsecaseImpl(
	arAssetsRepository service.ARAssetsRepository,
	qrCodeImageStorage service.QRCodeImageStorage,
	voiceModelAdapter voiceservice.VoiceModelAdapter,
	voiceService voiceservice.VoiceService,
	threeservice threeservice.ThreeDimentionalModelService,
) *ARAssetsUsecaseImpl {
	return &ARAssetsUsecaseImpl{
		arAssetsRepository:           arAssetsRepository,
		qrCodeImageStorage:           qrCodeImageStorage,
		voiceModelAdapter:            voiceModelAdapter,
		voiceService:                 voiceService,
		threeDimentionalModelService: threeservice,
	}
}

type ARAssetsCreateInput struct {
	UID                 string
	SpeakingDescription string
	ThreeDimentionalID  string
	File                multipart.File
	FileHeader          *multipart.FileHeader
}

type ARAssetsCreateOutput struct {
	ID string
}

func (u *ARAssetsUsecaseImpl) Create(input ARAssetsCreateInput) (ARAssetsCreateOutput, error) {
	var output ARAssetsCreateOutput

	// バリデーション & オブジェクト生成
	threedimentionalmodelID := id.ReconstructID(input.ThreeDimentionalID)
	uid := id.ReconstructID(input.UID)

	file := file.NewMyFile(input.File, input.FileHeader)
	qrCodeImage, err := arassets.NewQRCodeImage(file)
	if err != nil {
		return output, customerror.NewApplicationErrorWithoutDetails(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	speakingAsset, err := arassets.NewSpeakingAsset(uid, input.SpeakingDescription)
	if err != nil {
		return output, customerror.NewApplicationErrorWithoutDetails(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	arAssets := arassets.NewARAssets(
		speakingAsset,
		qrCodeImage,
		threedimentionalmodelID,
	)

	// 音声モデルの生成が完了しているかどうか
	isCompleted, err := u.voiceService.IsModelGenerated(uid)
	if err != nil {
		msg := "failed to check if voice model generation is complete"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	if !isCompleted {
		return output, customerror.NewApplicationErrorWithoutDetails(
			"voice model generation has not been completed",
			http.StatusBadRequest,
		)
	}

	// 3Dモデルが存在するかつ、ユーザーが所有しているか
	hasPermission, err := u.threeDimentionalModelService.HasUsePermission(threedimentionalmodelID, uid)
	if err != nil {
		msg := "failed to check if user has permission to use this 3D model"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	if !hasPermission {
		return output, customerror.NewApplicationErrorWithoutDetails(
			"user does not have permission to use this 3D model",
			http.StatusForbidden,
		)
	}

	// AIへ音声ファイル生成を依頼
	request := voiceservice.GenerateAudioFileRequest{
		UID:            arAssets.UserID().String(),
		OutputFilePath: speakingAsset.AudioPath(),
		Language:       "ja",
		Content:        speakingAsset.Description(),
	}

	err = u.voiceModelAdapter.GenerateAudioFile(request)
	if err != nil {
		msg := "failed to generate audio file"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// QRコードアイコン画像の保存
	err = u.qrCodeImageStorage.Save(qrCodeImage)
	if err != nil {
		msg := "failed to save QR code image"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// データベース保存
	err = u.arAssetsRepository.Save(arAssets)
	if err != nil {
		msg := "failed to save AR assets"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return ARAssetsCreateOutput{
		ID: arAssets.ID().String(),
	}, nil
}
