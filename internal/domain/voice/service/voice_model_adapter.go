package service

import "log"

type VoiceModelAdapter interface {
	GenerateAudioFile(request GenerateAudioFileRequest) error
}

type GenerateAudioFileRequest struct {
	UID            string
	OutputFilePath string
	Language       string
	Content        string
}

type MockVoiceModelAdapter struct{}

func (a *MockVoiceModelAdapter) GenerateAudioFile(_ GenerateAudioFileRequest) error {
	log.Print("Mock VoiceModel Adapter - Generate AudioFile")
	return nil
}
