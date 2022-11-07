package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

// tagCollectionHandler is the structure that handles the request for movies in cinemas.
type tagCollectionHandler struct {
	collectionService service.CollectionService
}

// NewTagCollectionHandler is constructor for tagCollectionHandler in this pkg - in cinema.
func NewTagCollectionHandler(uc service.CollectionService) handler.Handler {
	return &tagCollectionHandler{
		uc,
	}
}

func (h *tagCollectionHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/collection/{tag}", h.Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}", "delimiter", "{delimiter}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Films by tag
// @Description Films by tag "популярное" or "сейчас в кино"
// @tags completed
// @Produce json
// @Param tag         path  string true "tag name"
// @Param count_films query int    true "count films"
// @Param delimiter   query string true "last value while in is rating last returned film"
// @Success 200 {object} models.TagCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseCollectionNoSuchCollection "no such collection"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection/{tag} [GET]
func (h *tagCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewTagCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrValidation(stdErrors.Cause(err)))
		return
	}

	collection, err := h.collectionService.GetCollectionByTag(r.Context(), request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrCollection(stdErrors.Cause(err)))
		return
	}

	response := models.NewTagCollectionResponse(&collection)

	httpwrapper.Response(w, http.StatusOK, response)
}
