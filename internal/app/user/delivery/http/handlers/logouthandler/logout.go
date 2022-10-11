package logouthandler

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
)

type handler struct {
	userService userInterface.UserService
	authService authInterface.AuthService
}

func NewHandler(us userInterface.UserService, as authInterface.AuthService) interfaces.Handler {
	return &handler{
		us,
		as,
	}
}

func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), params.CookieKey, cookieStr)

	badCookie, err := h.authService.DeleteSession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper.NoContent(w)
}
