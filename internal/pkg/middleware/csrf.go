package middleware

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"net/http"
)

func SetCsrfMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Csrf-Token")

		if token == "" {
			httpwrapper.DefaultHandlerError(w, errors.ErrCsrfTokenNotFound)
			return
		}

		cookie, err := r.Cookie("session_id")
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.ErrNoCookie)
			return
		}

		user, err := GetUser

		correctToken, err := security.CheckCsrfToken()
	})
}
