package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// loginHandler is the structure that handles the request for auth.
type loginHandler struct {
	userService serviceUser.UserService
	authService serviceAuth.AuthService
}

// NewLoginHandler is constructor for loginHandler in this pkg - auth.
func NewLoginHandler(us serviceUser.UserService, as serviceAuth.AuthService) pkg.Handler {
	return &loginHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary User authentication
// @Description Sending login and password
// @tags user
// @Accept json
// @Produce json
// @Param user body models.UserLoginRequest true "Request body for login"
// @Success 200 {object} models.UserLoginResponse "successfully login"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchUser "such user not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/auth/login [POST]
func (h *loginHandler) Action(w http.ResponseWriter, r *http.Request) {
	loginRequest := models.NewUserLoginRequest()

	err := loginRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userLogged, err := h.userService.Login(r.Context(), user)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	newSession, err := h.authService.CreateSession(r.Context(), &userLogged)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("Set-Cookie", newSession)

	loginResponse := models.NewUserLoginResponse()

	httpwrapper.Response(w, http.StatusOK, loginResponse.ToPublic(&userLogged))
}
