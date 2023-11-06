package domain

import (
	"time"
)

type User struct {
	id        string
	createdAt time.Time
}

func NewUser() *User {
	createdAt := time.Now()

	return &User{
		id:        "test-id",
		createdAt: createdAt,
	}
}
