package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	authService "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	sessionService "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// authHandler is the structure that handles the request for auth.
type authHandler struct {
	authService    authService.AuthService
	sessionService sessionService.SessionService
}

// NewAuthHandler is constructor for authHandler in this pkg - auth.
func NewAuthHandler(as authService.AuthService, ss sessionService.SessionService) handler.Handler {
	return &authHandler{
		as,
		ss,
	}
}

func (h *authHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/auth", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Defining an authorized user
// @Description Sending login and password
// @tags completed
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
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
		return
	}

	cookie, err := r.Cookie(pkg.SessionCookieName)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrSessionNotExist))
		return
	}

	requestSession := mainModels.Session{
		ID: cookie.Value,
	}

	user, err := h.sessionService.GetUserBySession(r.Context(), requestSession)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}
	requestSession.User = &user

	token, err := security.CreateCsrfToken(&requestSession)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("X-CSRF-TOKEN", token)

	userAuth, err := h.authService.Auth(r.Context(), &user)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	authResponse := models.NewUserAuthResponse(&userAuth)

	httpwrapper.Response(w, http.StatusOK, authResponse)
}
