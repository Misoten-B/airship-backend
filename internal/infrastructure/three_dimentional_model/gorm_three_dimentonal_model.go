package threedimentionalmodel

import (
	"errors"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
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

// Save はThreeDimentionalModelの永続化を行ないます。
// GormARAssetsRepositoryと同様のトランザクションの問題があります。
func (r *GormThreeDimentionalModelRepository) Save(
	threeDimentionalModel threedimentionalmodel.ThreeDimensionalModel,
) error {
	id := threeDimentionalModel.ID().String()
	userID := threeDimentionalModel.UserID().String()

	// モデル生成
	tdmModel := model.ThreeDimentionalModel{
		ID:        id,
		ModelPath: threeDimentionalModel.Path().Value(),
	}
	ptdmModel := model.PersonalThreeDimentionalModel{
		ID:     tdmModel.ID,
		UserID: userID,
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

func (r *GormThreeDimentionalModelRepository) Find(id idlib.ID) (*threedimentionalmodel.ThreeDimensionalModel, error) {
	var threeDimentionalModel model.ThreeDimentionalModel

	// FIXME: Gorm取得の最適化
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

	path := shared.NewFilePath(threeDimentionalModel.ModelPath)
	if templateLen != 0 {
		return threedimentionalmodel.ReconstructThreeDimensionalModelTemplate(id, path), nil
	}

	uid := idlib.ReconstructID(threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID)
	return threedimentionalmodel.ReconstructThreeDimensionalModel(id, uid, path), nil
}

// GormARAssetsRepositoryと同様のトランザクションの問題があります。
func (r *GormThreeDimentionalModelRepository) Remove(id idlib.ID) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&model.PersonalThreeDimentionalModel{}, "id = ?", id.String()).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&model.ThreeDimentionalModel{}, "id = ?", id.String()).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *GormThreeDimentionalModelRepository) FindByID(id idlib.ID) (threedimentionalmodel.ReadModel, error) {
	var threeDimentionalModel model.ThreeDimentionalModel

	// TODO: 重複部分の修正または統合
	// FIXME: Gorm取得の最適化
	if err := r.db.Preload("PersonalThreeDimentionalModels").
		Preload("ThreeDimentionalModelTemplates").
		First(&threeDimentionalModel, id).Error; err != nil {
		return threedimentionalmodel.ReadModel{}, err
	}

	templateLen := len(threeDimentionalModel.ThreeDimentionalModelTemplates)
	personalLen := len(threeDimentionalModel.PersonalThreeDimentionalModels)

	if templateLen == 0 && personalLen == 0 {
		return threedimentionalmodel.ReadModel{}, errors.New("three dimentional model not found")
	}

	if templateLen != 0 {
		return threedimentionalmodel.NewTemplateReadModel(
			id.String(),
			threeDimentionalModel.ModelPath,
		), nil
	}

	if personalLen != 0 {
		uid := idlib.ReconstructID(threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID)
		return threedimentionalmodel.NewReadModel(
			id.String(),
			uid.String(),
			threeDimentionalModel.ModelPath,
		), nil
	}
	return threedimentionalmodel.ReadModel{}, errors.New("three dimentional model not found")
}

func (r *GormThreeDimentionalModelRepository) FindByUserID(
	userID idlib.ID,
) ([]threedimentionalmodel.ReadModel, error) {
	var threeDimentionalModels []model.ThreeDimentionalModel

	// FIXME: Gorm取得の最適化
	// すべてのThreeDimentionalModelTemplatesとuser_idが一致するPersonalThreeDimentionalModelsを取得
	if err := r.db.Preload("PersonalThreeDimentionalModels", "user_id = ?", userID).
		Preload("ThreeDimentionalModelTemplates").
		Find(&threeDimentionalModels).Error; err != nil {
		return nil, err
	}

	var readModels []threedimentionalmodel.ReadModel

	for _, threeDimentionalModel := range threeDimentionalModels {
		templateLen := len(threeDimentionalModel.ThreeDimentionalModelTemplates)
		personalLen := len(threeDimentionalModel.PersonalThreeDimentionalModels)

		if templateLen == 0 && personalLen == 0 {
			return nil, errors.New("three dimentional model not found")
		}

		if templateLen != 0 {
			readModels = append(readModels, threedimentionalmodel.NewTemplateReadModel(
				threeDimentionalModel.ID,
				threeDimentionalModel.ModelPath,
			))
		}

		if personalLen != 0 {
			readModels = append(readModels, threedimentionalmodel.NewReadModel(
				threeDimentionalModel.ID,
				threeDimentionalModel.PersonalThreeDimentionalModels[0].UserID,
				threeDimentionalModel.ModelPath,
			))
		}
	}

	return readModels, nil
}
