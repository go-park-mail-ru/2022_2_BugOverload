package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type ShortFilmCollectionResponse struct {
	Name       string `json:"name,omitempty" example:"Популярное"`
	Poster     string `json:"poster,omitempty"`
	CountLikes string `json:"count_likes,omitempty"`
	CountFilms string `json:"count_films,omitempty"`
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
