package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// reviewLikeHandler is the structure that handles the request for
// getting film by id.
type reviewsHandler struct {
	filmService serviceFilms.FilmsService
}

// NewReviewsHandler is constructor for reviewLikeHandler in this pkg - film.
func NewReviewsHandler(fs serviceFilms.FilmsService) pkg.Handler {
	return &reviewsHandler{
		fs,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Review for film
// @Description Getting reviews by film id
// @tags in_dev
// @Produce json
// @Param   id        path   int     true "film id"
// @Param   count     query  int     true "count reviews needed"
// @Param   delimiter query  string  no   "value attribute last review, not needed for first request"
// @Success 200 {array} models.ReviewResponse "return reviews"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such film"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/reviews [GET]
func (h *reviewsHandler) Action(w http.ResponseWriter, r *http.Request) {
	filmRecommendation, err := h.filmService.GerRecommendation(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	response := models.NewFilmResponse(&filmRecommendation)

	httpwrapper.Response(w, http.StatusOK, response.ToPublic())
}
