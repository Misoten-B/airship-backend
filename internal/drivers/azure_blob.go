package drivers

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/config"
	"github.com/Misoten-B/airship-backend/internal/file"
)

type AzureBlobDriver struct {
	connectionString string
}

const (
	sasExpiryDuration = 24 * time.Hour
)

func NewAzureBlobDriver(config *config.Config) *AzureBlobDriver {
	return &AzureBlobDriver{
		connectionString: config.AzureBlobStorageConnectionString,
	}
}

func (d *AzureBlobDriver) SaveBlob(containerName string, file file.File) error {
	ctx := context.Background()

	serviceClient, err := d.newClient()
	if err != nil {
		return err
	}

	_, err = serviceClient.UploadStream(ctx, containerName, file.FileHeader().Filename, file.File(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (d *AzureBlobDriver) GetBlobURL(containerName, blobName string) (string, error) {
	serviceClient, err := d.newClient()
	if err != nil {
		return "", err
	}

	blobClient := serviceClient.ServiceClient().
		NewContainerClient(containerName).
		NewBlobClient(blobName)

	permissions := sas.BlobPermissions{
		Read: true,
	}
	expiry := time.Now().Add(sasExpiryDuration)

	url, err := blobClient.GetSASURL(permissions, expiry, nil)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (d *AzureBlobDriver) newClient() (*azblob.Client, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(d.connectionString, nil)
	return serviceClient, err
}
