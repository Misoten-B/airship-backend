//nolint:revive,nolintlint // ←削除検討
package deletearassets

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"gorm.io/gorm"
)

type Usecase interface {
	Execute(input Input) error
}

type Input struct {
	ID     string
	UserID string
}

// Interactor はARアセット削除ユースケースの実装です。
type Interactor struct {
	db              *gorm.DB
	azureBlobDriver *drivers.AzureBlobDriver
}

func NewInteractor(
	db *gorm.DB,
	azureBlobDriver *drivers.AzureBlobDriver,
) *Interactor {
	return &Interactor{
		db:              db,
		azureBlobDriver: azureBlobDriver,
	}
}

// Execute はARアセット削除ユースケースを実行します。
// FIXME: ベタ書きだけど許してください
func (i *Interactor) Execute(input Input) error {
	// ARアセットの取得
	dbModel, err := i.fetchByID(input.ID)
	if err != nil {
		return err
	}

	// ARアセットの権限確認
	if dbModel.UserID != input.UserID {
		return customerror.NewApplicationErrorWithoutDetails(
			"permission denied",
			http.StatusForbidden,
		)
	}

	// データベースから削除する
	if err = i.deleteFromRepository(dbModel); err != nil {
		return err
	}

	// （存在するなら）ストレージからQRコード画像を削除する
	if dbModel.QRCodeImagePath != "" {
		if err = i.deleteQRCodeImageFromStorage(dbModel.QRCodeImagePath); err != nil {
			return err
		}
	}

	// ストレージから音声ファイルを削除する
	if err = i.deleteSpeakingAudioFromStorage(dbModel.SpeakingAsset.AudioPath); err != nil {
		return err
	}

	return nil
}

// interfaceに実装済みだけど面倒くさいのでここに書く
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

func (i *Interactor) deleteFromRepository(dbModel model.ARAsset) error {
	tx := i.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 参照している名刺の全削除
	if err := tx.Where("ar_asset_id = ?", dbModel.ID).Delete(&model.BusinessCard{}).Error; err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete business cards: %w", err),
			"failed to delete business cards",
			http.StatusInternalServerError,
		)
	}

	// ARアセットの削除
	if err := tx.Delete(&dbModel).Error; err != nil {
		tx.Rollback()

		// 外部キー制約による削除失敗
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return customerror.NewApplicationErrorWithoutDetails(
				"failed to delete ar asset: foreign key violated",
				http.StatusBadRequest,
			)
		}
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete ar asset: %w", err),
			"failed to delete ar asset",
			http.StatusInternalServerError,
		)
	}

	// 音声アセットの削除
	if err := tx.Delete(&dbModel.SpeakingAsset).Error; err != nil {
		tx.Rollback()
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete speaking asset: %w", err),
			"failed to delete speaking asset",
			http.StatusInternalServerError,
		)
	}

	err := tx.Commit().Error
	if err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to commit transaction: %w", err),
			"failed to commit transaction",
			http.StatusInternalServerError,
		)
	}

	return nil
}

func (i *Interactor) deleteQRCodeImageFromStorage(path string) error {
	if err := i.azureBlobDriver.DeleteBlob("qrcode-images", path); err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete qrcode image from storage: %w", err),
			"failed to delete qrcode image from storage",
			http.StatusInternalServerError,
		)
	}

	return nil
}

func (i *Interactor) deleteSpeakingAudioFromStorage(path string) error {
	if err := i.azureBlobDriver.DeleteBlob("voice-sounds", path); err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete speaking audio from storage: %w", err),
			"failed to delete speaking audio from storage",
			http.StatusInternalServerError,
		)
	}

	return nil
}
