package dto

type UserResponse struct {
	ID                string `json:"id"`
	RecordedModelPath string `json:"recordedModelPath"`
	IsToured          bool   `json:"isToured"`
}
