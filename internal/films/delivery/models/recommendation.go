package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type RecommendFilmResponse struct {
	ID               uint     `json:"film_id,omitempty" example:"23"`
	Name             string   `json:"film_name,omitempty" example:"Терминатор"`
	ShortDescription string   `json:"short_description,omitempty" example:"Идет борьба сопротивления людей против машин"`
	YearProd         string   `json:"year_prod,omitempty" example:"2008"`
	PosterHor        string   `json:"poster_hor,omitempty" example:"{{ссылка}}"`
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
		YearProd:         film.YearProd,
		PosterHor:        film.PosterHor,
		Genres:           film.Genres,
		Rating:           film.Rating,
	}
}
