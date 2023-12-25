//nolint:gomnd // factoryだから許して
package helper

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type AppModel struct {
	User                              *model.User
	TempThreeDimentionalModelTemplate *model.ThreeDimentionalModelTemplate
	PersonalThreeDimentionalModel     *model.PersonalThreeDimentionalModel
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	SpeakingAsset                     *model.SpeakingAsset
	ARAsset                           *model.ARAsset
	BusinessCardBackgroundTemplate    []*model.BusinessCardBackgroundTemplate
	PersonalBusinessCardBackground    *model.PersonalBusinessCardBackground
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       []*model.BusinessCardPartsCoordinate
	BusinessCard                      *model.BusinessCard
}

func NewAppModel() *AppModel {
	user := newUser()

	threeDimentionalModels := newThreeDimentionalModels()
	threeDimentionalModelTemplate := newThreeDimentionalModelTemplate(threeDimentionalModels[0].ID)
	personalThreeDimentionalModel := newPersonalThreeDimentionalModel(threeDimentionalModels[1].ID, user)

	speakingAsset := newSpeakingAsset(user)
	arAsset := newARAsset(user, speakingAsset, threeDimentionalModels[0])

	businessCardBackgrounds := []*model.BusinessCardBackground{
		{
			ID:        newID(),
			ColorCode: "#ffffff",
			ImagePath: "",
		},
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
		{ID: businessCardBackgrounds[2].ID},
		{ID: businessCardBackgrounds[3].ID},
	}
	personalBusinessCardBackground := &model.PersonalBusinessCardBackground{
		ID:     businessCardBackgrounds[1].ID,
		UserID: user.ID,
	}

	businessCardPartsCoordinate := newBusinessCardPartsCoordinate()

	businessCard := newBusinessCard(
		user,
		arAsset,
		businessCardPartsCoordinate[0],
		businessCardBackgrounds[0],
	)

	return &AppModel{
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

func newUser() *model.User {
	return &model.User{
		ID:                testdata.DEV_UID,
		RecordedModelPath: "seed_recorded_model.glb",
		IsToured:          false,
		Status:            model.GormStatusCompleted,
	}
}

func newThreeDimentionalModelTemplate(id string) *model.ThreeDimentionalModelTemplate {
	return &model.ThreeDimentionalModelTemplate{
		ID: id,
	}
}

func newPersonalThreeDimentionalModel(id string, user *model.User) *model.PersonalThreeDimentionalModel {
	return &model.PersonalThreeDimentionalModel{
		ID:     id,
		UserID: user.ID,
	}
}

func newThreeDimentionalModels() []*model.ThreeDimentionalModel {
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

func newSpeakingAsset(user *model.User) *model.SpeakingAsset {
	return &model.SpeakingAsset{
		ID:          newID(),
		UserID:      user.ID,
		Description: "こんにちは",
		AudioPath:   "seed_example.mp3",
	}
}

func newARAsset(
	user *model.User,
	speakingAsset *model.SpeakingAsset,
	threeDimentionalModel *model.ThreeDimentionalModel,
) *model.ARAsset {
	return &model.ARAsset{
		ID:                      "1",
		AccessCount:             0,
		QRCodeImagePath:         "seed_qr_code.png",
		UserID:                  user.ID,
		SpeakingAssetID:         speakingAsset.ID,
		ThreeDimentionalModelID: threeDimentionalModel.ID,
		Status:                  model.GormStatusCompleted,
	}
}

func newBusinessCardPartsCoordinate() []*model.BusinessCardPartsCoordinate {
	return []*model.BusinessCardPartsCoordinate{
		{
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
		},
		{
			ID:                newID(),
			DisplayNameX:      639,
			DisplayNameY:      83,
			CompanyNameX:      646,
			CompanyNameY:      230,
			DepartmentX:       646,
			DepartmentY:       284,
			OfficialPositionX: 646,
			OfficialPositionY: 323,
			PhoneNumberX:      646,
			PhoneNumberY:      455,
			EmailX:            646,
			EmailY:            408,
			PostalCodeX:       646,
			PostalCodeY:       544,
			AddressX:          646,
			AddressY:          588,
			QRCodeX:           156,
			QRCodeY:           209,
		},
		{
			ID:                newID(),
			DisplayNameX:      627,
			DisplayNameY:      153,
			CompanyNameX:      80,
			CompanyNameY:      82,
			DepartmentX:       192,
			DepartmentY:       182,
			OfficialPositionX: 406,
			OfficialPositionY: 182,
			PhoneNumberX:      640,
			PhoneNumberY:      397,
			EmailX:            640,
			EmailY:            357,
			PostalCodeX:       640,
			PostalCodeY:       468,
			AddressX:          640,
			AddressY:          507,
			QRCodeX:           156,
			QRCodeY:           303,
		},
	}
}

func newBusinessCard(
	user *model.User,
	arAsset *model.ARAsset,
	businessCardPartsCoordinate *model.BusinessCardPartsCoordinate,
	businessCardBackgrounds *model.BusinessCardBackground,
) *model.BusinessCard {
	return &model.BusinessCard{
		ID:                            "1",
		UserID:                        user.ID,
		ARAssetID:                     arAsset.ID,
		BusinessCardPartsCoordinateID: businessCardPartsCoordinate.ID,
		BusinessCardBackgroundID:      businessCardBackgrounds.ID,
		DisplayName:                   "山田太郎",
		CompanyName:                   "株式会社山田",
		Department:                    "開発部",
		OfficialPosition:              "部長",
		PhoneNumber:                   "090-1234-5678",
		Email:                         "yamada@example.com",
		PostalCode:                    "123-4567",
		Address:                       "東京都渋谷区",
	}
}

func newID() string {
	id, err := shared.NewID()
	if err != nil {
		panic(err)
	}

	return id.String()
}
