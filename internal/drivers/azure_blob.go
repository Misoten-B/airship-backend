package drivers

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/config"
	"github.com/Misoten-B/airship-backend/internal/file"
)

type AzureBlobDriver struct {
	connectionString string
}

type AzureFullPath struct {
	rootPath string
	token    string
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

func (d *AzureBlobDriver) GetContainerURL(containerName string) (AzureFullPath, error) {
	afp := AzureFullPath{}

	serviceClient, err := d.newClient()
	if err != nil {
		return afp, err
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	permissions := sas.ContainerPermissions{
		Read: true,
	}
	expiry := time.Now().Add(sasExpiryDuration)

	azureUrl, err := containerClient.GetSASURL(permissions, expiry, nil)
	if err != nil {
		return afp, err
	}

	parsedURL, err := url.Parse(azureUrl)
	if err != nil {
		return afp, err
	}

	afp.rootPath = parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path

	afp.token = parsedURL.Query().Encode()

	log.Print("afp: ", afp)
	return afp, nil
}

func (d *AzureBlobDriver) newClient() (*azblob.Client, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(d.connectionString, nil)
	return serviceClient, err
}
