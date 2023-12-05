package shared

type Status interface {
	Status() int
	Equal(other Status) bool
}

const (
	statusInProgressCode = 1
)

type StatusInProgress struct{}

func (s StatusInProgress) Status() int {
	return statusInProgressCode
}

func (s StatusInProgress) Equal(other Status) bool {
	return s.Status() == other.Status()
}

type StatusCompleted struct{}

const (
	statusCompletedCode = 1
)

func (s StatusCompleted) Status() int {
	return statusCompletedCode
}

func (s StatusCompleted) Equal(other Status) bool {
	return s.Status() == other.Status()
}
