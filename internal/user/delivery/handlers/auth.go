package handlers

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/params"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// authHandler is the structure that handles the request for auth.
type authHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewAuthHandler is constructor for authHandler in this pkg - auth.
func NewAuthHandler(us serviceUser.UserService, as serviceAuth.AuthService) interfaces.Handler {
	return &authHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *authHandler) Action(w http.ResponseWriter, r *http.Request) {
	authRequest := models.NewUserAuthRequest()

	err := authRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), params.CookieKey, cookieStr)

	user, err := h.authService.GetUserBySession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	httpwrapper.Response(w, http.StatusOK, authRequest.ToPublic(&user))
}
