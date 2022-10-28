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

// PopularFilmsHandler is the structure that handles the request for popular movies.
type PopularFilmsHandler struct {
	collectionService service.CollectionService
}

// NewPopularFilmsHandler is constructor for PopularFilmsHandler in this pkg - popular.
func NewPopularFilmsHandler(uc service.CollectionService) pkg.Handler {
	return &PopularFilmsHandler{
		uc,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Popular movies
// @Description Films from the "popular" category
// @tags collection
// @Produce json
// @Success 200 {object} models.FilmCollectionResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collections/popular [GET]
func (h *PopularFilmsHandler) Action(w http.ResponseWriter, r *http.Request) {
	collection, err := h.collectionService.GetPopular(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(stdErrors.Cause(err)))
		return
	}

	collectionPopular := models.NewFilmCollectionResponse(&collection)

	httpwrapper.Response(w, http.StatusOK, collectionPopular.ToPublic())
}
