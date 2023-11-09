package domain

import (
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
}

func NewUser() *User {
	createdAt := time.Now()

	return &User{
		ID:        "test-id",
		CreatedAt: createdAt,
	}
}
