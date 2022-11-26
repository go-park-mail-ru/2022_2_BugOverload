package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	serviceUserProfile "go-park-mail-ru/2022_2_BugOverload/internal/user/user/service"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// userProfileHandler is the structure that handles the request for auth.
type userProfileHandler struct {
	userProfileService serviceUserProfile.UserService
}

// NewUserProfileHandler is constructor for userProfileHandler in this pkg - settings.
func NewUserProfileHandler(us serviceUserProfile.UserService) handler.Handler {
	return &userProfileHandler{
		us,
	}
}

func (h *userProfileHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/user/profile/{id:[0-9]+}", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user stat
// @Description Getting user public info.
// @tags user
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
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	user, err := h.userProfileService.GetUserProfileByID(r.Context(), request.GetUser())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewUserProfileResponse(&user)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
