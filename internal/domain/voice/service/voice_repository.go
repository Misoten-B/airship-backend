package service

import (
	"log"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type VoiceRepository interface {
	FetchStatus(userID shared.ID) (shared.Status, error)
}

type MockVoiceRepository struct{}

func NewMockVoiceRepository() *MockVoiceRepository {
	return &MockVoiceRepository{}
}

func (m *MockVoiceRepository) FetchStatus(_ shared.ID) (shared.Status, error) {
	log.Print("Mock VoiceRepository - FetchStatus")
	return shared.StatusCompleted{}, nil
}
