package usecase

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/database/model"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ARAssetsUsecase interface {
	Create(input ARAssetsCreateInput) (ARAssetsCreateOutput, error)
}

type ARAssetsUsecaseImpl struct {
	qrCodeImageStorage service.QRCodeImageStorage
	voiceModelAdapter  voiceservice.VoiceModelAdapter
}

func NewARAssetsUsecaseImpl(
	qrCodeImageStorage service.QRCodeImageStorage,
	voiceModelAdapter voiceservice.VoiceModelAdapter,
) *ARAssetsUsecaseImpl {
	return &ARAssetsUsecaseImpl{
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

	// バリデーション
	ext := filepath.Ext(input.FileHeader.Filename)
	blobID, err := id.NewID()
	if err != nil {
		return output, err
	}
	blobName := fmt.Sprintf("%s%s", blobID.String(), ext)

	speekingAssetID, err := id.NewID()
	if err != nil {
		return output, err
	}
	audioPath := fmt.Sprintf("%s.wav", speekingAssetID.String())

	// 3Dモデルが存在するかつ、ユーザーが所有しているか

	// AIへ音声ファイル生成を依頼
	request := voiceservice.GenerateAudioFileRequest{
		UID:            input.UID,
		OutputFilePath: audioPath,
		Language:       "ja",
		Content:        input.SpeakingDescription,
	}
	err = u.voiceModelAdapter.GenerateAudioFile(request)
	if err != nil {
		return output, err
	}

	// QRコードアイコン画像保存
	err = u.qrCodeImageStorage.Save(blobName, input.File)
	if err != nil {
		return output, err
	}

	// データベース保存
	speekingAssetModel := model.SpeakingAsset{
		ID:          speekingAssetID.String(),
		UserID:      input.UID,
		Description: input.SpeakingDescription,
		AudioPath:   audioPath,
	}

	arAssetModel := model.ARAsset{
		ID:                      speekingAssetModel.ID,
		UserID:                  input.UID,
		SpeakingAssetID:         speekingAssetModel.ID,
		ThreeDimentionalModelID: input.ThreeDimentionalID,
		QRCodeImagePath:         blobName,
		AccessCount:             0,
	}

	db, err := database.ConnectDB()
	if err != nil {
		return output, err
	}

	tx := db.Begin()

	if err = tx.Create(&speekingAssetModel).Error; err != nil {
		return output, err
	}

	if err = tx.Create(&arAssetModel).Error; err != nil {
		return output, err
	}

	err = tx.Commit().Error
	if err != nil {
		return output, err
	}

	return ARAssetsCreateOutput{
		ID:                   arAssetModel.ID,
		SpeakingDescription:  speekingAssetModel.Description,
		ThreeDimentionalPath: "https://example.com",
		SpeakingAudioPath:    "https://example.com",
		QrcodeIconImagePath:  "https://example.com",
	}, nil
}
