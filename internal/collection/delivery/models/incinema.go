package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionInCinemaResponse struct {
	ID        uint     `json:"film_id,omitempty" example:"23"`
	Name      string   `json:"film_name,omitempty" example:"Game of Thrones"`
	ProdDate  string   `json:"prod_date,omitempty" example:"2014"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionInCinemaResponse struct {
	Name  string                             `json:"name,omitempty" example:"Сейчас в кино"`
	Films []filmInCollectionInCinemaResponse `json:"films,omitempty"`
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
			ProdDate:  value.ProdDate,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	return res
}

func (fcr *FilmCollectionInCinemaResponse) ToPublic() *FilmCollectionInCinemaResponse {
	return fcr
}
