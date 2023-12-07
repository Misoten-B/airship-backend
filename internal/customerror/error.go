package customerror

import (
	"fmt"
	"net/http"
)

type ApplicationError struct {
	message    string
	statusCode int
	// details は開発者向けのエラー詳細を格納します。
	details string
}

func NewApplicationError(err error, message string, statusCode int) error {
	details := fmt.Sprintf("%s: %s", message, err.Error())

	return &ApplicationError{
		message:    message,
		statusCode: statusCode,
		details:    details,
	}
}

func NewApplicationErrorWithoutDetails(message string, statusCode int) error {
	return &ApplicationError{
		message:    message,
		statusCode: statusCode,
		details:    message,
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

func (e *ApplicationError) IsInternalError() bool {
	return e.statusCode == http.StatusInternalServerError
}
