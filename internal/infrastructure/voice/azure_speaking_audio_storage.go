package voice

import (
	"github.com/Misoten-B/airship-backend/internal/drivers"
)

type AzureSpeakingAudioStorage struct {
	dirver *drivers.AzureBlobDriver
}

const (
	containerName = "speaking-audios"
)

func NewAzureSpeakingAudioStorage(driver *drivers.AzureBlobDriver) *AzureSpeakingAudioStorage {
	return &AzureSpeakingAudioStorage{
		dirver: driver,
	}
}

func (s *AzureSpeakingAudioStorage) GetAudioURL(name string) (string, error) {
	url, err := s.dirver.GetBlobURL(containerName, name)
	if err != nil {
		return "", err
	}

	return url, nil
}
