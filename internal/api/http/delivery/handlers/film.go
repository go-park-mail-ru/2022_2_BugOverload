package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// filmHandler is the structure that handles the request for
// getting film by id.
type filmHandler struct {
	filmService serviceFilms.FilmsService
}

// NewFilmHandler is constructor for filmHandler in this pkg - film.
func NewFilmHandler(fs serviceFilms.FilmsService) handler.Handler {
	return &filmHandler{
		fs,
	}
}

func (h *filmHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}", h.Action).
		Methods(http.MethodGet).
		Queries("count_images", "{count_images}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Film full info
// @Description Getting film info by id. ID, count_images - required.
// @tags film
// @Produce json
// @Param id  path int true "film id"
// @Param count_images  query int true "count images film"
// @Success 200 {object} models.FilmResponse "return film"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such film"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id} [GET]
func (h *filmHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewFilmRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	film, err := h.filmService.GetFilmByID(r.Context(), request.GetFilm(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewFilmResponse(&film)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
