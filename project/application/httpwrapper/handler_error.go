package httpwrapper

import (
	"errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
)

type ErrResponse struct {
	ErrMassage string `json:"error,omitempty"`
}

func DefHandlerError(w http.ResponseWriter, err error) {
	errResp := ErrResponse{
		err.Error(),
	}

	if errors.Is(err, errorshandlers.ErrNoCookie) {
		Response(w, http.StatusUnauthorized, errResp)
		return
	}

	if errors.Is(err, errorshandlers.ErrUnsupportedMediaType) {
		Response(w, http.StatusUnsupportedMediaType, errResp)
		return
	}

	if errors.Is(err, errorshandlers.ErrCookieNotExist) || errors.Is(err, errorshandlers.ErrLoginCombinationNotFound) {
		Response(w, http.StatusUnauthorized, errResp)
		return
	}

	if errors.Is(err, errorshandlers.ErrFilmNotFound) {
		Response(w, http.StatusNotFound, errResp)
		return
	}

	Response(w, http.StatusBadRequest, errResp)
}
