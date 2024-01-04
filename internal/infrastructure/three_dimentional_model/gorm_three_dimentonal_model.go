package threedimentionalmodel

import (
	"errors"

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
		// ID:     tdmModel.ID,
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
	var threeDimentionalModel model.ThreeDimentionalModel

	// TODO: 応急処置を治す
	// FIXME: Gorm取得の最適化
	if err := r.db.Preload("PersonalThreeDimentionalModels").
		Preload("ThreeDimentionalModelTemplates").
		First(&threeDimentionalModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrThreeDimentionalModelNotFound
		}
		return nil, err
	}

	// templateLen := len(threeDimentionalModel.ThreeDimentionalModelTemplates)
	// personalLen := len(threeDimentionalModel.PersonalThreeDimentionalModels)

	// if templateLen == 0 && personalLen == 0 {
	// return nil, errors.New("three dimentional model not found")
	// }

	// if templateLen != 0 {
	// return threedimentionalmodel.ReconstructThreeDimentionalModelTemplate(id), nil
	// }

	uid := shared.ReconstructID("threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID")
	return threedimentionalmodel.ReconstructThreeDimentionalModel(id, uid), nil
}

func (r *GormThreeDimentionalModelRepository) FindByID(id shared.ID) (threedimentionalmodel.ReadModel, error) {
	var threeDimentionalModel model.ThreeDimentionalModel

	// TODO: 応急処置を治す
	// TODO: 重複部分の修正または統合
	// FIXME: Gorm取得の最適化
	if err := r.db.Preload("PersonalThreeDimentionalModels").
		Preload("ThreeDimentionalModelTemplates").
		First(&threeDimentionalModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return threedimentionalmodel.ReadModel{}, service.ErrThreeDimentionalModelNotFound
		}
		return threedimentionalmodel.ReadModel{}, err
	}

	// templateLen := len(threeDimentionalModel.ThreeDimentionalModelTemplates)
	// personalLen := len(threeDimentionalModel.PersonalThreeDimentionalModels)

	// if templateLen == 0 && personalLen == 0 {
	// return threedimentionalmodel.ReadModel{}, errors.New("three dimentional model not found")
	// }

	// if templateLen != 0 {
	return threedimentionalmodel.NewTemplateReadModel(
		id.String(),
		threeDimentionalModel.ModelPath,
	), nil
	// }

	// if personalLen != 0 {
	// uid := shared.ReconstructID(threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID)
	// return threedimentionalmodel.NewReadModel(
	// id.String(),
	// uid.String(),
	// threeDimentionalModel.ModelPath,
	// ), nil
	// }
	// return threedimentionalmodel.ReadModel{}, errors.New("three dimentional model not found")
}

func (r *GormThreeDimentionalModelRepository) FindByUserID(
	userID shared.ID,
) ([]threedimentionalmodel.ReadModel, error) {
	var personalTDMs []model.PersonalThreeDimentionalModel
	var tDMTemplates []model.ThreeDimentionalModelTemplate

	r.db.Debug().Preload("ThreeDimentionalModel").Where("user_id = ?", userID).Find(&personalTDMs)
	r.db.Preload("ThreeDimentionalModel").Find(&tDMTemplates)

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
