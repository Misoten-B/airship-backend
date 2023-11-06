package model

import (
	"time"
)

// UserはユーザーテーブルのORMモデルです。
type User struct {
	ID                              string                           `gorm:"primaryKey"`
	BusinessCards                   []BusinessCard                   `gorm:"foreignKey:User"`
	PersonalBusinessCardBackgrounds []PersonalBusinessCardBackground `gorm:"foreignKey:User"`
	PersonalThreeDimentionalModels  []PersonalThreeDimentionalModel
	SpeakingAssets                  []SpeakingAsset
	RecordedVoicePath               string
	RecordedModelPath               string
	CreatedAt                       time.Time
	DeletedAt                       time.Time
	IsToured                        bool
}

// BusinessCardは名刺テーブルのORMモデルです。
type BusinessCard struct {
	ID   string `gorm:"primaryKey"`
	User string
	// ARAssets
	BusinessCardPathsCoordinate string
	BusinessCardBackground      string
	BusinessCardName            string
	DisplayName                 string
	CompanyName                 string
	Department                  string
	OfficialPosition            string
	PhoneNumber                 string
	Email                       string
	PostalCode                  string
	Address                     string
	QRCodeImagePath             string
	CreatedAt                   time.Time
	AccessCount                 int
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

// TODO: 排他サブタイプのモデル最適化

// BusinessCardBackgroundは名刺背景テーブルのORMモデルです。
type BusinessCardBackground struct {
	ID                             string         `gorm:"primaryKey"`
	BusinessCards                  []BusinessCard `gorm:"foreignKey:BusinessCardBackground"`
	BusinessCardBackgroundTemplate string
	PersonalBusinessCardBackground string
}

// BusinessCardBackgroundTemplateは名刺背景テンプレートテーブルのORMモデルです。
// BusinessCardBackgroundの排他的サブタイプです。
type BusinessCardBackgroundTemplate struct {
	ID                      string                   `gorm:"primaryKey"`
	BusinessCardBackgrounds []BusinessCardBackground `gorm:"foreignKey:BusinessCardBackgroundTemplate"`
	ColorCode               string
	ImagePath               string
}

// PersonalBusinessCardBackgroundはユーザー定義名刺背景テーブルのORMモデルです。
// BusinessCardBackgroundの排他的サブタイプです。
type PersonalBusinessCardBackground struct {
	ID                      string                   `gorm:"primaryKey"`
	BusinessCardBackgrounds []BusinessCardBackground `gorm:"foreignKey:PersonalBusinessCardBackground"`
	User                    string
	ColorCode               string
	ImagePath               string
}

// ThreeDimentionalModelは3DモデルテーブルのORMモデルです。
type ThreeDimentionalModel struct {
	ID string `gorm:"primaryKey"`
}

// ThreeDimentionalModelTemplateは3DモデルテンプレートテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type ThreeDimentionalModelTemplate struct {
	ID   string `gorm:"primaryKey"`
	Path string
}

// PersonalThreeDimentionalModelはユーザー定義3DモデルテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type PersonalThreeDimentionalModel struct {
	ID     string `gorm:"primaryKey"`
	UserID string
	Path   string
}

// SpeakingAssetは音声アセットテーブルのORMモデルです。
type SpeakingAsset struct {
	ID          string `gorm:"primaryKey"`
	UserID      string
	Description string
	AudioPath   string
}
