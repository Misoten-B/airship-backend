package dto

type UserResponse struct {
	ID                string `json:"id" binding:"required"`
	RecordedModelPath string `json:"recordedModelPath" binding:"required"`
	IsToured          bool   `json:"isToured" binding:"required"`
	Status            int    `json:"status" binding:"required"`
}
