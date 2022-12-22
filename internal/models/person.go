package models

//go:generate easyjson -all -disallow_unknown_fields person.go

type Person struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	OriginalName string  `json:"original_name,omitempty"`
	Birthday     string  `json:"birthday,omitempty"`
	Avatar       string  `json:"avatar,omitempty"`
	GrowthMeters float32 `json:"growth,omitempty"`
	Gender       string  `json:"gender,omitempty"`

	Death string `json:"death,omitempty"`

	CountFilms int `json:"count_films,omitempty"`

	Professions []string `json:"professions,omitempty"`
	Genres      []string `json:"genres,omitempty"`
	Images      []string `json:"images,omitempty"`

	BestFilms []Film `json:"best_films,omitempty"`
}
