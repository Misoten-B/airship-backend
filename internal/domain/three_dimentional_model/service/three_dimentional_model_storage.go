package service

import (
	"fmt"
	"log"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type ThreeDimentionalModelStorage interface {
	GetModelURL(modelName string) (string, error)
	GetContainerFullPath() (shared.ContainerFullPath, error)
}

type MockThreeDimentionalModelStorage struct{}

func NewMockThreeDimentionalModelStorage() *MockThreeDimentionalModelStorage {
	return &MockThreeDimentionalModelStorage{}
}

func (s *MockThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	log.Println("Mock 3D Model Storage - GetURL")

	url := fmt.Sprintf("http://example.com/mock/%s", modelName)
	return url, nil
}

func (s *MockThreeDimentionalModelStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	return &shared.MockContainerFullPath{}, nil
}
