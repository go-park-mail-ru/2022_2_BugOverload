package models

import "go-park-mail-ru/2022_2_BugOverload/internal/app/models"

// RecommendFilmRequest is structure for films handler
type RecommendFilmRequest struct {
	recommendedFilm models.Film
}

// SetFilm set film with actual fields for RecommendFilmRequest
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

// CreateResponse return Film struct for sending response for RecommendFilmRequest
func (rfr *RecommendFilmRequest) CreateResponse() models.Film {
	return rfr.recommendedFilm
}
