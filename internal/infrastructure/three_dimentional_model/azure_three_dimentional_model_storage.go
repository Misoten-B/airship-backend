package threedimentionalmodel

import (
	"github.com/Misoten-B/airship-backend/internal/drivers"
)

type AzureThreeDimentionalModelStorage struct {
	dirver *drivers.AzureBlobDriver
}

const (
	containerName = "three-dimentional-models"
)

func NewAzureThreeDimentionalModelStorage(driver *drivers.AzureBlobDriver) *AzureThreeDimentionalModelStorage {
	return &AzureThreeDimentionalModelStorage{
		dirver: driver,
	}
}

func (s *AzureThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	url, err := s.dirver.GetBlobURL(containerName, modelName)
	if err != nil {
		return "", err
	}

	return url, nil
}
