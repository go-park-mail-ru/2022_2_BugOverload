package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// logoutHandler is the structure that handles the request for auth.
type logoutHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewLogoutHandler is constructor for logoutHandler in this pkg - auth.
func NewLogoutHandler(us serviceUser.AuthService, as serviceAuth.SessionService) handler.Handler {
	return &logoutHandler{
		us,
		as,
	}
}

func (h *logoutHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/auth/logout", mw.CheckAuthMiddleware(h.Action)).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary User logout
// @Description Session delete. Needed auth
// @tags auth
// @Success 204 "successfully logout"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/auth/logout [GET]
func (h *logoutHandler) Action(w http.ResponseWriter, r *http.Request) {
	var logoutRequest models.UserLogoutRequest

	err := logoutRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
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

	badSession, err := h.authService.DeleteSession(r.Context(), requestSession)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	badCookie := &http.Cookie{
		Name:     pkg.SessionCookieName,
		Value:    badSession.ID,
		Expires:  time.Now().Add(-pkg.TimeoutLiveCookie),
		HttpOnly: true,
	}

	http.SetCookie(w, badCookie)

	httpwrapper.NoBody(w, http.StatusNoContent)
}
