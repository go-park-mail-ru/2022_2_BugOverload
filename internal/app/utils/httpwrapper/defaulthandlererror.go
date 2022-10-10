package httpwrapper

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	stdErrors "github.com/pkg/errors"
)

type ErrResponse struct {
	ErrMassage string `json:"error,omitempty"`
}

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

	//  Response(w, http.StatusBadRequest, errResp)
}
