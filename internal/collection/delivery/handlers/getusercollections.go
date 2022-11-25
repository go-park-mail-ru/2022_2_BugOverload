package handlers

import (
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
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

func (h *getUserCollectionsHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/user/collections", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodGet).
		Queries(
			"sort_param", "{sort_param}",
			"count_collections", "{count_collections}",
			"delimiter", "{delimiter}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Get user personal collections
// @Description Collections by create time OR update time
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param sort_param        query string true "create_time, update_time"
// @Param count_collections query int    true "count collections"
// @Param delimiter         query string true "last value while in is date last returned collection for create_time AND update_time for first for both is 'now'"
// @Success 200 {array} models.ShortFilmCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 404 "collection not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/user/collections [GET]
func (h *getUserCollectionsHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)
	if !ok {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewGetUserCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetUserCollections(r.Context(), &user, request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response := models.NewShortFilmCollectionResponse(collection)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
