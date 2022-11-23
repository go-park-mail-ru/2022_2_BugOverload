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

// getUserCollectionsHandler is the structure that handles the request for movies in cinemas.
type getUserCollectionsHandler struct {
	collectionService service.CollectionService
}

// NewGetUserCollectionsHandler is constructor for getUserCollectionsHandler in this pkg - in cinema.
func NewGetUserCollectionsHandler(uc service.CollectionService) handler.Handler {
	return &getUserCollectionsHandler{
		uc,
	}
}

func (h *getUserCollectionsHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/user/collections", h.Action).
		Methods(http.MethodGet).
		Queries(
			"sort_param", "{sort_param}",
			"count_collections", "{count_films}",
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
func (h *getUserCollectionsHandler) Action(w http.ResponseWriter, r *http.Request) {
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