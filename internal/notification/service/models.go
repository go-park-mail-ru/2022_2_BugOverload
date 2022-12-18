package service

import "go-park-mail-ru/2022_2_BugOverload/internal/models"

type AnonsFilmNotificationPayload struct {
	Name      string  `json:"name,omitempty" example:"Игра престолов"`
	ProdDate  string  `json:"prod_date,omitempty" example:"2014.01.13"`
	PosterVer string  `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32 `json:"rating,omitempty" example:"9.2"`
}

func NewAnonsFilmNotificationPayload(film models.Film) *AnonsFilmNotificationPayload {
	return &AnonsFilmNotificationPayload{
		Name:      film.Name,
		ProdDate:  film.ProdDate,
		PosterVer: film.PosterVer,
		Rating:    film.Rating,
	}
}
