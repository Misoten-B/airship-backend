package voice

import (
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/go-resty/resty/v2"
)

type ExternalAPIVoiceModelAdapter struct{}

func NewExternalAPIVoiceModelAdapter() *ExternalAPIVoiceModelAdapter {
	return &ExternalAPIVoiceModelAdapter{}
}

func (a *ExternalAPIVoiceModelAdapter) GenerateAudioFile(request service.GenerateAudioFileRequest) error {
	url := fmt.Sprintf("http://localhost:8080/voice-model/%s/sound", request.UID)
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"output_file_path": request.OutputFilePath,
			"language":         request.Language,
			"content":          request.Content,
		}).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}
