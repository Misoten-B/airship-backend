package domain

import (
	"time"

	"github.com/Misoten-B/airship-backend/internal/id"
)

type User struct {
	id        id.ID
	isToured  bool
	createdAt time.Time
	deletedAt time.Time
}

func NewUser() (*User, error) {
	createdAt := time.Now()

	id, err := id.NewID()
	if err != nil {
		return nil, err
	}

	return &User{
		id:        id,
		isToured:  false,
		createdAt: createdAt,
		deletedAt: time.Time{}, // zero value
	}, nil
}

func (u *User) ID() id.ID {
	return u.id
}

func (u *User) IsToured() bool {
	return u.isToured
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) DeletedAt() time.Time {
	return u.deletedAt
}
