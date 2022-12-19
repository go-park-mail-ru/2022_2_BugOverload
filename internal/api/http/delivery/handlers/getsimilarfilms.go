package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client"
)

// getSimilarFilmsHandler is the structure that handles the request for similar films
type getSimilarFilmsHandler struct {
	filmService client.WarehouseService
}

// NewGetSimilarFilmsHandler is constructor for getSimilarFilmsHandler in this pkg
func NewGetSimilarFilmsHandler(service client.WarehouseService) handler.Handler {
	return &getSimilarFilmsHandler{
		service,
	}
}

func (h *getSimilarFilmsHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id}/similar", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getting similar films collection
// @Description Return films with the same genres
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Success 200 {object} models.PremieresCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/film/{id}/similar [GET]
func (h *getSimilarFilmsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetSimilarFilmsRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	collection, err := h.filmService.GetSimilarFilms(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewGetSimilarFilmsCollectionResponse(&collection)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
