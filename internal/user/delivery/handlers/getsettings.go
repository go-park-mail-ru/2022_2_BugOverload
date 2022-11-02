package handlers

import (
	"net/http"

	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

// getSettingsHandler is the structure that handles the request for auth.
type getSettingsHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewGetSettingsHandler is constructor for getSettingsHandler in this pkg - settings.
func NewGetSettingsHandler(us serviceUser.AuthService, as serviceAuth.SessionService) pkg.Handler {
	return &getSettingsHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user stat and info
// @Description Getting user info and info for changes. Needed auth
// @tags in_dev
// @Produce json
// @Success 200 {object} models.GetUserSettingsResponse "successfully getting"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/settings [GET]
func (h *getSettingsHandler) Action(w http.ResponseWriter, r *http.Request) {
	// in dev
}