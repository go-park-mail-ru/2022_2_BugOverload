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

// signupHandler is the structure that handles the request for auth.
type signupHandler struct {
	authService    authService.AuthService
	sessionService sessionService.SessionService
}

// NewSingUpHandler is constructor for signupHandler in this pkg - auth.
func NewSingUpHandler(as authService.AuthService, ss sessionService.SessionService) handler.Handler {
	return &signupHandler{
		as,
		ss,
	}
}

func (h *signupHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/auth/signup", h.Action).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary New user registration
// @Description Sending login and password for registration. Email, password, nickname - required.
// @tags auth, completed
// @Accept json
// @Produce json
// @Param user body models.UserSignupRequest true "Request body for signup"
// @Success 201 {object} models.UserSignupResponse "successfully signup"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchUser "no such user"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/auth/signup [POST]
func (h *signupHandler) Action(w http.ResponseWriter, r *http.Request) {
	signupRequest := models.NewUserSignupRequest()

	err := signupRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(err))
		return
	}

	user, err := h.authService.Signup(r.Context(), signupRequest.GetUser())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		errors.CreateLog(r.Context(), err)
		return
	}

	newSession, err := h.sessionService.CreateSession(r.Context(), &user)
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

	cookieCSRF := &http.Cookie{
		Name:    "CSRF-TOKEN",
		Value:   token,
		Expires: time.Now().Add(pkg.TimeoutLiveCookie),
		Path:    pkg.GlobalCookiePath,
	}

	http.SetCookie(w, cookieCSRF)

	http.SetCookie(w, cookie)

	signupResponse := models.NewUserSignUpResponse(&user)

	httpwrapper.Response(w, http.StatusCreated, signupResponse)
}
