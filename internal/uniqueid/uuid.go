package uniqueid

import "github.com/google/uuid"

func NewUUID() (uuid.UUID, error) {
	return uuid.NewRandom()
}
