package usecase

import (
	"errors"
	"mime/multipart"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
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
	qrCodeImage, err := arassets.NewQRCodeImage(input.File, input.FileHeader)
	if err != nil {
		return output, err
	}
	speakingAsset, err := arassets.NewSpeakingAsset(input.UID, input.SpeakingDescription)
	if err != nil {
		return output, err
	}
	arAssets := arassets.NewARAssets(
		speakingAsset,
		qrCodeImage,
		input.ThreeDimentionalID,
	)

	threedimentionalmodelID := id.ReconstructID(input.ThreeDimentionalID)
	uid := id.ReconstructID(input.UID)

	// 音声モデルの生成が完了しているかどうか
	isCompleted, err := u.voiceService.IsModelGenerated(uid)
	if err != nil {
		return output, err
	}
	if !isCompleted {
		return output, errors.New("voice model generation has not been completed")
	}

	// 3Dモデルが存在するかつ、ユーザーが所有しているか
	hasPermission, err := u.threeDimentionalModelService.HasUsePermission(threedimentionalmodelID, uid)

	if err != nil {
		return output, err
	}
	if !hasPermission {
		return output, errors.New("user does not have permission to use this 3D model")
	}

	// AIへ音声ファイル生成を依頼
	request := voiceservice.GenerateAudioFileRequest{
		UID:            arAssets.UserID(),
		OutputFilePath: speakingAsset.AudioPath(),
		Language:       "ja",
		Content:        speakingAsset.Description(),
	}
	err = u.voiceModelAdapter.GenerateAudioFile(request)
	if err != nil {
		return output, err
	}

	// QRコードアイコン画像の保存
	err = u.qrCodeImageStorage.Save(qrCodeImage)
	if err != nil {
		return output, err
	}

	// データベース保存
	err = u.arAssetsRepository.Save(arAssets)
	if err != nil {
		return output, err
	}

	return ARAssetsCreateOutput{
		ID: arAssets.ID().String(),
	}, nil
}
