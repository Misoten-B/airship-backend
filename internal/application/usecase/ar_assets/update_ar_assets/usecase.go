package updatearassets

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/file"
	"gorm.io/gorm"
)

type Input struct {
	ID                  string
	UserID              string
	ThreeDimentionalID  string
	SpeakingDescription string
	QRCodeImage         *QRCodeImageInput
}

type QRCodeImageInput struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type Usecase interface {
	Execute(input Input) error
}

type Interactor struct {
	db                           *gorm.DB
	arAssetsRepository           service.ARAssetsRepository
	qrCodeImageStorage           service.QRCodeImageStorage
	voiceModelAdapter            voiceservice.VoiceModelAdapter
	voiceService                 voiceservice.VoiceService
	threeDimentionalModelService threeservice.ThreeDimentionalModelService
}

func NewInteractor(
	db *gorm.DB,
	arAssetsRepository service.ARAssetsRepository,
	qrCodeImageStorage service.QRCodeImageStorage,
	voiceModelAdapter voiceservice.VoiceModelAdapter,
	voiceService voiceservice.VoiceService,
	threeservice threeservice.ThreeDimentionalModelService,
) *Interactor {
	return &Interactor{
		db:                           db,
		arAssetsRepository:           arAssetsRepository,
		qrCodeImageStorage:           qrCodeImageStorage,
		voiceModelAdapter:            voiceModelAdapter,
		voiceService:                 voiceService,
		threeDimentionalModelService: threeservice,
	}
}

// Execute はARアセット更新ユースケースを実行します。
//
// QRコードアイコンの削除はここでは行わない。
func (i *Interactor) Execute(input Input) error {
	id := shared.ReconstructID(input.ID)
	userID := shared.ReconstructID(input.UserID)

	// 音声モデルが生成済みか
	// 生成と重複してるけどめんどくさいのでそのまま
	isCompleted, err := i.voiceService.IsModelGenerated(userID)
	if err != nil {
		msg := "failed to check if voice model generation is complete"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}
	if !isCompleted {
		return customerror.NewApplicationErrorWithoutDetails(
			"voice model generation has not been completed",
			http.StatusBadRequest,
		)
	}

	// ARアセットの取得
	prevARAssets, err := i.fetchByID(id.String())
	if err != nil {
		return err
	}

	// 権限確認
	if prevARAssets.UserID != input.UserID {
		return customerror.NewApplicationErrorWithoutDetails(
			"user does not have permission to use this AR assets",
			http.StatusForbidden,
		)
	}

	// バリデーション

	// （変更があれば）3Dモデルの権限確認
	if prevARAssets.ThreeDimentionalModelID != input.ThreeDimentionalID {
		permErr := i.hasPermission(shared.ReconstructID(input.ThreeDimentionalID), userID)

		if permErr != nil {
			return permErr
		}

		prevARAssets.ThreeDimentionalModel.ID = input.ThreeDimentionalID
	}

	// （変更があれば）AIへのリクエスト
	if prevARAssets.SpeakingAsset.Description != input.SpeakingDescription {
		request := voiceservice.GenerateAudioFileRequest{
			UID:        userID.String(),
			ARAssetsID: id.String(),
			Language:   "ja",
			Content:    input.SpeakingDescription,
		}

		if err = i.voiceModelAdapter.GenerateAudioFile(request); err != nil {
			msg := "failed to generate audio file"
			return customerror.NewApplicationError(
				err,
				msg,
				http.StatusInternalServerError,
			)
		}

		prevARAssets.SpeakingAsset.Description = input.SpeakingDescription
	}

	// （変更があれば）ストレージへの保存
	var qrCodeImage arassets.QRCodeImage
	if input.QRCodeImage != nil {
		myFile := file.NewMyFile(input.QRCodeImage.File, input.QRCodeImage.FileHeader)
		qrCodeImage, err = arassets.NewQRCodeImage(myFile)
		if err != nil {
			return customerror.NewApplicationErrorWithoutDetails(
				err.Error(),
				http.StatusBadRequest,
			)
		}

		// もし前のQRコードが存在していたら削除

		// ストレージ保存
		if err = i.qrCodeImageStorage.Save(qrCodeImage); err != nil {
			msg := "failed to save qr code image"
			return customerror.NewApplicationError(
				err,
				msg,
				http.StatusInternalServerError,
			)
		}

		prevARAssets.QRCodeImagePath = qrCodeImage.Name()
	}

	// データベース更新
	if err = i.update(prevARAssets); err != nil {
		return customerror.NewApplicationError(
			err,
			"failed to update ar asset",
			http.StatusInternalServerError,
		)
	}

	return nil
}

// deleteと重複しているけど
func (i *Interactor) fetchByID(id string) (model.ARAsset, error) {
	var arAssetModel model.ARAsset

	if err := i.db.Where("id = ?", id).
		Preload("SpeakingAsset").
		Preload("ThreeDimentionalModel").
		First(&arAssetModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return arAssetModel, customerror.NewApplicationErrorWithoutDetails(
				"ar asset not found",
				http.StatusNotFound,
			)
		}
		return arAssetModel, customerror.NewApplicationError(
			fmt.Errorf("failed to fetch ar asset: %w", err),
			"failed to fetch ar asset",
			http.StatusInternalServerError,
		)
	}

	return arAssetModel, nil
}

// createと重複しているけど
func (i *Interactor) hasPermission(threedimentionalmodelID shared.ID, userID shared.ID) error {
	hasPermission, err := i.threeDimentionalModelService.HasUsePermission(threedimentionalmodelID, userID)
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

func (i *Interactor) update(nextModel model.ARAsset) error {
	tx := i.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&nextModel.SpeakingAsset).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&nextModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
