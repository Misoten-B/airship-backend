package shared

type Status interface {
	Status() int
	Equal(other Status) bool
}

const (
	statusErrorCode = -1
)

type StatusError struct{}

func (s StatusError) Status() int {
	return statusErrorCode
}

func (s StatusError) Equal(other Status) bool {
	return s.Status() == other.Status()
}

const (
	statusNoneCode = 1
)

type StatusNone struct{}

func (s StatusNone) Status() int {
	return statusNoneCode
}

func (s StatusNone) Equal(other Status) bool {
	return s.Status() == other.Status()
}

const (
	statusInProgressCode = 2
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
	statusCompletedCode = 3
)

func (s StatusCompleted) Status() int {
	return statusCompletedCode
}

func (s StatusCompleted) Equal(other Status) bool {
	return s.Status() == other.Status()
}
