package helper

import (
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type AppModelPRD struct {
	User                              *model.User
	TempThreeDimentionalModelTemplate []*model.ThreeDimentionalModelTemplate
	ThreeDimentionalModels            []*model.ThreeDimentionalModel
	BusinessCardBackgroundTemplate    []*model.BusinessCardBackgroundTemplate
	BusinessCardBackgrounds           []*model.BusinessCardBackground
	BusinessCardPartsCoordinate       []*model.BusinessCardPartsCoordinate
	SpeakingAsset                     *model.SpeakingAsset
	ARAsset                           *model.ARAsset
}

func NewAppModelPRD() *AppModelPRD {
	threeDimentionalModels := newThreeDimentionalModelsPRD()
	threeDimentionalModelTemplate := newThreeDimentionalModelTemplatePRD(threeDimentionalModels)

	businessCardBackgrounds := []*model.BusinessCardBackground{
		{
			ID:        "background-template-1",
			ColorCode: "#ffffff",
			ImagePath: "background-template-1.png",
		},
		{
			ID:        "background-template-2",
			ColorCode: "#ffffff",
			ImagePath: "background-template-2.png",
		},
		{
			ID:        "background-template-3",
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
		SpeakingAsset:                     newSpeakingAssetsModelPRD(),
		ARAsset:                           newARAssetsModelPRD(threeDimentionalModels[0].ID),
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

func newThreeDimentionalModelTemplatePRD(tdms []*model.ThreeDimentionalModel) []*model.ThreeDimentionalModelTemplate {
	models := []*model.ThreeDimentionalModelTemplate{}
	for _, tdm := range tdms {
		models = append(models, &model.ThreeDimentionalModelTemplate{
			ID: tdm.ID,
		})
	}
	return models
}

func newThreeDimentionalModelsPRD() []*model.ThreeDimentionalModel {
	return []*model.ThreeDimentionalModel{
		{
			ID:        "chicken",
			ModelPath: "chicken.glb",
		},
		{
			ID:        "dog",
			ModelPath: "dog.glb",
		},
		{
			ID:        "pinguin",
			ModelPath: "pinguin.glb",
		},
		{
			ID:        "tiger",
			ModelPath: "tiger.glb",
		},
	}
}

func newSpeakingAssetsModelPRD() *model.SpeakingAsset {
	return &model.SpeakingAsset{
		ID:          testdata.DEV_UID,
		UserID:      testdata.DEV_UID,
		Description: "xxx",
		AudioPath:   fmt.Sprint(testdata.DEV_UID, ".wav"),
	}
}

func newARAssetsModelPRD(tdmID string) *model.ARAsset {
	return &model.ARAsset{
		ID:                      testdata.DEV_UID,
		AccessCount:             0,
		QRCodeImagePath:         "",
		Status:                  model.GormStatusCompleted,
		UserID:                  testdata.DEV_UID,
		SpeakingAssetID:         testdata.DEV_UID,
		ThreeDimentionalModelID: tdmID,
	}
}
