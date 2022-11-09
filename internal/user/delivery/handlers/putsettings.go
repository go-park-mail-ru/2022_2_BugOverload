package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// putSettingsHandler is the structure that handles the request for auth.
type putSettingsHandler struct {
	userService serviceUser.UserService
}

// NewPutSettingsHandler is constructor for putSettingsHandler in this pkg - settings.
func NewPutSettingsHandler(us serviceUser.UserService) handler.Handler {
	return &putSettingsHandler{
		us,
	}
}

func (h *putSettingsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/user/settings", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodPut)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Change user auth data
// @Description Request for change user settings and data. Needed auth.
// @Description Nickname required other not or new_password and cur_password required other not.
// @tags user, completed
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
	request := models.NewPutUserSettingsRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(stdErrors.Cause(err)))
		return
	}

	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrGetUserRequest))
		return
	}

	err = h.userService.ChangeUserProfileSettings(r.Context(), &user, request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrProfile(stdErrors.Cause(err)))
		errors.CreateLog(r.Context(), err)
		return
	}

	httpwrapper.NoBody(w, http.StatusNoContent)
}
