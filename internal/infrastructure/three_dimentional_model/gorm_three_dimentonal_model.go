package threedimentionalmodel

import (
	"errors"
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"gorm.io/gorm"
)

type GormThreeDimentionalModelRepository struct {
	db *gorm.DB
}

func NewGormThreeDimentionalModelRepository(db *gorm.DB) *GormThreeDimentionalModelRepository {
	return &GormThreeDimentionalModelRepository{
		db: db,
	}
}

// Save はThreeDimentionalModelの永続化を行ないます。
// GormARAssetsRepositoryと同様のトランザクションの問題があります。
func (r *GormThreeDimentionalModelRepository) Save(
	threeDimentionalModel threedimentionalmodel.ThreeDimentionalModel,
) error {
	id := threeDimentionalModel.ID().String()
	userID := threeDimentionalModel.UserID().String()

	// モデル生成
	tdmModel := model.ThreeDimentionalModel{
		ID:        id,
		ModelPath: threeDimentionalModel.FileName(),
	}
	ptdmModel := model.PersonalThreeDimentionalModel{
		ThreeDimentionalModel: tdmModel,
		UserID:                userID,
	}

	// トランザクション
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&tdmModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&ptdmModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *GormThreeDimentionalModelRepository) Find(id shared.ID) (*threedimentionalmodel.ThreeDimentionalModel, error) {
	var personalTDM model.PersonalThreeDimentionalModel

	err := r.db.
		Preload("ThreeDimentionalModel").
		First(&personalTDM, id).
		Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to fetch Personal Three Dimentional Model: %w", err)
		}
	} else {
		return threedimentionalmodel.ReconstructThreeDimentionalModel(id, shared.ID(personalTDM.UserID)), nil
	}

	var tDMTemplate model.ThreeDimentionalModelTemplate

	err = r.db.
		Preload("ThreeDimentionalModel").
		First(&tDMTemplate, id).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrThreeDimentionalModelNotFound
		}
		return nil, fmt.Errorf("failed to fetch Three Dimentional Model Template: %w", err)
	}
	return threedimentionalmodel.ReconstructThreeDimentionalModelTemplate(id), nil
}

func (r *GormThreeDimentionalModelRepository) FindByID(id shared.ID) (threedimentionalmodel.ReadModel, error) {
	var personalTDM model.PersonalThreeDimentionalModel

	err := r.db.
		Preload("ThreeDimentionalModel").
		First(&personalTDM, id).
		Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return threedimentionalmodel.ReadModel{}, fmt.Errorf("failed to fetch Personal Three Dimentional Model: %w", err)
		}
	} else {
		return threedimentionalmodel.NewReadModel(
			id.String(),
			personalTDM.UserID,
			personalTDM.ThreeDimentionalModel.ModelPath,
		), nil
	}

	var tDMTemplate model.ThreeDimentionalModelTemplate

	err = r.db.
		Preload("ThreeDimentionalModel").
		First(&tDMTemplate, id).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return threedimentionalmodel.ReadModel{}, service.ErrThreeDimentionalModelNotFound
		}
		return threedimentionalmodel.ReadModel{}, fmt.Errorf("failed to fetch Three Dimentional Model Template: %w", err)
	}
	return threedimentionalmodel.NewTemplateReadModel(id.String(), tDMTemplate.ThreeDimentionalModel.ModelPath), nil
}

func (r *GormThreeDimentionalModelRepository) FindByUserID(
	userID shared.ID,
) ([]threedimentionalmodel.ReadModel, error) {
	var personalTDMs []model.PersonalThreeDimentionalModel
	var tDMTemplates []model.ThreeDimentionalModelTemplate

	if err := r.db.
		Preload("ThreeDimentionalModel").
		Where("user_id = ?", userID).
		Find(&personalTDMs).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to fetch Personal Three Dimentional Models: %w", err)
		}
	}

	if err := r.db.
		Preload("ThreeDimentionalModel").
		Find(&tDMTemplates).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to fetch Three Dimentional Model Templates: %w", err)
		}
	}

	var readModels []threedimentionalmodel.ReadModel

	for _, tdm := range personalTDMs {
		readModels = append(readModels, threedimentionalmodel.NewReadModel(
			tdm.ThreeDimentionalModel.ID,
			tdm.UserID,
			tdm.ThreeDimentionalModel.ModelPath,
		))
	}

	for _, tdm := range tDMTemplates {
		readModels = append(readModels, threedimentionalmodel.NewTemplateReadModel(
			tdm.ThreeDimentionalModel.ID,
			tdm.ThreeDimentionalModel.ModelPath,
		))
	}

	return readModels, nil
}
