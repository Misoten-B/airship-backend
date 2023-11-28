package model

// TODO: 排他サブタイプのモデル最適化

// ThreeDimentionalModelは3DモデルテーブルのORMモデルです。
type ThreeDimentionalModel struct {
	ID                            string `gorm:"primaryKey"`
	ThreeDimentionalModelTemplate string `gorm:"default:null"`
	PersonalThreeDimentionalModel string `gorm:"default:null"`
	ARAssets                      []ARAsset
}

// ThreeDimentionalModelTemplateは3DモデルテンプレートテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type ThreeDimentionalModelTemplate struct {
	ID                     string                  `gorm:"primaryKey"`
	ThreeDimentionalModels []ThreeDimentionalModel `gorm:"foreignKey:ThreeDimentionalModelTemplate"`
	Path                   string
}

// PersonalThreeDimentionalModelはユーザー定義3DモデルテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type PersonalThreeDimentionalModel struct {
	ID                     string                  `gorm:"primaryKey"`
	ThreeDimentionalModels []ThreeDimentionalModel `gorm:"foreignKey:PersonalThreeDimentionalModel"`
	UserID                 string
	Path                   string
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
