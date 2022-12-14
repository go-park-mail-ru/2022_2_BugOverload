package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// loginHandler is the structure that handles the request for auth.
type loginHandler struct {
	authService client.AuthService
}

// NewLoginHandler is constructor for loginHandler in this pkg - auth.
func NewLoginHandler(as client.AuthService) handler.Handler {
	return &loginHandler{
		as,
	}
}

func (h *loginHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/auth/login", h.Action).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary User authentication
// @Description Sending login and password. Email and password - required.
// @tags auth
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
	request := models.NewUserLoginRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	userLogged, err := h.authService.Login(r.Context(), request.GetUser())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	newSession, err := h.authService.CreateSession(r.Context(), &userLogged)
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
		Name:     constparams.SessionCookieName,
		Value:    newSession.ID,
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	cookieCSRF := &http.Cookie{
		Name:    "CSRF-TOKEN",
		Value:   token,
		Expires: time.Now().Add(constparams.TimeoutLiveCookie),
		Path:    constparams.GlobalCookiePath,
	}

	http.SetCookie(w, cookieCSRF)

	http.SetCookie(w, cookie)

	response := models.NewUserLoginResponse(&userLogged)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
