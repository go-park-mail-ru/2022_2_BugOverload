package handlers

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type HandlerFilms struct {
	storage *database.FilmStorage
}

// NewHandlerFilms is constructor for HandlerFilms
func NewHandlerFilms(fs *database.FilmStorage) *HandlerFilms {
	return &HandlerFilms{fs}
}

// GetPopularFilms is handle getPopularFilms request
func (hf *HandlerFilms) GetPopularFilms(w http.ResponseWriter, r *http.Request) {
	var films []structs.Film
	for i := 0; i < 6; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		films = append(films, film)
	}

	if len(films) == 0 {
		http.Error(w, errorshandlers.ErrFilmsNotFound.Error(), http.StatusNotFound)
		return
	}

	response := structs.CreateFilmCollection("Popular films", films)

	httpwrapper.Response(w, http.StatusOK, response)
}

// GetFilmsInCinema is handle InCinema request
func (hf *HandlerFilms) GetFilmsInCinema(w http.ResponseWriter, r *http.Request) {
	var films []structs.Film
	for i := 6; i < 12; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		films = append(films, film)
	}

	if len(films) == 0 {
		http.Error(w, errorshandlers.ErrFilmsNotFound.Error(), http.StatusNotFound)
		return
	}

	response := structs.CreateFilmCollection("In cinema", films)

	httpwrapper.Response(w, http.StatusOK, response)
}

// GetFilmToPoster is handle film to poster request
func (hf *HandlerFilms) GetFilmToPoster(w http.ResponseWriter, r *http.Request) {
	response, err := hf.storage.GetFilm(uint(hf.storage.GetStorageLen() - 1))
	if err != nil {
		http.Error(w, errorshandlers.ErrFilmNotFound.Error(), http.StatusNotFound)
		return
	}

	httpwrapper.Response(w, http.StatusOK, response)
}
