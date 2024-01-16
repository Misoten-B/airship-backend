package drivers

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
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

func (d *AzureBlobDriver) SaveBlob(containerName string, file *file.File) error {
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

func (d *AzureBlobDriver) DeleteBlob(containerName string, blobName string) error {
	ctx := context.Background()

	serviceClient, err := d.newClient()
	if err != nil {
		return fmt.Errorf("failed to create service client: %w", err)
	}

	_, err = serviceClient.DeleteBlob(ctx, containerName, blobName, nil)
	if err != nil {
		return fmt.Errorf("failed to delete blob: %w", err)
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

func (d *AzureBlobDriver) GetContainerURL(containerName string) (AzureBlobContainerFullPath, error) {
	serviceClient, err := d.newClient()
	if err != nil {
		return AzureBlobContainerFullPath{}, err
	}

	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	permissions := sas.ContainerPermissions{
		Read: true,
	}
	expiry := time.Now().Add(sasExpiryDuration)

	url, err := containerClient.GetSASURL(permissions, expiry, nil)
	if err != nil {
		return AzureBlobContainerFullPath{}, err
	}

	afp, err := newContainerFullPath(url)
	if err != nil {
		return AzureBlobContainerFullPath{}, err
	}

	return afp, nil
}

func (d *AzureBlobDriver) newClient() (*azblob.Client, error) {
	serviceClient, err := azblob.NewClientFromConnectionString(d.connectionString, nil)
	return serviceClient, err
}

type AzureBlobContainerFullPath struct {
	rootPath string
	token    string
}

func newContainerFullPath(azureURL string) (AzureBlobContainerFullPath, error) {
	parsedURL, err := url.Parse(azureURL)
	if err != nil {
		return AzureBlobContainerFullPath{}, err
	}

	rootPath := parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path
	token := parsedURL.Query().Encode()

	return AzureBlobContainerFullPath{
		rootPath: rootPath,
		token:    token,
	}, nil
}

func (a *AzureBlobContainerFullPath) Path(name string) string {
	return fmt.Sprintf("%s/%s?%s", a.rootPath, name, a.token)
}
