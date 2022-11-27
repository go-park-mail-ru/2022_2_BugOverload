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

// dropFilmHandler is the structure that handles the request for
// drop film from collection.
type dropFilmHandler struct {
	collectionService service.CollectionService
}

// NewDropFilmHandler is constructor for dropFilmHandler in this pkg
func NewDropFilmHandler(uc service.CollectionService) handler.Handler {
	return &dropFilmHandler{
		uc,
	}
}

func (h *dropFilmHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/remove", h.Action).
		Methods(http.MethodGet)
}

func (h *dropFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewDropFilmRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	err = h.collectionService.DropFilmFromCollection(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.NoBody(w, http.StatusOK)
}
