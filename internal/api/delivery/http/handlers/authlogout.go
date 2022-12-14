package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// logoutHandler is the structure that handles the request for auth.
type logoutHandler struct {
	authService client.AuthService
}

// NewLogoutHandler is constructor for logoutHandler in this pkg - auth.
func NewLogoutHandler(as client.AuthService) handler.Handler {
	return &logoutHandler{
		as,
	}
}

func (h *logoutHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/auth/logout", mw.NeedAuthMiddleware(h.Action)).Methods(http.MethodDelete)
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
// @Router /api/v1/auth/logout [DELETE]
func (h *logoutHandler) Action(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(constparams.SessionCookieName)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	requestSession := &mainModels.Session{
		ID: cookie.Value,
	}

	badSession, err := h.authService.DeleteSession(r.Context(), requestSession)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	badCookie := &http.Cookie{
		Name:     constparams.SessionCookieName,
		Value:    badSession.ID,
		Expires:  time.Now().Add(-constparams.TimeoutLiveCookie),
		HttpOnly: true,
	}

	http.SetCookie(w, badCookie)

	wrapper.NoBody(w, http.StatusNoContent)
}
