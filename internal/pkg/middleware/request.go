package middleware

import (
	"net/http"
	"strconv"

	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

type RequestMiddleware struct{}

func NewRequestMiddleware() RequestMiddleware {
	return RequestMiddleware{}
}

func (umd *RequestMiddleware) SetSizeRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strLength := r.Header.Get("Content-Length")
		length, err := strconv.Atoi(strLength)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrConvertLength))
			return
		}

		if length > innerPKG.BufSizeRequest {
			httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(errors.ErrBigRequest))
			return
		}

		h.ServeHTTP(w, r)
	})
}
