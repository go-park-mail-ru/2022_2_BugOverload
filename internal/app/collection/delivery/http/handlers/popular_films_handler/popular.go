package popular_films_handler

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
)

// CollectionPopularHandler is structure for API films requests processing
type CollectionPopularHandler struct {
	storage *database.FilmStorage
}

// NewCollectionPopularHandler is constructor for NewCollectionPopularHandler
func NewCollectionPopularHandler(fs *database.FilmStorage) *CollectionPopularHandler {
	return &CollectionPopularHandler{
		fs,
	}
}

// Action is handle getPopularFilms request
func (hf *CollectionPopularHandler) Action(w http.ResponseWriter, r *http.Request) {
	var popularFilmRequest models.PopularFilmsRequest

	upperBound := 0
	if (hf.storage.GetStorageLen()-5)/2 > 0 {
		upperBound = (hf.storage.GetStorageLen() - 5) / 2
	}

	for i := 0; i < upperBound; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		popularFilmRequest.AddFilm(film)
	}

	if len(popularFilmRequest.FilmCollection) == 0 {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))

		return
	}

	response := popularFilmRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
