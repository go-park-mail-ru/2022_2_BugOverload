package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

type filmsInCinemaRequest struct {
	FilmCollection []models.Film
}

func NewFilmsInCinemaRequest(collection []models.Film) *filmsInCinemaRequest {
	return &filmsInCinemaRequest{
		collection,
	}
}

func (fcr *filmsInCinemaRequest) CreateResponse() *models.FilmCollection {
	return models.NewFilmCollection("Сейчас в кино", fcr.FilmCollection)
}
