package domain

import (
	"time"

	"github.com/Misoten-B/airship-backend/internal/database/model"
)

type User struct {
	id        string
	createdAt time.Time
}

func (u *User) toModel() *model.User {
	return &model.User{
		ID:        u.id,
		CreatedAt: u.createdAt,
	}
}

func NewUser() *User {
	createdAt := time.Now()

	return &User{
		id:        "test-id",
		createdAt: createdAt,
	}
}

func fromModel(user *model.User) *User {
	return &User{
		id:        user.ID,
		createdAt: user.CreatedAt,
	}
}
