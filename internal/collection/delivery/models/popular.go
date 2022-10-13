package models

import (
	models2 "go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type FilmsPopularRequest struct {
	FilmCollection []models2.Film
}

func NewFilmsPopularRequest(collection []models2.Film) *FilmsPopularRequest {
	return &FilmsPopularRequest{
		collection,
	}
}

func (fcr *FilmsPopularRequest) ToPublic() *models2.FilmCollection {
	return models2.NewFilmCollection("Популярное", fcr.FilmCollection)
}
