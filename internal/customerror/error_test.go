package customerror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Misoten-B/airship-backend/internal/customerror"
)

func TestApplicationError(t *testing.T) {
	err := customerror.NewApplicationErrorWithoutDetails(
		"test error",
		http.StatusBadRequest,
	)

	var appErr *customerror.ApplicationError
	if errors.As(err, &appErr) {
		t.Logf("error message: %s", appErr.Error())
		t.Logf("error status code: %d", appErr.StatusCode())
		t.Logf("error details: %s", appErr.Details())
	} else {
		t.Errorf("error is not *customerror.ApplicationError")
	}
}
