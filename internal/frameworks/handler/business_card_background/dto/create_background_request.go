package dto

type CreateBackgroundRequest struct {
	BusinessCardBackgroundColor string `form:"business_card_background_color" example:"#ffffff"`
}
