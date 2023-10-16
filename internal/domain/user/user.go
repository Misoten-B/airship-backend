package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	ModelKey  string    `json:"model_key"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
