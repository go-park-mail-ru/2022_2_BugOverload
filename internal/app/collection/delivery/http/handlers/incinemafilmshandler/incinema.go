package incinemafilmshandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
	collectionInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// handler is the structure that handles the request for movies in cinemas.
type handler struct {
	collectionService collectionInterface.CollectionService
}

// NewHandler is constructor for handler in this pkg - in cinema.
func NewHandler(uc collectionInterface.CollectionService) interfaces.Handler {
	return &handler{
		uc,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	collection, err := h.collectionService.GetInCinema(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	collectionInCinema := models.NewFilmsInCinemaRequest(collection)

	httpwrapper.Response(w, http.StatusOK, collectionInCinema.ToPublic())
}
