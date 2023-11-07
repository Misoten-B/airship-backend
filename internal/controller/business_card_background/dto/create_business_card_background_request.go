package dto

type CreateBusinessCardBackgroundRequest struct {
	BusinessCardBackgroundColor *string `json:"business_card_background" example:"#ffffff"`
	BusinessCardBackgroundImage *string `json:"business_card_background_image" example:"url" extensions:"x-nullable"`
}
