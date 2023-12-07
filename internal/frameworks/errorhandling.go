package frameworks

import (
	"errors"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/gin-gonic/gin"
)

// ErrorHandling はGinのHTTPハンドラ用のエラーハンドリングを行います。
// 内部エラーの場合はログを出力します。
func ErrorHandling(c *gin.Context, err error, statusCode int) {
	if err == nil {
		return
	}

	if statusCode == http.StatusInternalServerError {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			log.Printf("%s", appErr.Details())
		} else {
			log.Printf("%s", err.Error())
		}
	}

	c.JSON(statusCode, gin.H{"error": err.Error()})
}
