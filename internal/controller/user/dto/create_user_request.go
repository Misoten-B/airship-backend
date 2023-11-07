package dto

type CreateUserRequest struct {
	RecordedVoicePath *string `json:"recorded_voice_path" example:"url" extensions:"x-nullable"`
	RecordedModelPath *string `json:"recorded_model_path" example:"url" extensions:"x-nullable"`
	IsToured          *bool   `json:"is_toured" example:"false" extensions:"x-nullable"`
}
