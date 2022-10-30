package handlers

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
)

// userProfileHandler is the structure that handles the request for auth.
type userProfileHandler struct {
	userService serviceUser.AuthService
	authService serviceAuth.SessionService
}

// NewUserProfileHandler is constructor for userProfileHandler in this pkg - settings.
func NewUserProfileHandler(us serviceUser.AuthService, as serviceAuth.SessionService) pkg.Handler {
	return &userProfileHandler{
		us,
		as,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user stat
// @Description Getting user public info.
// @tags in_dev
// @Produce json
// @Param   id        path   int true "user id"
// @Success 200 {object} models.UserProfileResponse "successfully getting"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchUser "no such user"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/{id} [GET]
func (h *userProfileHandler) Action(w http.ResponseWriter, r *http.Request) {
	settingsRequest := models.NewUserProfileRequest()

	err := settingsRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user, err := h.authService.GetUserBySession(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	authResponse := models.NewUserProfileResponse(&user)

	httpwrapper.Response(w, http.StatusOK, authResponse.ToPublic())
}
