package content

import (
	"math/rand"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// HandlerFilms is structure for API films requests processing
type HandlerFilms struct {
	storage *database.FilmStorage
}

// NewHandlerFilms is constructor for HandlerFilms
func NewHandlerFilms(fs *database.FilmStorage) *HandlerFilms {
	return &HandlerFilms{
		fs,
	}
}

// PopularFilmsRequest is structure for films handler
type PopularFilmsRequest struct {
	filmCollection []structs.Film
}

// AddFilm adds film with actual fields for PopularFilmsRequest
func (pfr *PopularFilmsRequest) AddFilm(film structs.Film) {
	pfr.filmCollection = append(pfr.filmCollection, structs.Film{
		ID:        film.ID,
		Name:      film.Name,
		YearProd:  film.YearProd,
		PosterVer: film.PosterVer,
		Genres:    film.Genres,
		Rating:    film.Rating,
	})
}

// CreateResponse return FilmCollection struct for sending response in PopularFilmsRequest
func (pfr *PopularFilmsRequest) CreateResponse() structs.FilmCollection {
	return structs.CreateFilmCollection("Популярное", pfr.filmCollection)
}

// GetPopularFilms is handle getPopularFilms request
func (hf *HandlerFilms) GetPopularFilms(w http.ResponseWriter, r *http.Request) {
	var popularFilmRequest PopularFilmsRequest

	for i := hf.storage.GetStorageLen() - 5; i >= 0; i-- {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}
		popularFilmRequest.AddFilm(film)
	}

	if len(popularFilmRequest.filmCollection) == 0 {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrFilmNotFound)

		return
	}

	response := popularFilmRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}

// FilmsInCinemaRequest is structure for films handler
type FilmsInCinemaRequest struct {
	filmCollection []structs.Film
}

// AddFilm adds film with actual fields for FilmsInCinemaRequest
func (fcr *FilmsInCinemaRequest) AddFilm(film structs.Film) {
	fcr.filmCollection = append(fcr.filmCollection, structs.Film{
		ID:        film.ID,
		Name:      film.Name,
		YearProd:  film.YearProd,
		PosterVer: film.PosterVer,
		Genres:    film.Genres,
		Rating:    film.Rating,
	})
}

// CreateResponse return FilmCollection struct for sending response in FilmsInCinemaRequest
func (fcr *FilmsInCinemaRequest) CreateResponse() structs.FilmCollection {
	return structs.CreateFilmCollection("Сейчас в кино", fcr.filmCollection)
}

// GetFilmsInCinema is handle InCinema request
func (hf *HandlerFilms) GetFilmsInCinema(w http.ResponseWriter, r *http.Request) {
	var inCinemaRequest FilmsInCinemaRequest

	for i := 0; i < hf.storage.GetStorageLen()-4; i++ {
		film, err := hf.storage.GetFilm(uint(i))
		if err != nil {
			continue
		}

		inCinemaRequest.AddFilm(film)
	}

	if len(inCinemaRequest.filmCollection) == 0 {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrFilmNotFound)

		return
	}

	response := inCinemaRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}

// RecommendFilmRequest is structure for films handler
type RecommendFilmRequest struct {
	recommendedFilm structs.Film
}

// SetFilm set film with actual fields for RecommendFilmRequest
func (rfr *RecommendFilmRequest) SetFilm(film structs.Film) {
	rfr.recommendedFilm = structs.Film{
		ID:               film.ID,
		Name:             film.Name,
		ShortDescription: film.ShortDescription,
		YearProd:         film.YearProd,
		PosterHor:        film.PosterHor,
		Genres:           film.Genres,
		Rating:           film.Rating,
	}
}

// CreateResponse return Film struct for sending response for RecommendFilmRequest
func (rfr *RecommendFilmRequest) CreateResponse() structs.Film {
	return rfr.recommendedFilm
}

// GetRecommendedFilm is handle film to poster request
func (hf *HandlerFilms) GetRecommendedFilm(w http.ResponseWriter, r *http.Request) {
	var recommendFilmRequest RecommendFilmRequest

	max := hf.storage.GetStorageLen()
	min := max - 3

	film, err := hf.storage.GetFilm(uint(rand.Intn(max-min) + min))
	if err != nil {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrFilmNotFound)

		return
	}

	recommendFilmRequest.SetFilm(film)

	response := recommendFilmRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
