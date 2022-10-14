package httpwrapper

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// DefaultHandlerError is error handler that detects the type of error and gives an error response.
func DefaultHandlerError(w http.ResponseWriter, err error) {
	errResp := ErrResponse{
		err.Error(),
	}

	var errHTTP errors.DefaultValidationError
	if ok := stdErrors.As(err, &errHTTP); ok {
		Response(w, errHTTP.Code, errResp)
		return
	}

	var errAuth errors.AuthError
	if ok := stdErrors.As(err, &errAuth); ok {
		Response(w, errAuth.Code, errResp)
		return
	}

	var errFilms errors.FilmsError
	if ok := stdErrors.As(err, &errFilms); ok {
		Response(w, errFilms.Code, errResp)
		return
	}
}