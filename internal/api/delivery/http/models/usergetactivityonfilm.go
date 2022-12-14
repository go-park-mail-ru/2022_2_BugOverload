package models

import (
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields usergetactivityonfilm.go

type GetUserActivityOnFilmRequest struct {
	FilmID int
}

func NewUserActivityOnFilmRequest() GetUserActivityOnFilmRequest {
	return GetUserActivityOnFilmRequest{}
}

func (f *GetUserActivityOnFilmRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		return errors.ErrUnsupportedMediaType
	}

	var err error

	vars := mux.Vars(r)

	f.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	return nil
}

func (f *GetUserActivityOnFilmRequest) GetParams() *innerPKG.GetUserActivityOnFilmParams {
	return &innerPKG.GetUserActivityOnFilmParams{
		FilmID: f.FilmID,
	}
}

//easyjson:json
type NodeInUserCollectionResponse struct {
	ID     int    `json:"id,omitempty" example:"123"`
	Name   string `json:"name,omitempty" example:"Избранное"`
	IsUsed bool   `json:"is_used,omitempty" example:"true"`
}

//easyjson:json
type GetUserActivityOnFilmResponse struct {
	CountReviews int                            `json:"count_reviews,omitempty" example:"44"`
	Rating       int                            `json:"rating,omitempty" example:"5"`
	DateRating   string                         `json:"date_rating,omitempty" example:"2022.12.29"`
	Collections  []NodeInUserCollectionResponse `json:"collections,omitempty"`
}

func NewGetUserActivityOnFilmResponse(userActivity *models.UserActivity) *GetUserActivityOnFilmResponse {
	res := &GetUserActivityOnFilmResponse{
		CountReviews: userActivity.CountReviews,
		Rating:       userActivity.Rating,
		DateRating:   userActivity.DateRating,
		Collections:  make([]NodeInUserCollectionResponse, len(userActivity.Collections)),
	}

	for idx, value := range userActivity.Collections {
		res.Collections[idx].ID = value.ID
		res.Collections[idx].Name = value.Name
		res.Collections[idx].IsUsed = value.IsUsed
	}

	return res
}
