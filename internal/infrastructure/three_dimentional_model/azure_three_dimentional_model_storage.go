package threedimentionalmodel

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
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

func (s *AzureThreeDimentionalModelStorage) Save(
	threeDimentionalModel threedimentionalmodel.ThreeDimensionalModelFile,
) error {
	err := s.dirver.SaveBlob(containerName, threeDimentionalModel.File())
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	url, err := s.dirver.GetBlobURL(containerName, modelName)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *AzureThreeDimentionalModelStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	fullPath, err := s.dirver.GetContainerURL(containerName)
	if err != nil {
		return nil, err
	}

	return &fullPath, nil
}
