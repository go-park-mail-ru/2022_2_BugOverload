package incinema_films_handler

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
)

// CollectionInCinemaHandler is structure for API films requests processing
type CollectionInCinemaHandler struct {
	storage *database.FilmStorage
}

// NewCollectionInCinemaHandler is constructor for NewCollectionInCinemaHandler
func NewCollectionInCinemaHandler(fs *database.FilmStorage) *CollectionInCinemaHandler {
	return &CollectionInCinemaHandler{
		fs,
	}
}

// Action is handle InCinema request
func (hf *CollectionInCinemaHandler) Action(w http.ResponseWriter, r *http.Request) {
	var inCinemaRequest models.FilmsInCinemaRequest

	var upperBound, lowerBound int

	if hf.storage.GetStorageLen()-5 > 0 {
		upperBound = hf.storage.GetStorageLen() - 5
	}
	if (hf.storage.GetStorageLen()-5)/2 >= 0 {
		lowerBound = (hf.storage.GetStorageLen() - 5) / 2
	}

	for i := upperBound; i >= lowerBound; i-- {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}

		inCinemaRequest.AddFilm(film)
	}

	if len(inCinemaRequest.FilmCollection) == 0 {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))

		return
	}

	response := inCinemaRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
