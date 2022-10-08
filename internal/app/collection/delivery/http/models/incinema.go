package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

// FilmsInCinemaRequest is structure for films handler
type FilmsInCinemaRequest struct {
	FilmCollection []models.Film
}

// AddFilm adds film with actual fields for FilmsInCinemaRequest
func (fcr *FilmsInCinemaRequest) AddFilm(film models.Film) {
	fcr.FilmCollection = append(fcr.FilmCollection, models.Film{
		ID:        film.ID,
		Name:      film.Name,
		YearProd:  film.YearProd,
		PosterVer: film.PosterVer,
		Genres:    film.Genres,
		Rating:    film.Rating,
	})
}

// CreateResponse return FilmCollection struct for sending response in FilmsInCinemaRequest
func (fcr *FilmsInCinemaRequest) CreateResponse() models.FilmCollection {
	return models.CreateFilmCollection("Сейчас в кино", fcr.FilmCollection)
}

// PopularFilmsRequest is structure for films handler
type PopularFilmsRequest struct {
	FilmCollection []models.Film
}

// AddFilm adds film with actual fields for PopularFilmsRequest
func (pfr *PopularFilmsRequest) AddFilm(film models.Film) {
	pfr.FilmCollection = append(pfr.FilmCollection, models.Film{
		ID:        film.ID,
		Name:      film.Name,
		YearProd:  film.YearProd,
		PosterVer: film.PosterVer,
		Genres:    film.Genres,
		Rating:    film.Rating,
	})
}

// CreateResponse return FilmCollection struct for sending response in PopularFilmsRequest
func (pfr *PopularFilmsRequest) CreateResponse() models.FilmCollection {
	return models.CreateFilmCollection("Популярное", pfr.FilmCollection)
}
