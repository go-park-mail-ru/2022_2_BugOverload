package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// inCinemaHandler is the structure that handles the request for movies in cinemas.
type inCinemaHandler struct {
	collectionService service.CollectionService
}

// NewInCinemaHandler is constructor for inCinemaHandler in this pkg - in cinema.
func NewInCinemaHandler(uc service.CollectionService) pkg.Handler {
	return &inCinemaHandler{
		uc,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary In cinema movies
// @Description Films from the "in cinema" category
// @tags collection
// @Produce json
// @Success 200 {object} models.FilmCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/in_cinema [GET]
func (h *inCinemaHandler) Action(w http.ResponseWriter, r *http.Request) {
	collection, err := h.collectionService.GetInCinema(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	collectionInCinema := models.NewFilmCollectionResponse("Сейчас в кино", collection)

	httpwrapper.Response(w, http.StatusOK, collectionInCinema.ToPublic())
}
