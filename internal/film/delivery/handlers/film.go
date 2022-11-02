package handlers

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// filmHandler is the structure that handles the request for
// getting film by id.
type filmHandler struct {
	filmService serviceFilms.FilmsService
}

// NewFilmHandler is constructor for filmHandler in this pkg - film.
func NewFilmHandler(fs serviceFilms.FilmsService) pkg.Handler {
	return &filmHandler{
		fs,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Film full info
// @Description Getting film info by id
// @tags in_dev
// @Produce json
// @Param id  path int true "film id"
// @Success 200 {object} models.FilmResponse "return film"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such film"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id} [GET]
func (h *filmHandler) Action(w http.ResponseWriter, r *http.Request) {
	film, err := h.filmService.GerRecommendation(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	filmResponse := models.NewFilmResponse(&film)

	httpwrapper.Response(w, http.StatusOK, filmResponse)
}
