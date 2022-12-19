package models

import (
	"net/http"
	"strconv"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type PremieresCollectionRequest struct {
	CountFilms int
	Delimiter  int
}

func NewPremieresCollectionRequest() *PremieresCollectionRequest {
	return &PremieresCollectionRequest{}
}

func (p *PremieresCollectionRequest) Bind(r *http.Request) error {
	countFilms := r.FormValue("count_films")
	if countFilms == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	var err error
	p.CountFilms, err = strconv.Atoi(countFilms)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.CountFilms < 0 {
		return errors.ErrBadRequestParams
	}

	delimiter := r.FormValue("delimiter")
	if delimiter == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	p.Delimiter, err = strconv.Atoi(delimiter)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.Delimiter < 0 {
		return errors.ErrBadRequestParams
	}

	return nil
}

func (p *PremieresCollectionRequest) GetParams() *innerPKG.GetPremiersCollectionParams {
	return &innerPKG.GetPremiersCollectionParams{
		CountFilms: p.CountFilms,
		Delimiter:  p.Delimiter,
	}
}

type FilmPersonPremiersResponse struct {
	ID   int    `json:"id,omitempty" example:"123123"`
	Name string `json:"name,omitempty" example:"Стивен Спилберг"`
}

type PremieresCollectionFilm struct {
	ID              int     `json:"id,omitempty" example:"23"`
	Name            string  `json:"name,omitempty" example:"Игра престолов"`
	ProdDate        string  `json:"prod_date,omitempty" example:"2014.01.13"`
	PosterVer       string  `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating          float32 `json:"rating,omitempty" example:"9.2"`
	DurationMinutes int     `json:"duration_minutes,omitempty" example:"55"`
	Description     string  `json:"description,omitempty" example:"Британская лингвистка Алетея прилетает из Лондона"`

	Genres        []string                     `json:"genres,omitempty" example:"фэнтези,приключения"`
	ProdCountries []string                     `json:"prod_countries,omitempty" example:"США,Великобритания"`
	Directors     []FilmPersonPremiersResponse `json:"directors,omitempty"`
}

type PremieresCollectionResponse struct {
	Name        string                    `json:"name,omitempty" example:"Сейчас в кино"`
	Description string                    `json:"description,omitempty" example:"Здесь вы можете посмотреть новинки кинопроката"`
	Films       []PremieresCollectionFilm `json:"films,omitempty"`
}

func NewPremieresCollectionResponse(collection *models.Collection) *PremieresCollectionResponse {
	res := &PremieresCollectionResponse{
		Name:  collection.Name,
		Films: make([]PremieresCollectionFilm, len(collection.Films)),
	}

	for idx := range collection.Films {
		res.Films[idx].ID = collection.Films[idx].ID
		res.Films[idx].Name = collection.Films[idx].Name
		res.Films[idx].ProdDate = collection.Films[idx].ProdDate
		res.Films[idx].PosterVer = collection.Films[idx].PosterVer
		res.Films[idx].Rating = collection.Films[idx].Rating
		res.Films[idx].DurationMinutes = collection.Films[idx].DurationMinutes
		res.Films[idx].Description = collection.Films[idx].Description

		res.Films[idx].Genres = make([]string, 0)
		for index, val := range collection.Films[idx].Genres {
			res.Films[idx].Genres = append(res.Films[idx].Genres, val)
			if index > 0 {
				break
			}
		}

		res.Films[idx].ProdCountries = make([]string, 0)
		for index, val := range collection.Films[idx].ProdCountries {
			res.Films[idx].ProdCountries = append(res.Films[idx].ProdCountries, val)
			if index > 0 {
				break
			}
		}

		if len(collection.Films[idx].Directors) == 0 {
			continue
		}

		res.Films[idx].Directors = make([]FilmPersonPremiersResponse, 1)
		res.Films[idx].Directors[0].ID = collection.Films[idx].Directors[0].ID
		res.Films[idx].Directors[0].Name = collection.Films[idx].Directors[0].Name
	}

	return res
}
