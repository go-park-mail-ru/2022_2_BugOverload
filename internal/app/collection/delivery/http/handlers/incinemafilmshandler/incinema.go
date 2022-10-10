package incinemafilmshandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/models"
)

// handler is structure for API films requests processing
type handler struct {
	storage *memory.FilmStorage
}

// NewHandler is constructor for handler
func NewHandler(fs *memory.FilmStorage) interfaces.Handler {
	return &handler{
		fs,
	}
}

// tmp const
const countParts = 2
const countFilmPreview = 5

// Action is handle InCinema request
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	var inCinemaRequest models.FilmsInCinemaRequest

	var upperBound, lowerBound int

	if h.storage.GetStorageLen()-countFilmPreview > 0 {
		upperBound = h.storage.GetStorageLen() - countFilmPreview
	}
	if (h.storage.GetStorageLen()-countFilmPreview)/countParts >= 0 {
		lowerBound = (h.storage.GetStorageLen() - countFilmPreview) / countParts
	}

	for i := upperBound; i >= lowerBound; i-- {
		film, err := h.storage.GetFilm(uint(i))
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
