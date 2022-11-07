package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// authHandler is the structure that handles the request for auth.
type authHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewAuthHandler is constructor for authHandler in this pkg - auth.
func NewAuthHandler(us serviceUser.AuthService, as serviceAuth.SessionService) handler.Handler {
	return &authHandler{
		us,
		as,
	}
}

func (h *authHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/auth", h.Action).Methods(http.MethodGet)
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

	err := authRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	requestSession := mainModels.Session{
		ID: r.Cookies()[0].Value,
	}

	user, err := h.authService.GetUserBySession(r.Context(), requestSession)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	authResponse := models.NewUserAuthResponse(&user)

	httpwrapper.Response(w, http.StatusOK, authResponse)
}
