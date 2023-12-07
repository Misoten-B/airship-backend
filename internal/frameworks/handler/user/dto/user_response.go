package dto

type UserResponse struct {
	ID                string `json:"id"`
	RecordedModelPath string `json:"recorded_model_path"`
	IsToured          bool   `json:"is_toured"`
}
