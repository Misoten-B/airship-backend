package main

import (
	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/scripts/seed/helper"
	"gorm.io/gorm/clause"
)

func main() {
	// データベースに接続
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	appModel := helper.NewAppModel()

	// User
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.User)

	// ThreeDimentionalModelTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.TempThreeDimentionalModelTemplate)

	// PersonalThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.PersonalThreeDimentionalModel)

	// ThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ThreeDimentionalModels)

	// SpeakingAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.SpeakingAsset)

	// ARAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ARAsset)

	// BusinessCardBackgroundTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgroundTemplate)

	// PersonalBusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.PersonalBusinessCardBackground)

	// BusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgrounds)

	// BusinessCardPartsCoordinate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardPartsCoordinate)

	// BusinessCard
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCard)
}
