package main

import (
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

	seedPRD(db)
}

func seedPRD(db *gorm.DB) {
	appModel := helper.NewAppModelPRD()

	// User
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.User)

	// ThreeDimentionalModel
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ThreeDimentionalModels)

	// ThreeDimentionalModelTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.TempThreeDimentionalModelTemplate)

	// SpeakingAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.SpeakingAsset)

	// ARAsset
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.ARAsset)

	// BusinessCardBackground
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgrounds)

	// BusinessCardBackgroundTemplate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardBackgroundTemplate)

	// BusinessCardPartsCoordinate
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(appModel.BusinessCardPartsCoordinate)
}
