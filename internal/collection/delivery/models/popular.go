package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type FilmsPopularRequest struct {
	FilmCollection []models.Film
}

func NewFilmsPopularRequest(collection []models.Film) *FilmsPopularRequest {
	return &FilmsPopularRequest{
		collection,
	}
}

func (fcr *FilmsPopularRequest) ToPublic() *models.FilmCollection {
	return models.NewFilmCollection("Популярное", fcr.FilmCollection)
}
