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

	errResp := ErrResponse{
		ErrMassage: stdErrors.Cause(err).Error(),
	}

	code := errors.GetCode(err)

	Response(ctx, w, code, errResp)
}
