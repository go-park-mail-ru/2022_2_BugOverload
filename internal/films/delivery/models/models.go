package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type RecommendFilmRequest struct {
	ID               uint     `json:"film_id,omitempty"`
	Name             string   `json:"film_name,omitempty"`
	ShortDescription string   `json:"short_description,omitempty"`
	YearProd         string   `json:"year_prod,omitempty"`
	PosterHor        string   `json:"poster_hor,omitempty"`
	Rating           string   `json:"ratio,omitempty"`
	Genres           []string `json:"genres,omitempty"`
}

func NewRecommendFilmRequest() *RecommendFilmRequest {
	return &RecommendFilmRequest{}
}

func (rfr *RecommendFilmRequest) ToPublic(film *models.Film) models.Film {
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
