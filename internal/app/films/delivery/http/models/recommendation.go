package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

type RecommendFilmRequest struct {
	recommendedFilm models.Film
}

func (rfr *RecommendFilmRequest) SetFilm(film models.Film) {
	rfr.recommendedFilm = models.Film{
		ID:               film.ID,
		Name:             film.Name,
		ShortDescription: film.ShortDescription,
		YearProd:         film.YearProd,
		PosterHor:        film.PosterHor,
		Genres:           film.Genres,
		Rating:           film.Rating,
	}
}

func (rfr *RecommendFilmRequest) CreateResponse() models.Film {
	return rfr.recommendedFilm
}
