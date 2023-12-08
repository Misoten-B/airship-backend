package arassets

import (
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/id"
	"gorm.io/gorm"
)

type GormARAssetsRepository struct {
	db *gorm.DB
}

func NewGormARAssetsRepository(db *gorm.DB) *GormARAssetsRepository {
	return &GormARAssetsRepository{
		db: db,
	}
}

// Save はARAssetsの永続化を行ないます。
// 現状、トランザクションをこの層で行っています。トランザクションをどの層で行なうべきかは要検討
// https://sano11o1.com/posts/handle-transaction-in-usecase-layer
func (r *GormARAssetsRepository) Save(arassets arassets.ARAssets) error {
	speakingAsset := arassets.SpeakingAsset()
	qrCodeImage := arassets.QRCodeImage()

	id := arassets.ID().String()
	userID := arassets.UserID().String()
	threedimentionalmodelID := arassets.ThreeDimentionalModelID().String()

	speakingAssetModel := model.SpeakingAsset{
		ID:          id,
		UserID:      userID,
		Description: speakingAsset.Description(),
		AudioPath:   speakingAsset.AudioPath(),
	}

	arAssetModel := model.ARAsset{
		ID:                      id,
		UserID:                  userID,
		SpeakingAssetID:         speakingAssetModel.ID,
		ThreeDimentionalModelID: threedimentionalmodelID,
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

func (r *GormARAssetsRepository) FetchByID(id id.ID) (arassets.ReadModel, error) {
	var arAssetModel model.ARAsset
	var readModel arassets.ReadModel

	// FIXME: gorm 取得最適化
	if err := r.db.Model(&arAssetModel).
		Where("id = ?", id).
		Preload("SpeakingAsset").
		Preload("ThreeDimentionalModel").
		First(&arAssetModel).Error; err != nil {
		return readModel, err
	}

	isCreated := arAssetModel.Status == model.GormStatusCompleted
	return arassets.NewReadModel(
		arAssetModel.ID,
		arAssetModel.UserID,
		arAssetModel.SpeakingAsset.Description,
		arAssetModel.SpeakingAsset.AudioPath,
		arAssetModel.ThreeDimentionalModel.ModelPath,
		arAssetModel.QRCodeImagePath,
		isCreated,
	), nil
}
