package voice

import (
	"fmt"
	"mime/multipart"

	"github.com/go-resty/resty/v2"
)

type VallEXAdapter struct{}

func NewVallEXAdapter() *VallEXAdapter {
	return &VallEXAdapter{}
}

type GenerateVoiceModelRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func (a *VallEXAdapter) GenerateVoiceModel(userID string, request GenerateVoiceModelRequest) error {
	url := fmt.Sprintf("http://localhost:8000/voice-model/%s", userID)
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFileReader("file", request.FileHeader.Filename, request.File).
		Post(url)
	if err != nil {
		return err
	}

	return nil
}

type GenerateAudioFileRequest struct {
	ARAssetsID string
	Text       string
}

func (a *VallEXAdapter) GenerateAudioFile(userID string, request GenerateAudioFileRequest) error {
	url := fmt.Sprintf("http://localhost:8000/voice-model/%s/audio", userID)
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"ar_assets_id": request.ARAssetsID,
			"text":         request.Text,
		}).
		Post(url)
	if err != nil {
		return err
	}

	return nil
}
