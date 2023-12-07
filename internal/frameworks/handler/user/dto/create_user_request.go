package dto

type CreateUserRequest struct {
	IsToured bool `form:"is_toured" example:"false" extensions:"x-nullable"`
}
