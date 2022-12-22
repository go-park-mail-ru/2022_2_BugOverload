package models

//go:generate easyjson -all -disallow_unknown_fields collection.go

type Collection struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Poster      string `json:"poster,omitempty"`
	Time        string `json:"time,omitempty"`

	CountLikes int `json:"count_likes,omitempty"`
	CountFilms int `json:"count_films,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`
	CreateTime string `json:"create_time,omitempty"`

	Films []Film `json:"films,omitempty"`

	Author User `json:"author,omitempty"`
}
