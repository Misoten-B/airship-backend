package arassets

import (
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers"
)

type AzureQRCodeImageStorage struct {
	dirver *drivers.AzureBlobDriver
}

const (
	containerName = "qrcode-images"
)

func NewAzureQRCodeImageStorage(driver *drivers.AzureBlobDriver) *AzureQRCodeImageStorage {
	return &AzureQRCodeImageStorage{
		dirver: driver,
	}
}

func (s *AzureQRCodeImageStorage) Save(qrCodeImage arassets.QRCodeImage) error {
	err := s.dirver.SaveBlob(containerName, qrCodeImage.File())
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureQRCodeImageStorage) GetImageURL(name string) (string, error) {
	url, err := s.dirver.GetBlobURL(containerName, name)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *AzureQRCodeImageStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	fullPath, err := s.dirver.GetContainerURL(containerName)
	if err != nil {
		return nil, err
	}

	return &fullPath, nil
}
