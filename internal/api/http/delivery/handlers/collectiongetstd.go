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

// getStdCollectionHandler is the structure that handles the request for movies in cinemas.
type getStdCollectionHandler struct {
	collectionService client.WarehouseService
}

// NewStdCollectionHandler is constructor for getStdCollectionHandler in this pkg - in cinema.
func NewStdCollectionHandler(service client.WarehouseService) handler.Handler {
	return &getStdCollectionHandler{
		service,
	}
}

func (h *getStdCollectionHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/collection", h.Action).
		Methods(http.MethodGet).
		Queries(
			"target", "{target}",
			"key", "{key}",
			"sort_param", "{sort_param}",
			"count_films", "{count_films}",
			"delimiter", "{delimiter}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getter std collection (tag, genre)
// @Description Films by genre or tag by DESC rating or DESC date. prod_company, prod_country also in the future.
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param target      query string true "genre, tag"
// @Param key         query string true "for genre - comedy, tag - popular"
// @Param sort_param  query string true "rating, date"
// @Param count_films query int    true "count films"
// @Param delimiter   query string true "last value while in is rating last returned film for rating OR offset for date"
// @Success 200 {object} models.GetStdCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseCollectionNoSuchCollection "no such collection"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection [GET]
func (h *getStdCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetStdCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetStdCollection(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewStdCollectionResponse(&collection)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
