package wrapper

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// DefaultHandlerHTTPError is error handler that detects the type of error and gives an error response.
func DefaultHandlerHTTPError(ctx context.Context, w http.ResponseWriter, err error) {
	errors.CreateLog(ctx, err)

	errCause := stdErrors.Cause(err)

	code, exist := errors.GetErrorCodeHTTP(errCause)
	if !exist {
		errCause = stdErrors.Wrap(errCause, "Undefined error")
	}

	errResp := ErrResponse{
		ErrMassage: errCause.Error(),
	}

	Response(ctx, w, code, errResp)
}
