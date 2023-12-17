package dto

type BackgroundResponse struct {
	ID                          string `json:"id"`
	BusinessCardBackgroundColor string `json:"businessCardBackgroundColor" example:"#ffffff"`
	BusinessCardBackgroundImage string `json:"businessCardBackgroundImage" example:"url"`
}
