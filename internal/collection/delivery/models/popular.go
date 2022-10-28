package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionPopularResponse struct {
	ID        uint     `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  int      `json:"prod_year,omitempty" example:"2014"`
	EndYear   int      `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionPopularResponse struct {
	Name  string                            `json:"name,omitempty" example:"Популярное"`
	Films []filmInCollectionPopularResponse `json:"film,omitempty"`
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
			ProdYear:  value.ProdYear,
			EndYear:   value.EndYear,
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
