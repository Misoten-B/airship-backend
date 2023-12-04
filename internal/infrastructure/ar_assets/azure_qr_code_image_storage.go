package arassets

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Misoten-B/airship-backend/config"
)

type AzureQRCodeImageStorage struct {
	connectionString string
}

func NewAzureQRCodeImageStorage(config *config.Config) *AzureQRCodeImageStorage {
	return &AzureQRCodeImageStorage{
		connectionString: config.AzureBlobStorageConnectionString,
	}
}

func (s *AzureQRCodeImageStorage) Save(name string, file multipart.File) error {
	// ctx := context.Background()

	serviceClient, err := azblob.NewClientFromConnectionString(s.connectionString, nil)
	if err != nil {
		return err
	}

	url, err := serviceClient.ServiceClient().
		NewContainerClient("images").
		NewBlobClient("3b3baec3-7417-490f-bd2f-b8c700b4d1b1.png").
		GetSASURL(
			sas.BlobPermissions{
				Read: true,
			},
			time.Now().Add(24*time.Hour),
			nil,
		)
	fmt.Println(url)
	if err != nil {
		return err
	}

	// _, err = serviceClient.UploadStream(ctx, "images", name, file, &azblob.UploadStreamOptions{})
	// if err != nil {
	// 	return err
	// }
	return nil
}
