package popularfilmshandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// handler is structure for API films requests processing
type handler struct {
	storage *memory.FilmStorage
}

// NewHandler is constructor for handler
func NewHandler(fs *memory.FilmStorage) *handler {
	return &handler{
		fs,
	}
}

// tmp const
const countParts = 2
const countFilmPreview = 5

// Action is handle getPopularFilms request
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var popularFilmRequest models.PopularFilmsRequest

	upperBound := 0
	if (h.storage.GetStorageLen()-countFilmPreview)/countParts > 0 {
		upperBound = (h.storage.GetStorageLen() - countFilmPreview) / countParts
	}

	for i := 0; i < upperBound; i++ {
		film, err := h.storage.GetFilm(uint(i))
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
