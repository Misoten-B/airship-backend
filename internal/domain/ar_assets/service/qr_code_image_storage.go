package service

import (
	"fmt"
	"log"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type QRCodeImageStorage interface {
	Save(qrCodeImage arassets.QRCodeImage) error
	GetImageURL(name string) (string, error)
	GetContainerFullPath() (shared.ContainerFullPath, error)
}

type MockQRCodeImageStorage struct{}

func NewMockQRCodeImageStorage() *MockQRCodeImageStorage {
	return &MockQRCodeImageStorage{}
}

func (s *MockQRCodeImageStorage) Save(_ arassets.QRCodeImage) error {
	log.Println("Mock QRCode Image Storage - Save")
	return nil
}

func (s *MockQRCodeImageStorage) GetImageURL(name string) (string, error) {
	log.Println("Mock QRCode Image Storage - GetURL")

	url := fmt.Sprintf("http://example.com/mock/%s", name)
	return url, nil
}

func (s *MockQRCodeImageStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	return &shared.MockContainerFullPath{}, nil
}
