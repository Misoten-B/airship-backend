package database

import (
	"time"
)

// Userはデータベースのusersテーブルに対応するモデルです。
type User struct {
	ID                string `gorm:"primaryKey"`
	RecordedVoicePath string
	RecordedModelPath string
	IsToured          bool
	CreatedAt         time.Time
	DeletedAt         time.Time
}
