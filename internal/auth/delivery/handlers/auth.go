package handlers

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// authHandler is the structure that handles the request for auth.
type authHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewAuthHandler is constructor for authHandler in this pkg - auth.
func NewAuthHandler(us serviceUser.AuthService, as serviceAuth.SessionService) pkg.Handler {
	return &authHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Defining an authorized user
// @Description Sending login and password
// @tags auth
// @Produce json
// @Success 200 {object} models.UserAuthResponse "successfully auth"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/auth [GET]
func (h *authHandler) Action(w http.ResponseWriter, r *http.Request) {
	authRequest := models.NewUserAuthRequest()

	ctx, err := authRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user, err := h.authService.GetUserBySession(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	authResponse := models.NewUserAuthResponse(&user)

	httpwrapper.Response(w, http.StatusOK, authResponse.ToPublic())
}
