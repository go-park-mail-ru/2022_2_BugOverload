package models

import (
	"net/http"
	"strconv"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type GetStdCollectionRequest struct {
	Target     string
	Key        string
	SortParam  string
	SortDir    string
	CountFilms int
	Delimiter  string
}

func NewGetStdCollectionRequest() *GetStdCollectionRequest {
	return &GetStdCollectionRequest{}
}

func (p *GetStdCollectionRequest) Bind(r *http.Request) error {
	var err error

	p.Target = r.FormValue("target")
	if p.Target == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	p.Key = r.FormValue("key")
	if p.Key == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	p.SortParam = r.FormValue("sort_param")
	if p.SortParam == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	countFilms := r.FormValue("count_films")
	if countFilms == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	p.CountFilms, err = strconv.Atoi(countFilms)
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if p.CountFilms < 0 {
		return errors.ErrBadRequestParams
	}

	p.Delimiter = r.FormValue("delimiter")
	if p.Delimiter == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (p *GetStdCollectionRequest) GetParams() *innerPKG.GetCollectionParams {
	return &innerPKG.GetCollectionParams{
		Key:        p.Key,
		Target:     p.Target,
		SortParam:  p.SortParam,
		CountFilms: p.CountFilms,
		Delimiter:  p.Delimiter,
	}
}

type FilmTagCollectionResponse struct {
	ID        int      `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  string   `json:"prod_year,omitempty" example:"2014"`
	EndYear   string   `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9523542"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

type GetStdCollectionResponse struct {
	Name  string                      `json:"name,omitempty" example:"Сейчас в кино"`
	Films []FilmTagCollectionResponse `json:"films,omitempty"`
}

func NewStdCollectionResponse(collection *models.Collection) *GetStdCollectionResponse {
	res := &GetStdCollectionResponse{
		Name:  collection.Name,
		Films: make([]FilmTagCollectionResponse, len(collection.Films)),
	}

	for idx, value := range collection.Films {
		res.Films[idx] = FilmTagCollectionResponse{
			ID:        value.ID,
			Name:      value.Name,
			ProdYear:  value.ProdDate,
			EndYear:   value.EndYear,
			PosterVer: value.PosterVer,
			Rating:    value.Rating,
			Genres:    value.Genres,
		}
	}

	return res
}
