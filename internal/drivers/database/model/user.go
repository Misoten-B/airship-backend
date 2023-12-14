package model

import (
	"time"

	"gorm.io/gorm"
)

// UserはユーザーテーブルのORMモデルです。
type User struct {
	ID                              string `gorm:"primaryKey"`
	BusinessCards                   []BusinessCard
	PersonalBusinessCardBackgrounds []PersonalBusinessCardBackground
	PersonalThreeDimentionalModels  []PersonalThreeDimentionalModel
	ARAssets                        []ARAsset
	SpeakingAssets                  []SpeakingAsset
	RecordedModelPath               string
	CreatedAt                       time.Time
	DeletedAt                       gorm.DeletedAt
	IsToured                        bool
	// Status は音声モデルの生成状態を表します
	Status int
}
