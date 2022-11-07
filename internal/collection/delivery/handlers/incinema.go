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

// inCinemaHandler is the structure that handles the request for movies in cinemas.
type inCinemaHandler struct {
	collectionService service.CollectionService
}

// NewInCinemaHandler is constructor for inCinemaHandler in this pkg - in cinema.
func NewInCinemaHandler(uc service.CollectionService) handler.Handler {
	return &inCinemaHandler{
		uc,
	}
}

func (h *inCinemaHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/api/v1/collection/in_cinema", h.Action).Methods(http.MethodGet)
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *inCinemaHandler) Action(w http.ResponseWriter, r *http.Request) {
	collection, err := h.collectionService.GetInCinema(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	collectionInCinema := models.NewFilmInCollectionInCinemaResponse(&collection)

	httpwrapper.Response(w, http.StatusOK, collectionInCinema)
}
