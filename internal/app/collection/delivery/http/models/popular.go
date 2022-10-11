package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

type filmsPopularRequest struct {
	FilmCollection []models.Film
}

func NewFilmsPopularRequest(collection []models.Film) *filmsPopularRequest {
	return &filmsPopularRequest{
		collection,
	}
}

func (fcr *filmsPopularRequest) CreateResponse() *models.FilmCollection {
	return models.NewFilmCollection("Популярное", fcr.FilmCollection)
}
