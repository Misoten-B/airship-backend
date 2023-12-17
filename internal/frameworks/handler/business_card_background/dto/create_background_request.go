package dto

type CreateBackgroundRequest struct {
	BusinessCardBackgroundColor string `form:"businessCardBackgroundColor" example:"#ffffff"`
}
