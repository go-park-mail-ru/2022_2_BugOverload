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

// addFilmHandler is the structure that handles the request for
// adding film to collection.
type addFilmHandler struct {
	collectionService service.CollectionService
}

// NewAddFilmHandler is constructor for addFilmHandler in this pkg
func NewAddFilmHandler(uc service.CollectionService) handler.Handler {
	return &addFilmHandler{
		uc,
	}
}

func (h *addFilmHandler) Configure(r *mux.Router, mw *middleware.HTTPMiddleware) {
	r.HandleFunc("/api/v1/film/{id:[0-9]+}/save", h.Action).
		Methods(http.MethodGet)
}

func (h *addFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewAddFilmRequest()

	err := request.Bind(r)
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	err = h.collectionService.AddFilmToCollection(r.Context(), request.GetParams())
	if err != nil {
		wrapper.DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	wrapper.NoBody(w, http.StatusOK)
}
