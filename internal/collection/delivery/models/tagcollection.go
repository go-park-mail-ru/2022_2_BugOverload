package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type TagCollectionRequest struct {
	Tag        string
	CountFilms int
	Delimiter  string
}

func NewTagCollectionRequest() *TagCollectionRequest {
	return &TagCollectionRequest{}
}

func (p *TagCollectionRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	var err error
	p.Tag = vars["tag"]

	countFilms := r.FormValue("count_films")

	p.Delimiter = r.FormValue("delimiter")

	if countFilms == "" || p.Tag == "" || p.Delimiter == "" {
		return errors.NewErrValidation(errors.ErrQueryRequiredEmpty)
	}

	switch p.Tag {
	case innerPKG.TagFromPopular:
		p.Tag = innerPKG.TagInPopular
	case innerPKG.TagFromInCinema:
		p.Tag = innerPKG.TagInInCinema
	}

	p.CountFilms, err = strconv.Atoi(countFilms)
	if err != nil {
		return errors.NewErrValidation(errors.ErrConvertQuery)
	}

	if p.CountFilms <= 0 {
		return errors.NewErrValidation(errors.ErrQueryRequiredEmpty)
	}

	return nil
}

func (p *TagCollectionRequest) GetParams() *innerPKG.GetCollectionTagParams {
	return &innerPKG.GetCollectionTagParams{
		Tag:        p.Tag,
		CountFilms: p.CountFilms,
		Delimiter:  p.Delimiter,
	}
}

type filmTagCollectionResponse struct {
	ID        int      `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  int      `json:"prod_year,omitempty" example:"2014"`
	EndYear   int      `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type TagCollectionResponse struct {
	Name  string                      `json:"name,omitempty" example:"Сейчас в кино"`
	Films []filmTagCollectionResponse `json:"films,omitempty"`
}

func NewTagCollectionResponse(collection *models.Collection) *TagCollectionResponse {
	res := &TagCollectionResponse{
		Name:  collection.Name,
		Films: make([]filmTagCollectionResponse, len(collection.Films)),
	}

	for idx, value := range collection.Films {
		res.Films[idx] = filmTagCollectionResponse{
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
