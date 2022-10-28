package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionResponse struct {
	ID         uint     `json:"film_id,omitempty" example:"23"`
	Name       string   `json:"film_name,omitempty" example:"Game of Thrones"`
	ProdDate   string   `json:"prod_date,omitempty" example:"2014"`
	PosterVer  string   `json:"poster_ver,omitempty" example:"{{ссылка}}"`
	Rating     float32  `json:"rating,omitempty" example:"7.9"`
	Genres     []string `json:"genres,omitempty" example:"фэнтези,приключения"`
	CountLikes string   `json:"count_likes,omitempty"`
}

type FilmCollectionResponse struct {
	Name        string                     `json:"name,omitempty" example:"Популярное"`
	Description string                     `json:"description,omitempty"`
	Poster      string                     `json:"poster,omitempty"`
	Time        string                     `json:"time,omitempty"`
	CountLikes  string                     `json:"count_likes,omitempty"`
	Films       []filmInCollectionResponse `json:"films,omitempty"`
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
			ProdDate:  value.ProdDate,
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
