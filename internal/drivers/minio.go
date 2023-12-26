package drivers

import (
	"context"
	"fmt"

	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	minioEndpoint        = "localhost:9000"
	minioAccessKeyID     = "minio"
	minioSecretAccessKey = "minio123"
)

func newMinioClient() (*minio.Client, error) {
	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func MinioPutObject(bucketName string, objectName string, file file.File) error {
	client, err := newMinioClient()
	if err != nil {
		return fmt.Errorf("failed to create minio client: %w", err)
	}

	_, err = client.PutObject(
		context.Background(),
		bucketName,
		objectName,
		file.File(),
		file.FileHeader().Size,
		minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to put object: %w", err)
	}

	return nil
}
