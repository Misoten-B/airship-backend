package model

import "time"

// UserはユーザーテーブルのORMモデルです。
type User struct {
	ID                              string                           `gorm:"primaryKey"`
	BusinessCards                   []BusinessCard                   `gorm:"foreignKey:User"`
	PersonalBusinessCardBackgrounds []PersonalBusinessCardBackground `gorm:"foreignKey:User"`
	PersonalThreeDimentionalModels  []PersonalThreeDimentionalModel
	ARAssets                        []ARAsset
	SpeakingAssets                  []SpeakingAsset
	RecordedVoicePath               string
	RecordedModelPath               string
	CreatedAt                       time.Time
	DeletedAt                       time.Time
	IsToured                        bool
}
