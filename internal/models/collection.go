package models

type Collection struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Poster      string `json:"poster,omitempty"`
	Time        string `json:"time,omitempty"`
	Films       []Film `json:"films,omitempty"`
	CountLikes  string `json:"count_likes,omitempty"`
	CountFilms  string `json:"count_films,omitempty"`
}
