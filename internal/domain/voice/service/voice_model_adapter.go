package service

import "log"

type VoiceModelAdapter interface {
	GenerateAudioFile(request GenerateAudioFileRequest) error
}

type GenerateAudioFileRequest struct {
	UID        string
	ARAssetsID string
	Language   string
	Content    string
}

type MockVoiceModelAdapter struct{}

func NewMockVoiceModelAdapter() *MockVoiceModelAdapter {
	return &MockVoiceModelAdapter{}
}

func (a *MockVoiceModelAdapter) GenerateAudioFile(_ GenerateAudioFileRequest) error {
	log.Print("Mock VoiceModel Adapter - Generate AudioFile")
	return nil
}
