package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type RecommendFilmRequest struct {
	recommendedFilm models.Film
}

func NewRecommendFilmRequest(film models.Film) *RecommendFilmRequest {
	return &RecommendFilmRequest{
		models.Film{
			ID:               film.ID,
			Name:             film.Name,
			ShortDescription: film.ShortDescription,
			YearProd:         film.YearProd,
			PosterHor:        film.PosterHor,
			Genres:           film.Genres,
			Rating:           film.Rating,
		},
	}
}

func (rfr *RecommendFilmRequest) ToPublic() models.Film {
	return rfr.recommendedFilm
}
