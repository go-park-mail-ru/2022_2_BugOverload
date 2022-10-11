package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

type FilmsInCinemaRequest struct {
	FilmCollection []models.Film
}

func NewFilmsInCinemaRequest(collection []models.Film) *FilmsInCinemaRequest {
	return &FilmsInCinemaRequest{
		collection,
	}
}

func (fcr *FilmsInCinemaRequest) ToPublic() *models.FilmCollection {
	return models.NewFilmCollection("Сейчас в кино", fcr.FilmCollection)
}
