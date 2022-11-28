package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	mainModels "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client"
)

// getCollectionFilmsHandler is the structure that handles the request for movies in cinemas.
type getCollectionFilmsHandler struct {
	collectionService client.WarehouseService
}

// NewGetCollectionFilmsHandler is constructor for getCollectionFilmsHandler in this pkg - in cinema.
func NewGetCollectionFilmsHandler(service client.WarehouseService) handler.Handler {
	return &getCollectionFilmsHandler{
		service,
	}
}

func (h *getCollectionFilmsHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/collections/{id:[0-9]+}", mw.TryAuthMiddleware(h.Action)).
		Methods(http.MethodGet).
		Queries("sort_param", "{sort_param}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Getter std collection (tag, genre)
// @Description Films by genre or tag by DESC rating or DESC date. prod_company, prod_country also in the future.
// @Description All fields required
// @tags collection, completed
// @Produce json
// @Param sort_param  query string true "rating, date"
// @Success 200 {object} models.GetStdCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 403 "return error: forbidden"
// @Failure 404 {object} httpmodels.ErrResponseCollectionNoSuchCollection "no such collection"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collections/{id} [GET]
func (h *getCollectionFilmsHandler) Action(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constparams.CurrentUserKey).(mainModels.User)

	request := models.NewCollectionGetFilmsRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	var collection mainModels.Collection

	if ok {
		collection, err = h.collectionService.GetCollectionFilmsAuthorized(r.Context(), &user, request.GetParams())
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
			return
		}
	} else {
		collection, err = h.collectionService.GetCollectionFilmsNotAuthorized(r.Context(), request.GetParams())
		if err != nil {
			wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
			return
		}
	}

	response := models.NewCollectionGetFilmsResponse(&collection)

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
