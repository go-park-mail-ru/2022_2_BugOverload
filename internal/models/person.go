package models

type Person struct {
	ID          uint     `json:"user_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Birthday    string   `json:"birthday,omitempty"`
	Death       string   `json:"death,omitempty"`
	Gender      string   `json:"gender,omitempty"`
	CountFilms  int      `json:"count_films,omitempty"`
	Professions []string `json:"professions,omitempty"`
}
