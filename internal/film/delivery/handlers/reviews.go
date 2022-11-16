package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// reviewLikeHandler is the structure that handles the request for
// getting film by id.
type reviewsHandler struct {
	reviewsService service.FilmsService
}

// NewReviewsHandler is constructor for reviewLikeHandler in this pkg - film.
func NewReviewsHandler(fs service.FilmsService) handler.Handler {
	return &reviewsHandler{
		fs,
	}
}

func (h *reviewsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/reviews", h.Action).
		Methods(http.MethodGet).
		Queries("count_reviews", "{count_reviews}", "offset", "{offset}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Reviews for film
// @Description Getting reviews by film id. Count_reviews required and >= 0.
// @tags review, completed
// @Produce json
// @Param   id        path   int     true "film id"
// @Param   count_reviews     query  int     true "count reviews needed"
// @Param   offset    query  int     true "offset count"
// @Success 200 {array} models.ReviewResponse "return reviews"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such reviews"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/reviews [GET]
func (h *reviewsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewReviewsRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	reviews, err := h.reviewsService.GetReviewsByFilmID(r.Context(), request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	response := models.NewReviewsResponse(&reviews)

	httpwrapper.Response(r.Context(), w, http.StatusOK, response)
}
