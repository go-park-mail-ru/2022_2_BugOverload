package httpwrapper

import (
	stdErrors "github.com/pkg/errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errors"
)

type ErrResponse struct {
	ErrMassage string `json:"error,omitempty"`
}

func DefaultHandlerError(w http.ResponseWriter, err error) {
	errResp := ErrResponse{
		err.Error(),
	}

	var errHTTP errors.ErrHTTP
	if ok := stdErrors.As(err, &errHTTP); ok {
		Response(w, errHTTP.Code, errResp)
		return
	}

	var errAuth errors.ErrAuth
	if ok := stdErrors.As(err, &errAuth); ok {
		Response(w, errAuth.Code, errResp)
		return
	}

	if stdErrors.Is(err, errors.ErrNoCookie) {
		Response(w, http.StatusUnauthorized, errResp)
		return
	}

	if stdErrors.Is(err, errors.ErrCookieNotExist) || stdErrors.Is(err, errors.ErrLoginCombinationNotFound) {
		Response(w, http.StatusUnauthorized, errResp)
		return
	}

	if stdErrors.Is(err, errors.ErrFilmNotFound) {
		Response(w, http.StatusNotFound, errResp)
		return
	}

	Response(w, http.StatusBadRequest, errResp)
}
