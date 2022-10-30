package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type PersonResponse struct {
	ID           int      `json:"id,omitempty" example:"4526"`
	Name         string   `json:"name,omitempty" example:"Шон Коннери"`
	OriginalName string   `json:"original_name,omitempty" example:"Sean Connery"`
	Birthday     string   `json:"birthday,omitempty" example:"1930-08-25"`
	Death        string   `json:"death,omitempty" example:"2020-10-31"`
	Growth       float32  `json:"growth,omitempty" example:"1.9"`
	Gender       string   `json:"gender,omitempty" example:"male"`
	Avatar       string   `json:"avatar,omitempty" example:"4526"`
	CountFilms   int      `json:"count_films,omitempty" example:"218"`
	Professions  []string `json:"professions,omitempty" example:"актер,продюсер,режиссер"`
	Genres       []string `json:"genres,omitempty" example:"драма,боевик,триллер"`
	Images       []string `json:"images,omitempty" example:"1,2,3,4,5,6,7"`
}

func NewPersonResponse(person *models.Person) *PersonResponse {
	return &PersonResponse{
		ID:          person.ID,
		Name:        person.Name,
		Birthday:    person.OriginalName,
		Death:       person.Death,
		Growth:      person.Growth,
		Gender:      person.Gender,
		CountFilms:  person.CountFilms,
		Professions: person.Professions,
		Genres:      person.Genres,
	}
}

func (p *PersonResponse) ToPublic() *PersonResponse {
	return p
}
