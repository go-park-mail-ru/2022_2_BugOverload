package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionResponse struct {
	ID        uint     `json:"film_id,omitempty" example:"23"`
	Name      string   `json:"film_name,omitempty" example:"Game of Thrones"`
	YearProd  string   `json:"year_prod,omitempty" example:"2014"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{ссылка}}"`
	Rating    string   `json:"ratio,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type FilmCollectionResponse struct {
	Title string                     `json:"title,omitempty" example:"Популярное"`
	Films []filmInCollectionResponse `json:"films,omitempty"`
}

func NewFilmCollectionResponse(title string, films []models.Film) *FilmCollectionResponse {
	res := &FilmCollectionResponse{
		Title: title,
	}

	res.Films = make([]filmInCollectionResponse, len(films))

	for idx, value := range films {
		res.Films[idx] = filmInCollectionResponse{
			ID:        value.ID,
			Name:      value.Name,
			YearProd:  value.ProdDate,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	return res
}

func (fcr *FilmCollectionResponse) ToPublic() *FilmCollectionResponse {
	return fcr
}
