package voice

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

type AzureSpeakingAudioStorage struct {
	connectionString string
}

const (
	// sasExpiryDuration は、SASの有効期限の猶予です。
	sasExpiryDuration = 24 * time.Hour
	containerName     = "speaking-audios"
)

func NewAzureSpeakingAudioStorage(connectionString string) *AzureSpeakingAudioStorage {
	return &AzureSpeakingAudioStorage{
		connectionString: connectionString,
	}
}

func (s *AzureSpeakingAudioStorage) GetAudioURL(name string) (string, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(s.connectionString, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create service client: %w", err)
	}

	blobClient := serviceClient.ServiceClient().
		NewContainerClient(containerName).
		NewBlobClient(name)

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
