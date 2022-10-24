package handlers

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"time"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// logoutHandler is the structure that handles the request for auth.
type logoutHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewLogoutHandler is constructor for logoutHandler in this pkg - auth.
func NewLogoutHandler(us serviceUser.UserService, as serviceAuth.AuthService) pkg.Handler {
	return &logoutHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary User logout
// @Description Session delete
// @tags user
// @Success 204 "successfully logout"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "such cookie not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/auth/logout [GET]
func (h *logoutHandler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), pkg.SessionKey, cookieStr)

	badSession, err := h.authService.DeleteSession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    badSession,
		Expires:  time.Now().Add(-pkg.TimeoutLiveCookie),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	httpwrapper.NoContent(w)
}
