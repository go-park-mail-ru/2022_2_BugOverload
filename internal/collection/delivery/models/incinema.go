package models

import (
	models2 "go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type FilmsInCinemaRequest struct {
	FilmCollection []models2.Film
}

func NewFilmsInCinemaRequest(collection []models2.Film) *FilmsInCinemaRequest {
	return &FilmsInCinemaRequest{
		collection,
	}
}

func (fcr *FilmsInCinemaRequest) ToPublic() *models2.FilmCollection {
	return models2.NewFilmCollection("Сейчас в кино", fcr.FilmCollection)
}
