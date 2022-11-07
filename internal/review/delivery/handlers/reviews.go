package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// reviewLikeHandler is the structure that handles the request for
// getting film by id.
type reviewsHandler struct {
	filmService serviceFilms.FilmsService
}

// NewReviewsHandler is constructor for reviewLikeHandler in this pkg - film.
func NewReviewsHandler(fs serviceFilms.FilmsService) handler.Handler {
	return &reviewsHandler{
		fs,
	}
}

func (h *reviewsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("", h.Action).Methods(http.MethodGet)
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
	// in dev
}
