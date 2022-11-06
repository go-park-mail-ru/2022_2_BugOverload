package handlers

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUserProfile "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// userProfileHandler is the structure that handles the request for auth.
type userProfileHandler struct {
	userProfileService serviceUserProfile.UserService
}

// NewUserProfileHandler is constructor for userProfileHandler in this pkg - settings.
func NewUserProfileHandler(us serviceUserProfile.UserService) pkg.Handler {
	return &userProfileHandler{
		us,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user stat
// @Description Getting user public info.
// @tags completed
// @Produce json
// @Param   id        path   int true "user id"
// @Success 200 {object} models.UserProfileResponse "successfully getting"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseProfileNoSuchProfile "no such profile"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/profile/{id} [GET]
func (h *userProfileHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewUserProfileRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user, err := h.userProfileService.GetUserProfileByID(r.Context(), request.GetUser())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrProfile(stdErrors.Cause(err)))
		return
	}

	response := models.NewUserProfileResponse(&user)

	httpwrapper.Response(w, http.StatusOK, response)
}
