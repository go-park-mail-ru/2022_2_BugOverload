package models

type Collection struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Poster      string `json:"poster,omitempty"`
	Time        string `json:"time,omitempty"`
	Films       []Film `json:"films,omitempty"`
	CountLikes  int    `json:"count_likes,omitempty"`
	CountFilms  int    `json:"count_films,omitempty"`
}
