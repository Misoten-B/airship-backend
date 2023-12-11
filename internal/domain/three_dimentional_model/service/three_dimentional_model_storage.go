package service

import (
	"fmt"
	"log"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
)

type ThreeDimentionalModelStorage interface {
	Save(threeDimentionalModel threedimentionalmodel.ThreeDimentionalModel) error
	GetModelURL(modelName string) (string, error)
	GetContainerFullPath() (shared.ContainerFullPath, error)
}

type MockThreeDimentionalModelStorage struct{}

func NewMockThreeDimentionalModelStorage() *MockThreeDimentionalModelStorage {
	return &MockThreeDimentionalModelStorage{}
}

func (s *MockThreeDimentionalModelStorage) Save(_ threedimentionalmodel.ThreeDimentionalModel) error {
	log.Println("Mock 3D Model Storage - Save")

	return nil
}

func (s *MockThreeDimentionalModelStorage) GetModelURL(modelName string) (string, error) {
	log.Println("Mock 3D Model Storage - GetURL")

	url := fmt.Sprintf("http://example.com/mock/%s", modelName)
	return url, nil
}

func (s *MockThreeDimentionalModelStorage) GetContainerFullPath() (shared.ContainerFullPath, error) {
	return &shared.MockContainerFullPath{}, nil
}
