package service

import (
	"log"
	"mime/multipart"
)

type QRCodeImageStorage interface {
	// Save(file file.File) error
	Save(name string, file multipart.File) error
}

type MockQRCodeImageStorage struct{}

func (s *MockQRCodeImageStorage) Save(_ string, _ multipart.File) error {
	log.Println("Mock QRCode Image Storage - Save")
	return nil
}
