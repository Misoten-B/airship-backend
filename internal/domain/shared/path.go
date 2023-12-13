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

type FilePath struct {
	path string
}

func NewFilePath(path string) FilePath {
	return FilePath{
		path: path,
	}
}

func (f FilePath) Path() string {
	return f.path
}
