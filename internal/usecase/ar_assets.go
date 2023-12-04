package usecase

import (
	"mime/multipart"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
)

type ARAssetsUsecase interface {
	Create(input ARAssetsCreateInput) (ARAssetsCreateOutput, error)
}

type ARAssetsUsecaseImpl struct {
	arAssetsRepository service.ARAssetsRepository
	qrCodeImageStorage service.QRCodeImageStorage
	voiceModelAdapter  voiceservice.VoiceModelAdapter
}

func NewARAssetsUsecaseImpl(
	arAssetsRepository service.ARAssetsRepository,
	qrCodeImageStorage service.QRCodeImageStorage,
	voiceModelAdapter voiceservice.VoiceModelAdapter,
) *ARAssetsUsecaseImpl {
	return &ARAssetsUsecaseImpl{
		arAssetsRepository: arAssetsRepository,
		qrCodeImageStorage: qrCodeImageStorage,
		voiceModelAdapter:  voiceModelAdapter,
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
	ID                   string
	SpeakingDescription  string
	SpeakingAudioPath    string
	ThreeDimentionalPath string
	QrcodeIconImagePath  string
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

	blobName := qrCodeImage.Name()

	// 3Dモデルが存在するかつ、ユーザーが所有しているか

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
	err = u.qrCodeImageStorage.Save(blobName, qrCodeImage.File())
	if err != nil {
		return output, err
	}

	// QRコードアイコン画像のURL取得
	qrCodeImagePath, err := u.qrCodeImageStorage.GetImageURL(blobName)
	if err != nil {
		return output, err
	}

	// データベース保存
	err = u.arAssetsRepository.Save(arAssets)
	if err != nil {
		return output, err
	}

	return ARAssetsCreateOutput{
		ID:                   arAssets.ID().String(),
		SpeakingDescription:  speakingAsset.Description(),
		ThreeDimentionalPath: "https://example.com",
		SpeakingAudioPath:    "",
		QrcodeIconImagePath:  qrCodeImagePath,
	}, nil
}
