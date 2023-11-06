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
	BusinessCardPathsCoordinate string
	// BusinessCardBackground
	BusinessCardName string
	DisplayName      string
	CompanyName      string
	Department       string
	OfficialPosition string
	PhoneNumber      string
	Email            string
	PostalCode       string
	Address          string
	QRCodeImagePath  string
	CreatedAt        time.Time
	AccessCount      int
}

// BusinessCardPartsCoordinateは名刺パーツ座標テーブルのORMモデルです。
type BusinessCardPartsCoordinate struct {
	ID                string         `gorm:"primaryKey"`
	BusinessCards     []BusinessCard `gorm:"foreignKey:BusinessCardPathsCoordinate"`
	DisplayNameX      int
	DisplayNameY      int
	CompanyNameX      int
	CompanyNameY      int
	DepartmentX       int
	DepartmentY       int
	OfficialPositionX int
	OfficialPositionY int
	PhoneNumberX      int
	PhoneNumberY      int
	EmailX            int
	EmailY            int
	PostalCodeX       int
	PostalCodeY       int
	AddressX          int
	AddressY          int
	QRCodeX           int
	QRCodeY           int
}
