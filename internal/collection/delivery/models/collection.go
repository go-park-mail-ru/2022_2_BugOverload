package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionResponse struct {
	ID        uint     `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  int      `json:"prod_year,omitempty" example:"2014"`
	EndYear   int      `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionResponse struct {
	Name        string                     `json:"name,omitempty" example:"Популярное"`
	Description string                     `json:"description,omitempty"  example:"Лучшие фильмы на данный момент"`
	Poster      string                     `json:"poster,omitempty" example:"42"`
	Time        string                     `json:"time,omitempty" example:"2023.01.04 15:12:23'"`
	CountLikes  float32                    `json:"count_likes,omitempty" example:"502"`
	Films       []filmInCollectionResponse `json:"film,omitempty"`
}

func NewFilmCollectionResponse(collection *models.Collection) *FilmCollectionResponse {
	res := &FilmCollectionResponse{
		Name:        collection.Name,
		Description: collection.Description,
		Poster:      collection.Poster,
		Time:        collection.Time,
		CountLikes:  collection.CountLikes,
	}

	res.Films = make([]filmInCollectionResponse, len(collection.Films))

	for idx, value := range collection.Films {
		res.Films[idx] = filmInCollectionResponse{
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

func (fcr *FilmCollectionResponse) ToPublic() *FilmCollectionResponse {
	return fcr
}
