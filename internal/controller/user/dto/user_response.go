package dto

type UserResponse struct {
	ID                string `json:"id"`
	RecordedVoicePath string `json:"recorded_voice_path"`
	RecordedModelPath string `json:"recorded_model_path"`
	IsToured          bool   `json:"is_toured"`
}
