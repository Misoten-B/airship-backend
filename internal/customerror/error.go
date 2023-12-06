package customerror

type ApplicationError struct {
	message    string
	statusCode int
	// details は開発者向けのエラー詳細を格納します。
	details string
}

func NewApplicationError(message string, statusCode int, details string) error {
	return &ApplicationError{
		message:    message,
		statusCode: statusCode,
		details:    details,
	}
}

func (e *ApplicationError) Error() string {
	return e.message
}

func (e *ApplicationError) StatusCode() int {
	return e.statusCode
}

// Details は開発者向けのエラー詳細を返します。
func (e *ApplicationError) Details() string {
	return e.details
}
