package service

import (
	"fmt"
	"log"
)

type SpeakingAudioStorage interface {
	GetAudioURL(name string) (string, error)
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
