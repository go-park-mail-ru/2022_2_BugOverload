package handlers

import (
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
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
	r.HandleFunc("/api/v1/user/collections", mw.CheckAuthMiddleware(mw.SetCsrfMiddleware(h.Action))).
		Methods(http.MethodGet).
		Queries(
			"sort_param", "{sort_param}",
			"count_collections", "{count_collections}",
			"delimiter", "{delimiter}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getter std collection (tag, genre)
// @Description Films by genre or tag by DESC rating or DESC date. prod_company, prod_country also in the future.
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param sort_param        query string true "create_time, update_time"
// @Param count_collections query int    true "count collections"
// @Param delimiter         query string true "last value while in is date last returned collection for create_time AND update_time"
// @Success 200 {object} models.GetStdCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 404 {object} httpmodels.ErrResponseCollectionNoSuchCollection "no such collection"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection [GET]
func (h *getUserCollectionsHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(pkg.CurrentUserKey).(mainModels.User)
	if !ok {
		httpwrapper.DefaultHandlerError(r.Context(), w, errors.ErrGetUserRequest)
		return
	}

	request := models.NewGetUserCollectionRequest()

	err := request.Bind(r)
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	collection, err := h.collectionService.GetUserCollections(r.Context(), &user, request.GetParams())
	if err != nil {
		httpwrapper.DefaultHandlerError(r.Context(), w, err)
		return
	}

	response := models.NewShortFilmCollectionResponse(collection)

	httpwrapper.Response(r.Context(), w, http.StatusOK, response)
}
