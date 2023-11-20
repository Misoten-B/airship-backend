package id

import "github.com/google/uuid"

type ID string

func NewID() (ID, error) {
	id, err := uuid.NewRandom()
	return ID(id.String()), err
}

func (i ID) String() string {
	return string(i)
}
