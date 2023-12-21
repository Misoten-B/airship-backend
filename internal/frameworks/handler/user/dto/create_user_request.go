package dto

type CreateUserRequest struct {
	IsToured bool `form:"isToured" example:"false"`
}
