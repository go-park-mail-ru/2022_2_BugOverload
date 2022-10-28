package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type RecommendFilmResponse struct {
	ID               uint     `json:"id,omitempty" example:"23"`
	Name             string   `json:"name,omitempty" example:"Терминатор"`
	ShortDescription string   `json:"short_description,omitempty" example:"Идет борьба сопротивления людей против машин"`
	ProdYear         int      `json:"prod_year,omitempty" example:"2008"`
	EndYear          int      `json:"end_year,omitempty" example:"2013"`
	PosterHor        string   `json:"poster_hor,omitempty" example:"{{ключ}}"`
	Rating           string   `json:"ratio,omitempty" example:"8.8"`
	Genres           []string `json:"genres,omitempty" example:"фантастика,боевик"`
}

func NewRecommendFilmResponse() *RecommendFilmResponse {
	return &RecommendFilmResponse{}
}

func (rfr *RecommendFilmResponse) ToPublic(film *models.Film) models.Film {
	return models.Film{
		ID:               film.ID,
		Name:             film.Name,
		ShortDescription: film.ShortDescription,
		ProdYear:         film.ProdYear,
		EndYear:          film.EndYear,
		PosterHor:        film.PosterHor,
		Genres:           film.Genres,
		Rating:           film.Rating,
	}
}
