package handlers

import (
	"context"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/params"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// logoutHandler is the structure that handles the request for auth.
type logoutHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewLogoutHandler is constructor for logoutHandler in this pkg - auth.
func NewLogoutHandler(us serviceUser.UserService, as serviceAuth.AuthService) interfaces.Handler {
	return &logoutHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *logoutHandler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(r)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), params.CookieKey, cookieStr)

	badCookie, err := h.authService.DeleteSession(ctx)
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper2.NoContent(w)
}
