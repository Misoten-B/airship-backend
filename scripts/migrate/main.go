package main

import (
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
)

func main() {
	// データベースに接続
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// マイグレーション
	if err = db.AutoMigrate(
		&model.User{},
		&model.SpeakingAsset{},
		&model.ThreeDimentionalModel{},
		&model.PersonalThreeDimentionalModel{}, &model.ThreeDimentionalModelTemplate{},
		&model.ARAsset{},
		&model.BusinessCardBackground{},
		&model.PersonalBusinessCardBackground{}, &model.BusinessCardBackgroundTemplate{},
		&model.BusinessCardPartsCoordinate{},
		&model.BusinessCard{},
	); err != nil {
		panic(err)
	}
}
