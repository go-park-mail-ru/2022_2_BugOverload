package models

type Person struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	OriginalName string   `json:"original_name,omitempty"`
	Birthday     string   `json:"birthday,omitempty"`
	Avatar       string   `json:"avatar,omitempty"`
	Death        string   `json:"death,omitempty"`
	Growth       float32  `json:"growth,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	CountFilms   int      `json:"count_films,omitempty"`
	Professions  []string `json:"professions,omitempty"`
	Genres       []string `json:"genres,omitempty"`
	Images       []string `json:"images,omitempty"`
}
