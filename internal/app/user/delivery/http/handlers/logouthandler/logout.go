package logouthandler

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// handler is structure for API auth, login and signup processing
type handler struct {
	userService userInterface.UserService
	authService authInterface.AuthService
}

// NewHandler is constructor for handler
func NewHandler(us userInterface.UserService, as authInterface.AuthService) *handler {
	return &handler{
		us,
		as,
	}
}

// Action is handling request for check current client cookie and return user data
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), "cookie", cookieStr)

	badCookie, err := h.authService.DeleteSession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper.NoContent(w)
}
