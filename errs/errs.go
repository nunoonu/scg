package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type Errs struct {
	code    int
	message string
}

func New(httpStatusCode int, msg string) error {

	err := &Errs{
		code:    httpStatusCode,
		message: msg,
	}
	return errors.Wrap(err, msg)

}

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	ctxErrs := ctx.Errors

	for _, ginErr := range ctxErrs {

		errCauseFromService := errors.Cause(ginErr.Err.(error))
		e, ok := errCauseFromService.(*Errs)
		if ok {
			ctx.AbortWithStatusJSON(e.code, e.message)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "default")
		}

		return
	}
}

func (e *Errs) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", e.code, e.message)
}
