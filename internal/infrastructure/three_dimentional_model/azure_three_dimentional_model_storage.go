package threedimentionalmodel

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/config"
)

type AzureThreeDimentionalModelStorage struct {
	connectionString string
}

const (
	sasExpiryDuration = 24 * time.Hour
	containerName     = "three-dimentional-models"
)

func NewAzureThreeDimentionalModelStorage(config *config.Config) *AzureThreeDimentionalModelStorage {
	return &AzureThreeDimentionalModelStorage{
		connectionString: config.AzureBlobStorageConnectionString,
	}
}

func (s *AzureThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(s.connectionString, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create service client: %w", err)
	}

	blobClient := serviceClient.ServiceClient().
		NewContainerClient(containerName).
		NewBlobClient(modelName)

	permissions := sas.BlobPermissions{
		Read: true,
	}
	expiry := time.Now().Add(sasExpiryDuration)

	url, err := blobClient.GetSASURL(permissions, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get SAS URL: %w", err)
	}

	return url, nil
}
