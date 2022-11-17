package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type RecommendFilmResponse struct {
	ID               int    `json:"id,omitempty" example:"23"`
	Name             string `json:"name,omitempty" example:"Терминатор"`
	ShortDescription string `json:"short_description,omitempty" example:"Идет борьба сопротивления людей против машин"`
	ProdYear         string `json:"prod_year,omitempty" example:"2008"`
	PosterHor        string `json:"poster_hor,omitempty" example:"{{ключ}}"`

	EndYear string `json:"end_year,omitempty" example:"2013"`

	Rating float32 `json:"rating,omitempty" example:"8.8"`

	Genres []string `json:"genres,omitempty" example:"фантастика,боевик"`
}

func NewRecommendFilmResponse(film *models.Film) *RecommendFilmResponse {
	return &RecommendFilmResponse{
		ID:               film.ID,
		Name:             film.Name,
		ShortDescription: film.ShortDescription,
		ProdYear:         film.ProdDate[:len(innerPKG.OnlyDate)],
		EndYear:          film.EndYear,
		PosterHor:        film.PosterHor,
		Genres:           film.Genres,
		Rating:           film.Rating,
	}
}
