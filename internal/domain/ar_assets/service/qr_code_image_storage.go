package service

import (
	"fmt"
	"log"
	"mime/multipart"
)

type QRCodeImageStorage interface {
	// Save(file file.File) error
	Save(name string, file multipart.File) error
	GetImageURL(name string) (string, error)
}

type MockQRCodeImageStorage struct{}

func (s *MockQRCodeImageStorage) Save(_ string, _ multipart.File) error {
	log.Println("Mock QRCode Image Storage - Save")
	return nil
}

func (s *MockQRCodeImageStorage) GetImageURL(name string) (string, error) {
	log.Println("Mock QRCode Image Storage - GetURL")

	url := fmt.Sprintf("http://example.com/%s", name)
	return url, nil
}
