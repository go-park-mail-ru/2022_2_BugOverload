package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionPopularResponse struct {
	ID        uint     `json:"film_id,omitempty" example:"23"`
	Name      string   `json:"film_name,omitempty" example:"Game of Thrones"`
	ProdDate  string   `json:"prod_date,omitempty" example:"2014"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionPopularResponse struct {
	Name  string                            `json:"name,omitempty" example:"Популярное"`
	Films []filmInCollectionPopularResponse `json:"films,omitempty"`
}

func NewFilmInCollectionPopularResponse(collection *models.Collection) *FilmCollectionPopularResponse {
	res := &FilmCollectionPopularResponse{
		Name: collection.Name,
	}

	res.Films = make([]filmInCollectionPopularResponse, len(collection.Films))

	for idx, value := range collection.Films {
		res.Films[idx] = filmInCollectionPopularResponse{
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

func (fcr *FilmCollectionPopularResponse) ToPublic() *FilmCollectionPopularResponse {
	return fcr
}
