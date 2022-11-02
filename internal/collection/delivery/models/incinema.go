package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionInCinemaResponse struct {
	ID        int      `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  int      `json:"prod_year,omitempty" example:"2014"`
	EndYear   int      `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionInCinemaResponse struct {
	Name  string                             `json:"name,omitempty" example:"Сейчас в кино"`
	Films []filmInCollectionInCinemaResponse `json:"film,omitempty"`
}

func NewFilmInCollectionInCinemaResponse(collection *models.Collection) *FilmCollectionInCinemaResponse {
	res := &FilmCollectionInCinemaResponse{
		Name: collection.Name,
	}

	res.Films = make([]filmInCollectionInCinemaResponse, len(collection.Films))

	for idx, value := range collection.Films {
		res.Films[idx] = filmInCollectionInCinemaResponse{
			ID:        value.ID,
			Name:      value.Name,
			ProdYear:  value.ProdYear,
			EndYear:   value.EndYear,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	return res
}