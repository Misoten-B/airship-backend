package deleteqrcodeicon

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

// Interactor はQRコードアイコン削除ユースケースの実装です。
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

// Execute はQRコードアイコン削除ユースケースを実行します。
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

	// 値が初期値なら何もしない
	if dbModel.QRCodeImagePath == "" {
		return customerror.NewApplicationErrorWithoutDetails(
			"qr code icon not found",
			http.StatusNotFound,
		)
	}

	// データベースから値を削除する（値を初期値に）
	qrCodeImagePath := dbModel.QRCodeImagePath
	dbModel.QRCodeImagePath = ""
	if err = i.db.Save(&dbModel).Error; err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to update ar asset: %w", err),
			"failed to update ar asset",
			http.StatusInternalServerError,
		)
	}

	// ストレージからQRコード画像を削除する
	if err = i.azureBlobDriver.DeleteBlob("qrcode-images", qrCodeImagePath); err != nil {
		return customerror.NewApplicationError(
			fmt.Errorf("failed to delete qr code image: %w", err),
			"failed to delete qr code image",
			http.StatusInternalServerError,
		)
	}

	return nil
}

// deleteとまったく同じことを書いてる
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
