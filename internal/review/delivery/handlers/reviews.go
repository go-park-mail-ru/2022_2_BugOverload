package handlers

import (
	"context"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/review/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/review/service"
)

// reviewLikeHandler is the structure that handles the request for
// getting film by id.
type reviewsHandler struct {
	reviewsService service.ReviewService
}

// NewReviewsHandler is constructor for reviewLikeHandler in this pkg - film.
func NewReviewsHandler(fs service.ReviewService) pkg.Handler {
	return &reviewsHandler{
		fs,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Reviews for film
// @Description Getting reviews by film id
// @tags completed
// @Produce json
// @Param   id        path   int     true "film id"
// @Param   count     query  int     true "count reviews needed"
// @Param   offset    query  int     true "offset count"
// @Success 200 {array} models.ReviewResponse "return reviews"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseFilmNoSuchFilm "no such reviews"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/reviews [GET]
func (h *reviewsHandler) Action(w http.ResponseWriter, r *http.Request) {
	reviewsRequest := models.NewReviewsRequest()

	err := reviewsRequest.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	requestParams := reviewsRequest.GetParams()

	ctx := context.WithValue(r.Context(), pkg.GetReviewsParamsKey, requestParams)

	reviews, err := h.reviewsService.GetReviewsByFilmID(ctx)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrReview(stdErrors.Cause(err)))
		return
	}

	reviewsResponse := models.NewReviewsResponse(&reviews)

	httpwrapper.Response(w, http.StatusOK, reviewsResponse)
}
