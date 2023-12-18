package main

import (
	"os"

	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/scripts/seed/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	// データベースに接続
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	value, ok := os.LookupEnv("DEV_MODE")
	if ok && value == "true" {
		// 開発環境
		seed(db)
	} else {
		// 本番環境
		seedPRD(db)
	}
}

func seed(db *gorm.DB) {
	appModel := helper.NewAppModel()

	// User
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.User)

	// ThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ThreeDimentionalModels)

	// ThreeDimentionalModelTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.TempThreeDimentionalModelTemplate)

	// PersonalThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.PersonalThreeDimentionalModel)

	// SpeakingAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.SpeakingAsset)

	// ARAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ARAsset)

	// BusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgrounds)

	// BusinessCardBackgroundTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgroundTemplate)

	// PersonalBusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.PersonalBusinessCardBackground)

	// BusinessCardPartsCoordinate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardPartsCoordinate)

	// BusinessCard
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCard)
}

func seedPRD(db *gorm.DB) {
	appModel := helper.NewAppModelPRD()

	// ThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ThreeDimentionalModels)

	// ThreeDimentionalModelTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.TempThreeDimentionalModelTemplate)

	// BusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgrounds)

	// BusinessCardBackgroundTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgroundTemplate)

	// BusinessCardPartsCoordinate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardPartsCoordinate)
}
