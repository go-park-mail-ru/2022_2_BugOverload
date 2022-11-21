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

// getStdCollectionHandler is the structure that handles the request for movies in cinemas.
type getStdCollectionHandler struct {
	collectionService service.CollectionService
}

// NewStdCollectionHandler is constructor for getStdCollectionHandler in this pkg - in cinema.
func NewStdCollectionHandler(uc service.CollectionService) handler.Handler {
	return &getStdCollectionHandler{
		uc,
	}
}

func (h *getStdCollectionHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
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
// @Summary Films by genre or tag by DESK rating or DESK date
// @Description Films by tag "популярное" or "сейчас в кино". Key, count_films, delimiter - required.
// @tags collection
// @Produce json
// @Param target      query string true "genre, tag, prod_company, prod_country"
// @Param key         query string true "for genre - comedy, tag, prod_company, prod_country"
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
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetStdCollection(r.Context(), request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	response := models.NewStdCollectionResponse(&collection)

	httpwrapper.Response(r.Context(), w, http.StatusOK, response)
}
