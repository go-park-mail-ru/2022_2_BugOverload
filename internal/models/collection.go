package models

type Collection struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Poster      string  `json:"poster,omitempty"`
	Time        string  `json:"time,omitempty"`
	Films       []Film  `json:"films,omitempty"`
	CountLikes  float32 `json:"count_likes,omitempty"`
	CountFilms  float32 `json:"count_films,omitempty"`
}
