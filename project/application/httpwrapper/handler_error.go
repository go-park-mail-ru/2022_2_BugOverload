package httpwrapper

import (
	"errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
)

//  type ErrResponse struct {
//  }

func DefHandlerError(w http.ResponseWriter, err error) {
	body := `{"error":"` + err.Error() + `"}`

	if errors.Is(err, errorshandlers.ErrUnsupportedMediaType) {
		Response(w, http.StatusUnsupportedMediaType, body)
		return
	}

	if errors.Is(err, errorshandlers.ErrCookieNotExist) {
		Response(w, http.StatusForbidden, body)
		return
	}

	Response(w, http.StatusBadRequest, body)
}
