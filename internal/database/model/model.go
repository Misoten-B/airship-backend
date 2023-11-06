package model

import (
	"time"
)

// UserはユーザーテーブルのORMモデルです。
type User struct {
	ID                string         `gorm:"primaryKey"`
	BusinessCards     []BusinessCard `gorm:"foreignKey:User"`
	RecordedVoicePath string
	RecordedModelPath string
	CreatedAt         time.Time
	DeletedAt         time.Time
	IsToured          bool
}

// BusinessCardは名刺テーブルのORMモデルです。
type BusinessCard struct {
	ID   string `gorm:"primaryKey"`
	User string
	// ARAssets
	// BusinessCardPathsCordinate
	// BusinessCardBackground
	BusinessCardName string
	DisplayName      string
	CompanyName      string
	Department       string
	OfficialPosition string
	PhoneNumber      string
	Email            string
	Address          string
	QRCodeImagePath  string
	CreatedAt        time.Time
	AccessCount      int
}
