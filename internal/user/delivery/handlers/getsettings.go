package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
)

// getSettingsHandler is the structure that handles the request for auth.
type getSettingsHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewGetSettingsHandler is constructor for getSettingsHandler in this pkg - settings.
func NewGetSettingsHandler(us serviceUser.AuthService, as serviceAuth.SessionService) handler.Handler {
	return &getSettingsHandler{
		us,
		as,
	}
}

func (h *getSettingsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("", h.Action).Methods(http.MethodGet)
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
	request := models.NewGetUserSettingsRequest()

	err := request.Bind(r)
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

	response := models.NewGetUserSettingsResponse(&user)

	httpwrapper.Response(w, http.StatusOK, response)
}
