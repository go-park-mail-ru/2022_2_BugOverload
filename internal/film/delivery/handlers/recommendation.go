package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// recommendationFilmHandler is the structure that handles the request for
// getting recommendation film for authorized ot unauthorized user.
type recommendationFilmHandler struct {
	filmService serviceFilms.FilmsService
}

// NewRecommendationFilmHandler is constructor for recommendationFilmHandler in this pkg - recommendation film.
func NewRecommendationFilmHandler(fs serviceFilms.FilmsService) handler.Handler {
	return &recommendationFilmHandler{
		fs,
	}
}

func (h *recommendationFilmHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/film/recommendation", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Recommendation film
// @Description Getting a recommended movie for the user
// @tags completed
// @Produce json
// @Success 200 {object} models.RecommendFilmResponse "returns recommendation film movies for auth user or not auth"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/recommendation [GET]
func (h *recommendationFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
	filmRecommendation, err := h.filmService.GetRecommendation(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	response := models.NewRecommendFilmResponse(&filmRecommendation)

	httpwrapper.Response(w, http.StatusOK, response)
}
