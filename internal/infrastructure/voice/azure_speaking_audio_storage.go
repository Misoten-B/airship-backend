package voice

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers"
)

type AzureSpeakingAudioStorage struct {
	dirver *drivers.AzureBlobDriver
}

const (
	containerName = "voice-sounds"
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

func (s *AzureSpeakingAudioStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	fullPath, err := s.dirver.GetContainerURL(containerName)
	if err != nil {
		return nil, err
	}

	return &fullPath, nil
}
