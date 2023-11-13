package main

import (
	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/database/model"
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
		&model.ARAsset{},
		&model.PersonalThreeDimentionalModel{}, &model.ThreeDimentionalModelTemplate{},
		&model.ThreeDimentionalModel{},
		&model.PersonalBusinessCardBackground{}, &model.BusinessCardBackgroundTemplate{},
		&model.BusinessCardBackground{},
		&model.BusinessCardPartsCoordinate{},
		&model.BusinessCard{},
	); err != nil {
		panic(err)
	}
}