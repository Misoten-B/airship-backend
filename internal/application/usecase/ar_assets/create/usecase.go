package usecase

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/file"
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
	threedimentionalmodelID := shared.ReconstructID(input.ThreeDimentionalID)
	uid := shared.ReconstructID(input.UID)

	var qrCodeImage arassets.QRCodeImage
	if input.File != nil {
		var err error

		file := file.NewMyFile(input.File, input.FileHeader)
		qrCodeImage, err = arassets.NewQRCodeImage(file)
		if err != nil {
			return output, customerror.NewApplicationErrorWithoutDetails(
				err.Error(),
				http.StatusBadRequest,
			)
		}
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
	err = u.hasPermission(threedimentionalmodelID, uid)
	if err != nil {
		return output, err
	}

	// AIへ音声ファイル生成を依頼
	request := voiceservice.GenerateAudioFileRequest{
		UID:        arAssets.UserID().String(),
		ARAssetsID: arAssets.ID().String(),
		Language:   "ja",
		Content:    speakingAsset.Description(),
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
	if input.File != nil {
		err = u.saveStorage(qrCodeImage)
		if err != nil {
			return output, err
		}
	}

	// データベース保存
	err = u.saveModel(arAssets)
	if err != nil {
		return output, err
	}

	return ARAssetsCreateOutput{
		ID: arAssets.ID().String(),
	}, nil
}

func (u *ARAssetsUsecaseImpl) hasPermission(threedimentionalmodelID shared.ID, userID shared.ID) error {
	hasPermission, err := u.threeDimentionalModelService.HasUsePermission(threedimentionalmodelID, userID)
	if err != nil {
		if errors.Is(err, threeservice.ErrThreeDimentionalModelNotFound) {
			return customerror.NewApplicationErrorWithoutDetails(
				"3D model not found",
				http.StatusNotFound,
			)
		}
		msg := "failed to check if user has permission to use this 3D model"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	if !hasPermission {
		return customerror.NewApplicationErrorWithoutDetails(
			"user does not have permission to use this 3D model",
			http.StatusForbidden,
		)
	}
	return nil
}

func (u *ARAssetsUsecaseImpl) saveStorage(qrCodeImage arassets.QRCodeImage) error {
	err := u.qrCodeImageStorage.Save(qrCodeImage)

	if err != nil {
		msg := "failed to save QR code image"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	return nil
}

func (u *ARAssetsUsecaseImpl) saveModel(arAssets arassets.ARAssets) error {
	err := u.arAssetsRepository.Save(arAssets)
	if err != nil {
		msg := "failed to save AR assets"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	return nil
}
