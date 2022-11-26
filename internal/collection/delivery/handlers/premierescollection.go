package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// premieresCollectionHandler is the structure that handles the request for movies in cinemas.
type premieresCollectionHandler struct {
	collectionService service.CollectionService
}

// NewPremieresCollectionHandler is constructor for getStdCollectionHandler in this pkg - in cinema.
func NewPremieresCollectionHandler(uc service.CollectionService) handler.Handler {
	return &premieresCollectionHandler{
		uc,
	}
}

func (h *premieresCollectionHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/premieres", h.Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}", "delimiter", "{delimiter}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getter premiers collection
// @Description Return films with production date > current date.
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param count_films query int    true "count films"
// @Param delimiter   query string true "offset"
// @Success 200 {object} models.PremieresCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/premieres [GET]
func (h *premieresCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPremieresCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetPremieresCollection(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewPremieresCollectionResponse(&collection)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
