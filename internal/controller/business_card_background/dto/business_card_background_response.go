package dto

type BusinessCardBackgroundResponse struct {
	ID                          string `json:"id"`
	BusinessCardBackgroundColor string `json:"business_card_background_color" example:"#ffffff"`
	BusinessCardBackgroundImage string `json:"business_card_background_image" example:"url"`
}
