package model

import "time"

// BusinessCardは名刺テーブルのORMモデルです。
type BusinessCard struct {
	ID                            string `gorm:"primaryKey"`
	UserID                        string `gorm:"size:255"`
	ARAssetID                     string `gorm:"size:255"`
	BusinessCardPartsCoordinateID string `gorm:"size:255"`
	BusinessCardBackgroundID      string `gorm:"size:255"`
	BusinessCardName              string
	DisplayName                   string
	CompanyName                   string
	Department                    string
	OfficialPosition              string
	PhoneNumber                   string
	Email                         string
	PostalCode                    string
	Address                       string
	CreatedAt                     time.Time
}

// BusinessCardPartsCoordinateは名刺パーツ座標テーブルのORMモデルです。
type BusinessCardPartsCoordinate struct {
	ID                string `gorm:"primaryKey"`
	BusinessCards     []BusinessCard
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

// BusinessCardBackgroundは名刺背景テーブルのORMモデルです。
type BusinessCardBackground struct {
	ID                              string `gorm:"primaryKey"`
	ColorCode                       string
	ImagePath                       string
	BusinessCardBackgroundTemplates []BusinessCardBackgroundTemplate `gorm:"foreignKey:ID"`
	PersonalBusinessCardBackgrounds []PersonalBusinessCardBackground `gorm:"foreignKey:ID"`
	BusinessCards                   []BusinessCard
}

// BusinessCardBackgroundTemplateは名刺背景テンプレートテーブルのORMモデルです。
// BusinessCardBackgroundの排他的サブタイプです。
type BusinessCardBackgroundTemplate struct {
	ID string `gorm:"primaryKey"`
}

// PersonalBusinessCardBackgroundはユーザー定義名刺背景テーブルのORMモデルです。
// BusinessCardBackgroundの排他的サブタイプです。
type PersonalBusinessCardBackground struct {
	ID     string `gorm:"primaryKey"`
	UserID string `gorm:"size:255"`
}
