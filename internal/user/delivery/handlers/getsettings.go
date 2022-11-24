package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// getSettingsHandler is the structure that handles the request for auth.
type getSettingsHandler struct {
	userService serviceUser.UserService
}

// NewGetSettingsHandler is constructor for getActivityOnFilmHandler in this pkg - settings.
func NewGetSettingsHandler(us serviceUser.UserService) handler.Handler {
	return &getSettingsHandler{
		us,
	}
}

func (h *getSettingsHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/user/settings", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user stat and info
// @Description Getting user info and info for changes. Needed auth
// @tags user
// @Produce json
// @Success 200 {object} models.GetUserSettingsResponse "successfully getting"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/settings [GET]
func (h *getSettingsHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	userProfile, err := h.userService.GetUserProfileSettings(r.Context(), &user)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewGetUserSettingsResponse(&userProfile)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
