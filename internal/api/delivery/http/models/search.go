package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields search.go

type SearchRequest struct {
	query string
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{}
}

func (p *SearchRequest) Bind(r *http.Request) error {
	p.query = r.FormValue("q")
	if p.query == "" {
		return errors.ErrQueryRequiredEmpty
	}

	return nil
}

func (p *SearchRequest) GetParams() *constparams.SearchParams {
	return &constparams.SearchParams{
		Query: p.query,
	}
}

//easyjson:json
type SearchFilmPersonResponse struct {
	ID   int    `json:"id,omitempty" example:"123123"`
	Name string `json:"name,omitempty" example:"Стивен Спилберг"`
}

//easyjson:json
type SearchFilmResponse struct {
	ID           int     `json:"id,omitempty" example:"23"`
	Name         string  `json:"name,omitempty" example:"Игра престолов"`
	OriginalName string  `json:"original_name,omitempty" example:"Game of Thrones"`
	ProdYear     string  `json:"prod_date,omitempty" example:"2014.01.13"` // WARNING json name <> go name, must be prod_year in both
	PosterVer    string  `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating       float32 `json:"rating,omitempty" example:"9.2"`

	Genres        []string                   `json:"genres,omitempty" example:"фэнтези,приключения"`
	ProdCountries []string                   `json:"prod_countries,omitempty" example:"США,Великобритания"`
	Directors     []SearchFilmPersonResponse `json:"directors,omitempty"`

	EndYear string `json:"end_year,omitempty" example:"2019.01.13"`
}

//easyjson:json
type SearchPersonResponse struct {
	ID           int      `json:"id,omitempty" example:"23"`
	Name         string   `json:"name,omitempty" example:"Шон Коннери"`
	OriginalName string   `json:"original_name,omitempty" example:"Sean Connery"`
	Birthday     string   `json:"birthday,omitempty" example:"1930.08.25"`
	Avatar       string   `json:"avatar,omitempty" example:"4526"`
	CountFilms   int      `json:"count_films,omitempty" example:"218"`
	Professions  []string `json:"professions,omitempty" example:"актер,продюсер,режиссер"`
}

//easyjson:json
type SearchResponse struct {
	Films   []SearchFilmResponse   `json:"films,omitempty"`
	Series  []SearchFilmResponse   `json:"serials,omitempty"`
	Persons []SearchPersonResponse `json:"persons,omitempty"`
}

func newSearchFilmPersonsResponse(filmPersons []models.FilmPerson) []SearchFilmPersonResponse {
	res := make([]SearchFilmPersonResponse, len(filmPersons))

	for idx, val := range filmPersons {
		res[idx].ID = val.ID
		res[idx].Name = val.Name
	}

	return res
}

func NewSearchResponse(resp models.Search) (SearchResponse, error) {
	if len(resp.Films) == 0 && len(resp.Serials) == 0 && len(resp.Persons) == 0 {
		return SearchResponse{}, errors.ErrNotFoundInDB
	}

	films := make([]SearchFilmResponse, len(resp.Films))

	for idx, val := range resp.Films {
		films[idx].ID = val.ID
		films[idx].Name = val.Name
		films[idx].OriginalName = val.OriginalName
		films[idx].ProdYear = val.ProdDate
		films[idx].PosterVer = val.PosterVer
		films[idx].Genres = val.Genres
		films[idx].ProdCountries = val.ProdCountries
		films[idx].Directors = newSearchFilmPersonsResponse(val.Directors)
	}

	serials := make([]SearchFilmResponse, len(resp.Serials))

	for idx, val := range resp.Serials {
		serials[idx].ID = val.ID
		serials[idx].Name = val.Name
		serials[idx].OriginalName = val.OriginalName
		serials[idx].ProdYear = val.ProdDate
		serials[idx].PosterVer = val.PosterVer
		serials[idx].Genres = val.Genres
		serials[idx].ProdCountries = val.ProdCountries
		serials[idx].Directors = newSearchFilmPersonsResponse(val.Directors)

		serials[idx].EndYear = val.EndYear
	}

	persons := make([]SearchPersonResponse, len(resp.Persons))

	for idx, val := range resp.Persons {
		persons[idx].ID = val.ID
		persons[idx].Name = val.Name
		persons[idx].OriginalName = val.OriginalName
		persons[idx].Birthday = val.Birthday
		persons[idx].Avatar = val.Avatar
		persons[idx].CountFilms = val.CountFilms
		persons[idx].Professions = val.Professions
	}

	return SearchResponse{
		Films:   films,
		Series:  serials,
		Persons: persons,
	}, nil
}
