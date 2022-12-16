package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client"
)

// searchHandler is the structure that handles the request for search.
type searchHandler struct {
	searchService client.WarehouseService
}

// NewSearchHandler is constructor for searchHandler in this pkg.
func NewSearchHandler(service client.WarehouseService) handler.Handler {
	return &searchHandler{
		service,
	}
}

func (h *searchHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/search", h.Action).Methods(http.MethodGet).
		Queries("q", "{q}")
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Search films, serials and persons
// @Description Getting film, serials and persons info. WARNING for films no fields end_year.
// @tags search, completed
// @Produce json
// @Param q  query string true "search body"
// @Success 200 {object} models.SearchResponse "successfully search"
// @Failure 400 "return error"
// @Failure 404 "return error: not found"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/search [GET]
func (h *searchHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewSearchRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, errBind)
		return
	}

	searchRes, err := h.searchService.Search(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	response, err := models.NewSearchResponse(searchRes)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.Response(r.Context(), w, http.StatusOK, response)
}
