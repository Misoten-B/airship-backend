package helper

import (
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type AppModelPRD struct {
	User                              *model.User
	TempThreeDimentionalModelTemplate *model.ThreeDimentionalModelTemplate
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	BusinessCardBackgroundTemplate    []*model.BusinessCardBackgroundTemplate
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       []*model.BusinessCardPartsCoordinate
}

func NewAppModelPRD() *AppModelPRD {
	threeDimentionalModels := newThreeDimentionalModelsPRD()
	threeDimentionalModelTemplate := newThreeDimentionalModelTemplatePRD(threeDimentionalModels[0].ID)

	businessCardBackgrounds := []*model.BusinessCardBackground{
		{
			ID:        newID(),
			ColorCode: "#ffffff",
			ImagePath: "background-template-1.png",
		},
		{
			ID:        newID(),
			ColorCode: "#ffffff",
			ImagePath: "background-template-2.png",
		},
		{
			ID:        newID(),
			ColorCode: "#ffffff",
			ImagePath: "background-template-3.png",
		},
	}

	businessCardBackgroundTemplate := []*model.BusinessCardBackgroundTemplate{
		{ID: businessCardBackgrounds[0].ID},
		{ID: businessCardBackgrounds[1].ID},
		{ID: businessCardBackgrounds[2].ID},
	}

	businessCardPartsCoordinate := newBusinessCardPartsCoordinate()

	return &AppModelPRD{
		User:                              newUser(),
		TempThreeDimentionalModelTemplate: threeDimentionalModelTemplate,
		ThreeDimentionalModels:            threeDimentionalModels,
		BusinessCardBackgroundTemplate:    businessCardBackgroundTemplate,
		BusinessCardBackgrounds:           businessCardBackgrounds,
		BusinessCardPartsCoordinate:       businessCardPartsCoordinate,
	}
}

func newUser() *model.User {
	return &model.User{
		ID:                testdata.DEV_UID,
		RecordedModelPath: fmt.Sprintf("%s.npz", testdata.DEV_UID),
		IsToured:          true,
		Status:            model.GormStatusCompleted,
	}
}

func newThreeDimentionalModelTemplatePRD(id string) *model.ThreeDimentionalModelTemplate {
	return &model.ThreeDimentionalModelTemplate{
		ID: id,
	}
}

func newThreeDimentionalModelsPRD() []*model.ThreeDimentionalModel {
	return []*model.ThreeDimentionalModel{
		{
			ID:        newID(),
			ModelPath: "chicken.glb",
		},
		{
			ID:        newID(),
			ModelPath: "dog.glb",
		},
		{
			ID:        newID(),
			ModelPath: "pinguin.glb",
		},
		{
			ID:        newID(),
			ModelPath: "tiger.glb",
		},
	}
}
