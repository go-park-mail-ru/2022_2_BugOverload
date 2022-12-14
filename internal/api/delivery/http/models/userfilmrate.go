package models

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields userfilmrate.go

//easyjson:json
type FilmRateRequest struct {
	FilmID int `json:"-"`
	Score  int `json:"score,omitempty" example:"4"`
}

func NewFilmRateRequest() FilmRateRequest {
	return FilmRateRequest{}
}

func (f *FilmRateRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	f.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != constparams.ContentTypeJSON {
		return errors.ErrUnsupportedMediaType
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrBadBodyRequest
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if len(body) == 0 {
		return errors.ErrEmptyBody
	}

	err = easyjson.Unmarshal(body, f)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	return nil
}

func (f *FilmRateRequest) GetParams() *constparams.FilmRateParams {
	return &constparams.FilmRateParams{
		FilmID: f.FilmID,
		Score:  f.Score,
	}
}

//easyjson:json
type FilmRateResponse struct {
	Rating       float32 `json:"rating,omitempty" example:"9.2"`
	CountRatings int     `json:"count_ratings,omitempty" example:"22"`
}

func NewFilmRateResponse(film *models.Film) *FilmRateResponse {
	return &FilmRateResponse{
		Rating:       film.Rating,
		CountRatings: film.CountRatings,
	}
}
