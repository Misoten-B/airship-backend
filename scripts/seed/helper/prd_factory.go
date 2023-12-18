//nolint:gomnd // factoryだから許して
package helper

import (
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
)

type AppModelPRD struct {
	User                              *model.User
	TempThreeDimentionalModelTemplate *model.ThreeDimentionalModelTemplate
	PersonalThreeDimentionalModel     *model.PersonalThreeDimentionalModel
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	SpeakingAsset                     *model.SpeakingAsset
	ARAsset                           *model.ARAsset
	BusinessCardBackgroundTemplate    []*model.BusinessCardBackgroundTemplate
	PersonalBusinessCardBackground    *model.PersonalBusinessCardBackground
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       *model.BusinessCardPartsCoordinate
	BusinessCard                      *model.BusinessCard
}

func NewAppModelPRD() *AppModelPRD {
	user := newUserPRD()

	threeDimentionalModels := newThreeDimentionalModelsPRD()
	threeDimentionalModelTemplate := newThreeDimentionalModelTemplatePRD(threeDimentionalModels[0].ID)
	personalThreeDimentionalModel := newPersonalThreeDimentionalModelPRD()

	speakingAsset := newSpeakingAssetPRD()
	arAsset := newARAssetPRD()

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
	personalBusinessCardBackground := &model.PersonalBusinessCardBackground{}

	businessCardPartsCoordinate := newBusinessCardPartsCoordinatePRD()

	businessCard := newBusinessCardPRD()

	return &AppModelPRD{
		User:                              user,
		TempThreeDimentionalModelTemplate: threeDimentionalModelTemplate,
		PersonalThreeDimentionalModel:     personalThreeDimentionalModel,
		ThreeDimentionalModels:            threeDimentionalModels,
		SpeakingAsset:                     speakingAsset,
		ARAsset:                           arAsset,
		BusinessCardBackgroundTemplate:    businessCardBackgroundTemplate,
		PersonalBusinessCardBackground:    personalBusinessCardBackground,
		BusinessCardBackgrounds:           businessCardBackgrounds,
		BusinessCardPartsCoordinate:       businessCardPartsCoordinate,
		BusinessCard:                      businessCard,
	}
}

func newUserPRD() *model.User {
	return &model.User{}
}

func newThreeDimentionalModelTemplatePRD(id string) *model.ThreeDimentionalModelTemplate {
	return &model.ThreeDimentionalModelTemplate{
		ID: id,
	}
}

func newPersonalThreeDimentionalModelPRD() *model.PersonalThreeDimentionalModel {
	return &model.PersonalThreeDimentionalModel{}
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
			ModelPath: "dog.glb",
		},
		{
			ID:        newID(),
			ModelPath: "dog.glb",
		},
	}
}

func newSpeakingAssetPRD() *model.SpeakingAsset {
	return &model.SpeakingAsset{}
}

func newARAssetPRD() *model.ARAsset {
	return &model.ARAsset{}
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

func newBusinessCardPRD() *model.BusinessCard {
	return &model.BusinessCard{}
}
