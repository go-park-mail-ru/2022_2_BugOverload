package incinemafilmshandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
)

// CollectionInCinemaHandler is structure for API films requests processing
type CollectionInCinemaHandler struct {
	storage *memory.FilmStorage
}

// NewCollectionInCinemaHandler is constructor for NewCollectionInCinemaHandler
func NewCollectionInCinemaHandler(fs *memory.FilmStorage) *CollectionInCinemaHandler {
	return &CollectionInCinemaHandler{
		fs,
	}
}

// tmp const
const countParts = 2
const countFilmPreview = 5

// Action is handle InCinema request
func (hf *CollectionInCinemaHandler) Action(w http.ResponseWriter, r *http.Request) {
	var inCinemaRequest models.FilmsInCinemaRequest

	var upperBound, lowerBound int

	if hf.storage.GetStorageLen()-countFilmPreview > 0 {
		upperBound = hf.storage.GetStorageLen() - countFilmPreview
	}
	if (hf.storage.GetStorageLen()-countFilmPreview)/countParts >= 0 {
		lowerBound = (hf.storage.GetStorageLen() - countFilmPreview) / countParts
	}

	for i := upperBound; i >= lowerBound; i-- {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}

		inCinemaRequest.AddFilm(film)
	}

	if len(inCinemaRequest.FilmCollection) == 0 {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrFilms(errors2.ErrFilmNotFound))

		return
	}

	response := inCinemaRequest.CreateResponse()

	httpwrapper2.Response(w, http.StatusOK, response)
}
