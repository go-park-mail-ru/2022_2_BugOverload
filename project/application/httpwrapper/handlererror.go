package httpwrapper

import (
	"errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"net/http"
)

func DefHandlerError(w http.ResponseWriter, err error) {
	if errors.Is(err, errorshandlers.ErrUnsupportedMediaType) {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	http.Error(w, err.Error(), http.StatusBadRequest)
}
