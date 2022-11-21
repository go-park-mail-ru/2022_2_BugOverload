package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/models"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

// getActivityOnFilmHandler is the structure that handles the request for auth.
type getActivityOnFilmHandler struct {
	userService serviceUser.UserService
}

// NewGetActivityOnFilmHandler is constructor for getActivityOnFilmHandler in this pkg - settings.
func NewGetActivityOnFilmHandler(us serviceUser.UserService) handler.Handler {
	return &getActivityOnFilmHandler{
		us,
	}
}

func (h *getActivityOnFilmHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/user_activity", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting user info with film
// @Description Getting user collections, rating on film. Needed auth
// @tags user
// @Produce json
// @Param id  path int true "film id"
// @Success 200 {object} models.GetUserActivityOnFilmResponse "successfully getting"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/user_activity [GET]
func (h *getActivityOnFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		httpwrapper.DefaultHandlerError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewUserActivityOnFilmRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	userActivity, err := h.userService.GetUserActivityOnFilm(r.Context(), &user, request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	response := models.NewGetUserActivityOnFilmResponse(&userActivity)

	httpwrapper.Response(r.Context(), w, http.StatusOK, response)
}
