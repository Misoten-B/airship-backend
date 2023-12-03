package model

// ThreeDimentionalModelは3DモデルテーブルのORMモデルです。
type ThreeDimentionalModel struct {
	ID                             string `gorm:"primaryKey"`
	Path                           string
	ARAssets                       []ARAsset
	ThreeDimentionalModelTemplates []ThreeDimentionalModelTemplate `gorm:"foreignKey:ID"`
	PersonalThreeDimentionalModels []PersonalThreeDimentionalModel `gorm:"foreignKey:ID"`
}

// ThreeDimentionalModelTemplateは3DモデルテンプレートテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type ThreeDimentionalModelTemplate struct {
	ID string `gorm:"primaryKey"`
}

// PersonalThreeDimentionalModelはユーザー定義3DモデルテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type PersonalThreeDimentionalModel struct {
	ID     string `gorm:"primaryKey"`
	UserID string
}

// SpeakingAssetは音声アセットテーブルのORMモデルです。
type SpeakingAsset struct {
	ID          string `gorm:"primaryKey"`
	ARAssets    []ARAsset
	UserID      string
	Description string
	AudioPath   string
}

// ARAssetはARアセットテーブルのORMモデルです。
type ARAsset struct {
	ID                      string         `gorm:"primaryKey"`
	BusinessCards           []BusinessCard `gorm:"foreignKey:ARAsset"`
	UserID                  string
	SpeakingAssetID         string
	ThreeDimentionalModelID string
}
