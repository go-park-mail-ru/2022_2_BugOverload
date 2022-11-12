package httpwrapper

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// DefaultHandlerError is error handler that detects the type of error and gives an error response.
func DefaultHandlerError(ctx context.Context, w http.ResponseWriter, err error) {
	errors.CreateLog(ctx, err)

	errCause := stdErrors.Cause(err)

	errResp := ErrResponse{
		ErrMassage: errCause.Error(),
	}

	code := errors.GetCode(errCause)

	Response(ctx, w, code, errResp)
}
