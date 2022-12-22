package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields filmgetsimilar.go

//easyjson:json
type GetSimilarFilmsRequest struct {
	FilmID int `json:"film_id,omitempty" example:"23"`
}

func NewGetSimilarFilmsRequest() *GetSimilarFilmsRequest {
	return &GetSimilarFilmsRequest{}
}

func (p *GetSimilarFilmsRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	p.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	return nil
}

func (p *GetSimilarFilmsRequest) GetParams() *constparams.GetSimilarFilmsParams {
	return &constparams.GetSimilarFilmsParams{
		FilmID: p.FilmID,
	}
}

//easyjson:json
type SimilarFilmCollectionResponse struct {
	ID        int      `json:"id,omitempty" example:"23"`
	Name      string   `json:"name,omitempty" example:"Game of Thrones"`
	ProdYear  string   `json:"prod_year,omitempty" example:"2014"`
	EndYear   string   `json:"end_year,omitempty" example:"2013"`
	PosterVer string   `json:"poster_ver,omitempty" example:"{{key}}"`
	Rating    float32  `json:"rating,omitempty" example:"7.9523542"`
	Genres    []string `json:"genres,omitempty" example:"фэнтези,приключения"`
}

//easyjson:json
type GetSimilarFilmsCollectionResponse struct {
	Name  string                          `json:"name,omitempty" example:"Похожие фильмы"`
	Films []SimilarFilmCollectionResponse `json:"films,omitempty"`
}

func NewGetSimilarFilmsCollectionResponse(collection *models.Collection) *GetSimilarFilmsCollectionResponse {
	res := &GetSimilarFilmsCollectionResponse{
		Name:  collection.Name,
		Films: make([]SimilarFilmCollectionResponse, len(collection.Films)),
	}

	for idx, value := range collection.Films {
		res.Films[idx] = SimilarFilmCollectionResponse{
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
