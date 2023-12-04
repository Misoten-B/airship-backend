package service

import (
	"fmt"
	"log"
)

type ThreeDimentionalModelStorage interface {
	GetModelURL(modelName string) (string, error)
}

type MockThreeDimentionalModelStorage struct{}

func (s *MockThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	log.Println("Mock 3D Model Storage - GetURL")

	url := fmt.Sprintf("http://example.com/%s", modelName)
	return url, nil
}
