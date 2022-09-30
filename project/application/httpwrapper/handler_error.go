package httpwrapper

import (
	"errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
)

func DefHandlerError(w http.ResponseWriter, err error) {
	if errors.Is(err, errorshandlers.ErrUnsupportedMediaType) {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	if errors.Is(err, errorshandlers.ErrCookieNotExist) {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.Error(w, err.Error(), http.StatusBadRequest)
}
