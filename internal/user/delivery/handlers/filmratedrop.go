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

// filmRateDropHandler is the structure that handles the request for auth.
type filmRateDropHandler struct {
	userService serviceUser.UserService
}

// NewFilmRateDropHandler is constructor for filmRateDropHandler in this pkg - settings.
func NewFilmRateDropHandler(us serviceUser.UserService) handler.Handler {
	return &filmRateDropHandler{
		us,
	}
}

func (h *filmRateDropHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/film/{id}/rate/drop", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodDelete)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Drop user rate on film
// @Description  Drop user rate on film by filmID
// @tags user
// @Produce json
// @Param   id    path  int    true "film id"
// @Success 200 {object} models.FilmRateDropResponse "successfully drop rate"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/rate/drop [DELETE]
func (h *filmRateDropHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewFilmRateDropRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	film, errService := h.userService.FilmRateDrop(r.Context(), &user, request.GetParams())
	if errService != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errService)
		return
	}

	response := models.NewFilmRateDropResponse(&film)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
