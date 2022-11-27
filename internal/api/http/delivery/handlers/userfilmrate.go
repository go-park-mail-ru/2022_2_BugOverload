package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"net/http"

	"github.com/gorilla/mux"

	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// filmRateHandler is the structure that handles the request for auth.
type filmRateHandler struct {
	userService serviceUser.UserService
}

// NewFilmRateHandler is constructor for filmRateHandler in this pkg - settings.
func NewFilmRateHandler(us serviceUser.UserService) handler.Handler {
	return &filmRateHandler{
		us,
	}
}

func (h *filmRateHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id}/rate", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).Methods(http.MethodPost)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Set user rate on film
// @Description  Set user rate on film by filmID. Score is required and 10 >= score >= 0.
// @tags user
// @Produce json
// @Param   id    path  int    true "film id"
// @Param score body models.FilmRateRequest true "Request body for rate film"
// @Success 200 {object} models.FilmRateResponse "successfully rate"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseAuthNoSuchCookie "no such cookie"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/rate [POST]
func (h *filmRateHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewFilmRateRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	film, errService := h.userService.FilmRate(r.Context(), &user, request.GetParams())
	if errService != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errService)
		return
	}

	response := models.NewFilmRateResponse(&film)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
