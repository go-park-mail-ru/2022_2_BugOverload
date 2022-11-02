package handlers

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

// tagCollectionHandler is the structure that handles the request for movies in cinemas.
type tagCollectionHandler struct {
	collectionService service.CollectionService
}

// NewTagCollectionHandler is constructor for tagCollectionHandler in this pkg - in cinema.
func NewTagCollectionHandler(uc service.CollectionService) pkg.Handler {
	return &tagCollectionHandler{
		uc,
	}
}

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary In cinema movies
// @Description Films by tag "популярное" or "сейчас в кино"
// @tags collection
// @Produce json
// @Param tag  path string true "tag name"
// @Success 200 {object} models.TagCollectionResponse "returns an array of movies"
// @Failure 400 {object} httpmodels.ErrResponseFilmNoSuchFilm "return error"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection/{tag} [GET]
func (h *tagCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	// in dev
}
