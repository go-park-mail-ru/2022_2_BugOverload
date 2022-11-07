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

// putSettingsHandler is the structure that handles the request for auth.
type putSettingsHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewPutSettingsHandler is constructor for putSettingsHandler in this pkg - settings.
func NewPutSettingsHandler(us serviceUser.AuthService, as serviceAuth.SessionService) handler.Handler {
	return &putSettingsHandler{
		us,
		as,
	}
}

func (h *putSettingsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("", mw.SetCsrfMiddleware(mw.CheckAuthMiddleware(h.Action))).Methods(http.MethodPut)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Change user auth data
// @Description Request for change user settings and data. Needed auth
// @tags in_dev
// @Produce json
// @Param user body models.UserPutSettingsRequest true "Request body for change user data"
// @Success 204 "successfully changes"
// @Failure 400 {object} httpmodels.ErrResponseAuthDefault "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 403 {object} httpmodels.ErrResponseAuthWrongLoginCombination "wrong pass"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/settings [PUT]
func (h *putSettingsHandler) Action(w http.ResponseWriter, r *http.Request) {
	getUserSettingsRequest := models.NewPutUserSettingsRequest()

	err := getUserSettingsRequest.Bind(r)
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

	getUserSettingsResponse := models.NewGetUserSettingsResponse(&user)

	httpwrapper.Response(w, http.StatusOK, getUserSettingsResponse)
}
