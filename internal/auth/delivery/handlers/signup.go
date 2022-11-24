package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	authService "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
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

func (h *signupHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/auth/signup", h.Action).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary New user registration
// @Description Sending login and password for registration. Email, password, nickname - required.
// @tags auth
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
	request := models.NewUserSignupRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	user, err := h.authService.Signup(r.Context(), request.GetUser())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	newSession, err := h.sessionService.CreateSession(r.Context(), &user)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	token, err := security.CreateCsrfToken(&newSession)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
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

	response := models.NewUserSignUpResponse(&user)

	wrapper.Response(r.Context(), w, http.StatusCreated, response)
}
