package domain

import "gorm.io/gorm"

type User struct {
	ID                string `json:"id"`
	RecordedVoicePath string
	RecordedModelPath string
	IsToured          bool   `json:"is_toured"`
	ModelKey          string `json:"model_key"`
}

func (user *User) toModel() {}

// TODO: database/modelに移動させる
// 場所は仮
type UserModel struct {
	gorm.Model
}

func (model *UserModel) toEntity() {
}
