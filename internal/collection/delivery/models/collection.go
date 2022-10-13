package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type filmInCollectionRequest struct {
	ID               uint     `json:"film_id,omitempty"`
	Name             string   `json:"film_name,omitempty"`
	ShortDescription string   `json:"short_description,omitempty"`
	YearProd         string   `json:"year_prod,omitempty"`
	PosterVer        string   `json:"poster_ver,omitempty"`
	Rating           string   `json:"ratio,omitempty"`
	Genres           []string `json:"genres,omitempty"`
}

type FilmCollectionRequest struct {
	Title string                    `json:"title,omitempty"`
	Films []filmInCollectionRequest `json:"films,omitempty"`
}

func NewFilmCollectionRequest(title string, films []models.Film) *FilmCollectionRequest {
	res := &FilmCollectionRequest{
		Title: title,
	}

	res.Films = make([]filmInCollectionRequest, len(films))

	for idx, value := range films {
		res.Films[idx] = filmInCollectionRequest{
			ID:               value.ID,
			Name:             value.Name,
			ShortDescription: value.ShortDescription,
			YearProd:         value.YearProd,
			PosterVer:        value.PosterVer,
			Rating:           value.Rating,
			Genres:           value.Genres,
		}
	}

	return res
}

func (fcr *FilmCollectionRequest) ToPublic() *FilmCollectionRequest {
	return fcr
}
