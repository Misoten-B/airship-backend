//nolint:gomnd // factoryだから許して
package helper

import (
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
)

type AppModelPRD struct {
	TempThreeDimentionalModelTemplate *model.ThreeDimentionalModelTemplate
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	BusinessCardBackgroundTemplate    []*model.BusinessCardBackgroundTemplate
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       *model.BusinessCardPartsCoordinate
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

	businessCardPartsCoordinate := newBusinessCardPartsCoordinatePRD()

	return &AppModelPRD{
		TempThreeDimentionalModelTemplate: threeDimentionalModelTemplate,
		ThreeDimentionalModels:            threeDimentionalModels,
		BusinessCardBackgroundTemplate:    businessCardBackgroundTemplate,
		BusinessCardBackgrounds:           businessCardBackgrounds,
		BusinessCardPartsCoordinate:       businessCardPartsCoordinate,
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

func newBusinessCardPartsCoordinatePRD() *model.BusinessCardPartsCoordinate {
	return &model.BusinessCardPartsCoordinate{
		ID:                newID(),
		DisplayNameX:      112,
		DisplayNameY:      266,
		CompanyNameX:      116,
		CompanyNameY:      98,
		DepartmentX:       116,
		DepartmentY:       152,
		OfficialPositionX: 116,
		OfficialPositionY: 200,
		PhoneNumberX:      116,
		PhoneNumberY:      478,
		EmailX:            116,
		EmailY:            428,
		PostalCodeX:       116,
		PostalCodeY:       574,
		AddressX:          116,
		AddressY:          614,
		QRCodeX:           760,
		QRCodeY:           209,
	}
}
