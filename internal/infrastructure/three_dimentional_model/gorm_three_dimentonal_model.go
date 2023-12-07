package threedimentionalmodel

import (
	"errors"

	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	idlib "github.com/Misoten-B/airship-backend/internal/id"
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

func (r *GormThreeDimentionalModelRepository) Find(id idlib.ID) (*threedimentionalmodel.ThreeDimentionalModel, error) {
	var threeDimentionalModel model.ThreeDimentionalModel

	if err := r.db.Preload("PersonalThreeDimentionalModels").
		Preload("ThreeDimentionalModelTemplates").
		First(&threeDimentionalModel, id).Error; err != nil {
		return nil, err
	}

	templateLen := len(threeDimentionalModel.ThreeDimentionalModelTemplates)
	personalLen := len(threeDimentionalModel.PersonalThreeDimentionalModels)

	if templateLen == 0 && personalLen == 0 {
		return nil, errors.New("three dimentional model not found")
	}

	if templateLen != 0 {
		return threedimentionalmodel.ReconstructThreeDimentionalModelTemplate(id), nil
	}

	uid := idlib.ReconstructID(threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID)
	return threedimentionalmodel.ReconstructThreeDimentionalModel(id, uid), nil
}
