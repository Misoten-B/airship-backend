package service

import (
	"log"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type VoiceRepository interface {
	FetchStatus(userID id.ID) (shared.Status, error)
}

type MockVoiceRepository struct{}

func (m *MockVoiceRepository) FetchStatus(_ id.ID) (shared.Status, error) {
	log.Print("Mock VoiceRepository - FetchStatus")
	return shared.StatusCompleted{}, nil
}
