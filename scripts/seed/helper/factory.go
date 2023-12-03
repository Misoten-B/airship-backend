package helper

import (
	"github.com/Misoten-B/airship-backend/internal/database/model"
	"github.com/Misoten-B/airship-backend/internal/id"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type AppModel struct {
	User                              *model.User
	TempThreeDimentionalModelTemplate *model.ThreeDimentionalModelTemplate
	PersonalThreeDimentionalModel     *model.PersonalThreeDimentionalModel
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	SpeakingAsset                     *model.SpeakingAsset
	ARAsset                           *model.ARAsset
	BusinessCardBackgroundTemplate    *model.BusinessCardBackgroundTemplate
	PersonalBusinessCardBackground    *model.PersonalBusinessCardBackground
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       *model.BusinessCardPartsCoordinate
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
			ID:        "1",
			ColorCode: "#000000",
			ImagePath: "https://example.com/background-template.png",
		},
		{
			ID:        "2",
			ColorCode: "#ffffff",
			ImagePath: "https://example.com/background-personal.png",
		},
	}
	businessCardBackgroundTemplate := &model.BusinessCardBackgroundTemplate{
		ID: businessCardBackgrounds[0].ID,
	}
	personalBusinessCardBackground := &model.PersonalBusinessCardBackground{
		ID:     businessCardBackgrounds[1].ID,
		UserID: user.ID,
	}

	businessCardPartsCoordinate := newBusinessCardPartsCoordinate()

	businessCard := newBusinessCard(
		user,
		arAsset,
		businessCardPartsCoordinate,
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
		RecordedVoicePath: "https://example.com/recorded_voice.mp3",
		RecordedModelPath: "https://example.com/recorded_model.gltf",
		IsToured:          false,
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
			ID:        "1",
			ModelPath: "https://example.com/3d_model.gltf",
		},
		{
			ID:        "2",
			ModelPath: "https://example.com/3d_model.gltf",
		},
	}
}

func newSpeakingAsset(user *model.User) *model.SpeakingAsset {
	return &model.SpeakingAsset{
		ID:     newID(),
		UserID: user.ID,
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
		QRCodeImagePath:         "https://example.com/qr_code.png",
		UserID:                  user.ID,
		SpeakingAssetID:         speakingAsset.ID,
		ThreeDimentionalModelID: threeDimentionalModel.ID,
	}
}

func newBusinessCardPartsCoordinate() *model.BusinessCardPartsCoordinate {
	return &model.BusinessCardPartsCoordinate{
		ID:                newID(),
		DisplayNameX:      0,
		DisplayNameY:      0,
		CompanyNameX:      0,
		CompanyNameY:      0,
		DepartmentX:       0,
		DepartmentY:       0,
		OfficialPositionX: 0,
		OfficialPositionY: 0,
		PhoneNumberX:      0,
		PhoneNumberY:      0,
		EmailX:            0,
		EmailY:            0,
		PostalCodeX:       0,
		PostalCodeY:       0,
		AddressX:          0,
		AddressY:          0,
		QRCodeX:           0,
		QRCodeY:           0,
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
	id, err := id.NewID()
	if err != nil {
		panic(err)
	}

	return id.String()
}
