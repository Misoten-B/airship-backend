package shared

import (
	"github.com/Misoten-B/airship-backend/internal/uniqueid"
)

type ID string

const (
	invalidID = ID("")
)

func NewID() (ID, error) {
	value, err := uniqueid.NewULID()
	if err != nil {
		return invalidID, err
	}

	return ID(value.String()), nil
}

func ReconstructID(id string) ID {
	return ID(id)
}

func (i ID) String() string {
	return string(i)
}

func (i ID) Equals(other ID) bool {
	return i == other
}

func (i ID) IsValid() bool {
	return !i.Equals(invalidID)
}
