package model

import "time"

// UserはユーザーテーブルのORMモデルです。
type User struct {
	ID                              string `gorm:"primaryKey"`
	BusinessCards                   []BusinessCard
	PersonalBusinessCardBackgrounds []PersonalBusinessCardBackground
	PersonalThreeDimentionalModels  []PersonalThreeDimentionalModel
	ARAssets                        []ARAsset
	SpeakingAssets                  []SpeakingAsset
	RecordedVoicePath               string
	RecordedModelPath               string
	CreatedAt                       time.Time
	DeletedAt                       time.Time
	IsToured                        bool
}
