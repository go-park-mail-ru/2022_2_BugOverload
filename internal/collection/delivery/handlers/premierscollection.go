package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// premiersCollectionHandler is the structure that handles the request for movies in cinemas.
type premiersCollectionHandler struct {
	collectionService service.CollectionService
}

// NewPremiersCollectionHandler is constructor for getStdCollectionHandler in this pkg - in cinema.
func NewPremiersCollectionHandler(uc service.CollectionService) handler.Handler {
	return &premiersCollectionHandler{
		uc,
	}
}

func (h *premiersCollectionHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/collection/premiers", h.Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getter premiers collection
// @Description Return films with production date > current date.
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param count_films query int    true "count films"
// @Success 200 {object} models.PremiersCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection/premiers [GET]
func (h *premiersCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewPremiersCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetPremiersCollection(r.Context(), request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	response := models.NewPremiersCollectionResponse(&collection)

	httpwrapper.Response(r.Context(), w, http.StatusOK, response)
}
