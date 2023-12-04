package arassets

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/config"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
)

type AzureQRCodeImageStorage struct {
	connectionString string
}

const (
	// sasExpiryDuration は、SASの有効期限の猶予です。
	sasExpiryDuration = 24 * time.Hour
	containerName     = "images"
)

func NewAzureQRCodeImageStorage(config *config.Config) *AzureQRCodeImageStorage {
	return &AzureQRCodeImageStorage{
		connectionString: config.AzureBlobStorageConnectionString,
	}
}

func (s *AzureQRCodeImageStorage) Save(qrCodeImage arassets.QRCodeImage) error {
	ctx := context.Background()

	serviceClient, err := s.newClient()
	if err != nil {
		return err
	}

	_, err = serviceClient.UploadStream(ctx, containerName, qrCodeImage.Name(), qrCodeImage.File(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *AzureQRCodeImageStorage) GetImageURL(qrCodeImage arassets.QRCodeImage) (string, error) {
	serviceClient, err := s.newClient()
	if err != nil {
		return "", err
	}

	blobClient := serviceClient.ServiceClient().
		NewContainerClient(containerName).
		NewBlobClient(qrCodeImage.Name())

	permissions := sas.BlobPermissions{
		Read: true,
	}
	expiry := time.Now().Add(sasExpiryDuration)

	url, err := blobClient.GetSASURL(permissions, expiry, nil)
	if err != nil {
		return "", err
	}

	return url, err
}

func (s *AzureQRCodeImageStorage) newClient() (*azblob.Client, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(s.connectionString, nil)
	return serviceClient, err
}
