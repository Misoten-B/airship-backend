package shared

type ContainerFullPath interface {
	Path(name string) string
}

type MockContainerFullPath struct{}

func NewMockContainerFullPath() *MockContainerFullPath {
	return &MockContainerFullPath{}
}

func (m *MockContainerFullPath) Path(name string) string {
	return "http://example.com/mock/" + name
}
