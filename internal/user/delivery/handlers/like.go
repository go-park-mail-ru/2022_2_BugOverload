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
type reviewLikeHandler struct {
	filmService serviceFilms.FilmsService
}

// NewReviewLikeHandler is constructor for reviewLikeHandler in this pkg - film.
func NewReviewLikeHandler(fs serviceFilms.FilmsService) handler.Handler {
	return &reviewLikeHandler{
		fs,
	}
}

func (h *reviewLikeHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Review like
// @Description Set like or unset for review by film id and review id. User id get from cookie
// @tags review, not_actual
// @Produce json
// @Param   id        path   int true "film id"
// @Param   review_id query  int true "review id"
// @Success 204 "success"
// @Failure 400 "return error"
// @Failure 401 {object} httpmodels.ErrResponseAuthNoCookie "no cookie"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such film or no such cookie or no such review"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/review/like [POST]
func (h *reviewLikeHandler) Action(w http.ResponseWriter, r *http.Request) {
	// in dev
	//  vars := mux.Vars(r)
}
