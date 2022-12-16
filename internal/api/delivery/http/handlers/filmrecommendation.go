package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// recommendationFilmHandler is the structure that handles the request for
// getting recommendation film for authorized ot unauthorized user.
type recommendationFilmHandler struct {
	filmService client.WarehouseService
}

// NewRecommendationFilmHandler is constructor for recommendationFilmHandler in this pkg - recommendation film.
func NewRecommendationFilmHandler(service client.WarehouseService) handler.Handler {
	return &recommendationFilmHandler{
		service,
	}
}

func (h *recommendationFilmHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/recommendation", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Recommendation film
// @Description Getting a recommended movie for the user
// @tags film
// @Produce json
// @Success 200 {object} models.RecommendFilmResponse "returns recommendation film movies for auth user or not auth"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/recommendation [GET]
func (h *recommendationFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
	filmRecommendation, err := h.filmService.GetRecommendation(r.Context())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewRecommendFilmResponse(&filmRecommendation)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
