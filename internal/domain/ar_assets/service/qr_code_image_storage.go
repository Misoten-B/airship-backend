package service

import (
	"fmt"
	"log"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
)

type QRCodeImageStorage interface {
	Save(qrCodeImage arassets.QRCodeImage) error
	GetImageURL(qrCodeImage arassets.QRCodeImage) (string, error)
}

type MockQRCodeImageStorage struct{}

func (s *MockQRCodeImageStorage) Save(_ arassets.QRCodeImage) error {
	log.Println("Mock QRCode Image Storage - Save")
	return nil
}

func (s *MockQRCodeImageStorage) GetImageURL(qrCodeImage arassets.QRCodeImage) (string, error) {
	log.Println("Mock QRCode Image Storage - GetURL")

	url := fmt.Sprintf("http://example.com/%s", qrCodeImage.Name())
	return url, nil
}
