package voice

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/go-resty/resty/v2"
)

type ExternalAPIVoiceModelAdapter struct{}

func NewExternalAPIVoiceModelAdapter() *ExternalAPIVoiceModelAdapter {
	return &ExternalAPIVoiceModelAdapter{}
}

func (a *ExternalAPIVoiceModelAdapter) GenerateAudioFile(request service.GenerateAudioFileRequest) error {
	url := fmt.Sprintf("https://airship-ml.japaneast.cloudapp.azure.com/voice-model/%s/sound", request.UID)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"ar_assets_id": request.ARAssetsID,
			"language":     request.Language,
			"content":      request.Content,
		}).
		Post(url)
	if resp.StatusCode() != http.StatusOK {
		return errors.New("failed to generate audio file")
	}

	if err != nil {
		return err
	}
	return nil
}
