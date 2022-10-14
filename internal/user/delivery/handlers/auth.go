package handlers

import (
	"context"
	stdErrors "github.com/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"time"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// authHandler is the structure that handles the request for auth.
type authHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewAuthHandler is constructor for authHandler in this pkg - auth.
func NewAuthHandler(us serviceUser.UserService, as serviceAuth.AuthService) pkg.Handler {
	return &authHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Defining an authorized user
// @Description Sending login and password
// @tags user
// @Produce json
// @Success 200 {object} models.UserAuthResponse "successfully auth"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "such cookie not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/auth [GET]
func (h *authHandler) Action(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)

	authRequest := models.NewUserAuthRequest()

	err := authRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	ctx := context.WithValue(r.Context(), pkg.CookieKey, cookieStr)

	user, err := h.authService.GetUserBySession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	authResponse := models.NewUserAuthResponse()

	httpwrapper.Response(w, http.StatusOK, authResponse.ToPublic(&user))
}
