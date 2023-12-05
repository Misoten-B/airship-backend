package arassets

import (
	"github.com/Misoten-B/airship-backend/internal/database/model"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"gorm.io/gorm"
)

type GormARAssetsRepository struct {
	db *gorm.DB
}

func NewGormARAssetsRepository(db *gorm.DB) service.ARAssetsRepository {
	return GormARAssetsRepository{
		db: db,
	}
}

// Save はARAssetsの永続化を行ないます。
// 現状、トランザクションをこの層で行っています。トランザクションをどの層で行なうべきかは要検討
// https://sano11o1.com/posts/handle-transaction-in-usecase-layer
func (r GormARAssetsRepository) Save(arassets arassets.ARAssets) error {
	speakingAsset := arassets.SpeakingAsset()
	qrCodeImage := arassets.QRCodeImage()

	speakingAssetModel := model.SpeakingAsset{
		ID:          speakingAsset.ID().String(),
		UserID:      speakingAsset.UserID().String(),
		Description: speakingAsset.Description(),
		AudioPath:   speakingAsset.AudioPath(),
	}

	arAssetModel := model.ARAsset{
		ID:                      arassets.ID().String(),
		UserID:                  arassets.UserID().String(),
		SpeakingAssetID:         speakingAssetModel.ID,
		ThreeDimentionalModelID: arassets.ThreeDimentionalModelID().String(),
		QRCodeImagePath:         qrCodeImage.Name(),
		AccessCount:             arassets.AccessCount(),
		Status:                  model.GormStatusInProgress,
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&speakingAssetModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&arAssetModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
