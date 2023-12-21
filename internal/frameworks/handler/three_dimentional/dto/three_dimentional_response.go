package dto

type ThreeDimentionalResponse struct {
	ID   string `json:"id" binding:"required"`
	Path string `json:"path" binding:"required"`
}
