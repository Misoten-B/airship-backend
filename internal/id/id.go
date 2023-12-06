package id

import "github.com/google/uuid"

type ID string

func NewID() (ID, error) {
	id, err := uuid.NewRandom()
	return ID(id.String()), err
}

func ReconstructID(id string) ID {
	return ID(id)
}

func (i ID) String() string {
	return string(i)
}
