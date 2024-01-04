package model

// ThreeDimentionalModelは3DモデルテーブルのORMモデルです。
type ThreeDimentionalModel struct {
	ID        string `gorm:"primaryKey"`
	ModelPath string
	ARAssets  []ARAsset
}

// ThreeDimentionalModelTemplateは3DモデルテンプレートテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type ThreeDimentionalModelTemplate struct {
	ThreeDimentionalModelID string                `gorm:"size:255"`
	ThreeDimentionalModel   ThreeDimentionalModel `gorm:"foreignKey:ThreeDimentionalModelID"`
}

// PersonalThreeDimentionalModelはユーザー定義3DモデルテーブルのORMモデルです。
// ThreeDimentionalModelの排他的サブタイプです。
type PersonalThreeDimentionalModel struct {
	ThreeDimentionalModelID string `gorm:"size:255"`
	ThreeDimentionalModel   ThreeDimentionalModel
	UserID                  string `gorm:"size:255"`
}

// SpeakingAssetは音声アセットテーブルのORMモデルです。
type SpeakingAsset struct {
	ID          string `gorm:"primaryKey"`
	UserID      string `gorm:"size:255"`
	Description string
	AudioPath   string
}

// ARAssetはARアセットテーブルのORMモデルです。
type ARAsset struct {
	ID                      string `gorm:"primaryKey"`
	AccessCount             int
	QRCodeImagePath         string
	Status                  int
	UserID                  string `gorm:"size:255"`
	User                    User
	SpeakingAssetID         string `gorm:"size:255"`
	SpeakingAsset           SpeakingAsset
	ThreeDimentionalModelID string `gorm:"size:255"`
	ThreeDimentionalModel   ThreeDimentionalModel
}
