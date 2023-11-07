package dto

type CreateUserRequest struct {
	IsToured *bool `json:"is_toured" example:"false" extensions:"x-nullable"`
}
