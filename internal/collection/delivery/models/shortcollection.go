package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type ShortFilmCollectionResponse struct {
	Name       string  `json:"name,omitempty" example:"Популярное"`
	Poster     string  `json:"poster,omitempty" example:"42"`
	CountLikes float32 `json:"count_likes,omitempty" example:"1023"`
	CountFilms float32 `json:"count_films,omitempty"  example:"10"`
}

func NewShortFilmCollectionResponse(collection *models.Collection) *ShortFilmCollectionResponse {
	res := &ShortFilmCollectionResponse{
		Name:       collection.Name,
		Poster:     collection.Poster,
		CountLikes: collection.CountLikes,
		CountFilms: collection.CountFilms,
	}

	return res
}

func (fcr *ShortFilmCollectionResponse) ToPublic() *ShortFilmCollectionResponse {
	return fcr
}
