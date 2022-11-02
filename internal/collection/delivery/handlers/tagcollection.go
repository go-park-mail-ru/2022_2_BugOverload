package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
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
// @Description Films from the "in cinema" category
// @tags collection
// @Produce json
// @Success 200 {object} models.FilmCollectionInCinemaResponse "returns an array of movies"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /api/v1/collection/{tag} [GET]
func (h *tagCollectionHandler) Action(w http.ResponseWriter, r *http.Request) {
	// in dev
}
