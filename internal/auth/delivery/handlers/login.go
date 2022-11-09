package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	authService "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	sessionService "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// loginHandler is the structure that handles the request for auth.
type loginHandler struct {
	authService    authService.AuthService
	sessionService sessionService.SessionService
}

// NewLoginHandler is constructor for loginHandler in this pkg - auth.
func NewLoginHandler(as authService.AuthService, ss sessionService.SessionService) handler.Handler {
	return &loginHandler{
		as,
		ss,
	}
}

func (h *loginHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/auth/login", h.Action).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary User authentication
// @Description Sending login and password. Email and password - required.
// @tags auth, completed
// @Accept json
// @Produce json
// @Param user body models.UserLoginRequest true "Request body for login"
// @Success 200 {object} models.UserLoginResponse "successfully login"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 403 {object} httpmodels.ErrResponseAuthWrongLoginCombination "wrong combination"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchUser "no such user"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/auth/login [POST]
func (h *loginHandler) Action(w http.ResponseWriter, r *http.Request) {
	loginRequest := models.NewUserLoginRequest()

	err := loginRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	userLogged, err := h.authService.Login(r.Context(), loginRequest.GetUser())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	newSession, err := h.sessionService.CreateSession(r.Context(), &userLogged)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	token, err := security.CreateCsrfToken(&newSession)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	w.Header().Set("X-CSRF-TOKEN", token)

	cookie := &http.Cookie{
		Name:     pkg.SessionCookieName,
		Value:    newSession.ID,
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		Path:     pkg.GlobalCookiePath,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	loginResponse := models.NewUserLoginResponse(&userLogged)

	httpwrapper.Response(w, http.StatusOK, loginResponse)
}
