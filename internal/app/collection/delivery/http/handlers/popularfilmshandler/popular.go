package popularfilmshandler

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// CollectionPopularHandler is structure for API films requests processing
type CollectionPopularHandler struct {
	storage *memory.FilmStorage
}

// NewCollectionPopularHandler is constructor for NewCollectionPopularHandler
func NewCollectionPopularHandler(fs *memory.FilmStorage) *CollectionPopularHandler {
	return &CollectionPopularHandler{
		fs,
	}
}

//const countParts = 2

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
