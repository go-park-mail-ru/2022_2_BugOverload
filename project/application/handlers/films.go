package handlers

import (
	"net/http"
	"strconv"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"

	"github.com/wonderivan/logger"
)

type HandlerFilms struct {
	storage *database.FilmStorage
}

// NewHandlerFilms is constructor for HandlerFilms
func NewHandlerFilms(fs *database.FilmStorage) *HandlerFilms {
	return &HandlerFilms{fs}
}

// Handle getPopularFilms request
func (hf *HandlerFilms) GetPopularFilms(w http.ResponseWriter, r *http.Request) {
	// Логируем входящий HTTP запрос

	// Get six films from storage to response
	var films []structs.Film
	for i := 0; i < 6; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		films = append(films, film)
	}

	if len(films) == 0 {
		logger.Error(strconv.FormatInt(int64(len(films)), 10) + " 404 get popular films")
		http.Error(w, errorshandlers.ErrFilmsNotFound.Error(), http.StatusNotFound)
		return
	}

	// Create response (film collection)
	response := structs.CreateFilmCollection("Popular films", films)

	// Send film collection to client
	httpwrapper.ResponseOK(w, http.StatusOK, response)
}

// Handle InCinema request
func (hf *HandlerFilms) GetFilmsInCinema(w http.ResponseWriter, r *http.Request) {
	// Логируем входящий HTTP запрос

	// Get six films from storage to response
	var films []structs.Film
	for i := 6; i < 12; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		films = append(films, film)
	}

	if len(films) == 0 {
		logger.Error(strconv.FormatInt(int64(hf.storage.GetStorageLen()), 10) + " 404 get films in cinema")
		http.Error(w, errorshandlers.ErrFilmsNotFound.Error(), http.StatusNotFound)
		return
	}

	// Create response (film collection)
	response := structs.CreateFilmCollection("In cinema", films)

	// Send film collection to client
	httpwrapper.ResponseOK(w, http.StatusOK, response)
}

// Handle film to poster request
func (hf *HandlerFilms) GetFilmToPoster(w http.ResponseWriter, r *http.Request) {
	// Логируем входящий HTTP запрос

	// Create response (film collection)
	response, err := hf.storage.GetFilm(12)
	if err != nil {
		logger.Error(err.Error() + " 404 get film to poster")
		http.Error(w, errorshandlers.ErrFilmNotFound.Error(), http.StatusNotFound)
		return
	}

	// Send film collection to client
	httpwrapper.ResponseOK(w, http.StatusOK, response)
}
