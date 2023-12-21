package dto

type BackgroundResponse struct {
	ID                          string `json:"id" binding:"required"`
	BusinessCardBackgroundColor string `json:"businessCardBackgroundColor" example:"#ffffff" binding:"required"`
	BusinessCardBackgroundImage string `json:"businessCardBackgroundImage" example:"url"`
}
