package service

import (
	"fmt"
	"log"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type SpeakingAudioStorage interface {
	GetAudioURL(name string) (string, error)
	GetContainerFullPath() (shared.ContainerFullPath, error)
}

type MockSpeakingAudioStorage struct{}

func NewMockSpeakingAudioStorage() *MockSpeakingAudioStorage {
	return &MockSpeakingAudioStorage{}
}

func (s *MockSpeakingAudioStorage) GetAudioURL(name string) (string, error) {
	log.Println("Mock Speaking Audio Storage - GetURL")

	url := fmt.Sprintf("http://example.com/mock/%s", name)
	return url, nil
}

func (s *MockSpeakingAudioStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	return &shared.MockContainerFullPath{}, nil
}
