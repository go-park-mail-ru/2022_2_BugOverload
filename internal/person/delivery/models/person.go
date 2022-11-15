package models

import (
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type PersonRequest struct {
	ID          int
	CountFilms  int
	CountImages int
}

func NewPersonRequest() *PersonRequest {
	return &PersonRequest{}
}

func (p *PersonRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	p.ID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	p.CountFilms, err = strconv.Atoi(r.FormValue("count_films"))
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.CountFilms < 0 {
		return errors.ErrBadQueryParams
	}

	p.CountImages, err = strconv.Atoi(r.FormValue("count_images"))
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.CountImages < 0 {
		return errors.ErrBadQueryParams
	}

	return nil
}

func (p *PersonRequest) GetPerson() *models.Person {
	return &models.Person{
		ID: p.ID,
	}
}

func (p *PersonRequest) GetParams() *innerPKG.GetPersonParams {
	return &innerPKG.GetPersonParams{
		CountFilms:  p.CountFilms,
		CountImages: p.CountImages,
	}
}

type filmInPersonResponse struct {
	ID        int      `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  string   `json:"prod_year,omitempty" example:"2014"`
	EndYear   string   `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type PersonResponse struct {
	Name         string                 `json:"name,omitempty" example:"Шон Коннери"`
	OriginalName string                 `json:"original_name,omitempty" example:"Sean Connery"`
	Birthday     string                 `json:"birthday,omitempty" example:"1930.08.25"`
	Death        string                 `json:"death,omitempty" example:"2020.10.31"`
	Growth       float32                `json:"growth,omitempty" example:"1.9"`
	Gender       string                 `json:"gender,omitempty" example:"male"`
	Avatar       string                 `json:"avatar,omitempty" example:"4526"`
	CountFilms   int                    `json:"count_films,omitempty" example:"218"`
	Professions  []string               `json:"professions,omitempty" example:"актер,продюсер,режиссер"`
	Genres       []string               `json:"genres,omitempty" example:"драма,боевик,триллер"`
	BestFilms    []filmInPersonResponse `json:"best_films,omitempty"`

	Images []string `json:"images,omitempty" example:"1,2,3,4,5,6,7"`
}

func NewPersonResponse(person *models.Person) *PersonResponse {
	res := &PersonResponse{
		Name:         person.Name,
		Birthday:     person.Birthday,
		OriginalName: person.OriginalName,
		Death:        person.Death,
		Growth:       person.Growth,
		Gender:       person.Gender,
		CountFilms:   person.CountFilms,
		Professions:  person.Professions,
		Genres:       person.Genres,
		Avatar:       person.Avatar,
		Images:       person.Images,
		BestFilms:    make([]filmInPersonResponse, len(person.BestFilms)),
	}

	for idx, value := range person.BestFilms {
		res.BestFilms[idx] = filmInPersonResponse{
			ID:        value.ID,
			Name:      value.Name,
			ProdYear:  value.ProdYear,
			EndYear:   value.EndYear,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	return res
}
