package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type FilmRateRequest struct {
	FilmID int     `json:"-"`
	Score  float32 `json:"score,omitempty" example:"4.2"`
}

func NewFilmRateRequest() FilmRateRequest {
	return FilmRateRequest{}
}

func (f *FilmRateRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	f.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.NewErrValidation(errors.ErrConvertQuery)
	}

	return nil
}

func (f *FilmRateRequest) GetParams() *innerPKG.FilmRateParams {
	return &innerPKG.FilmRateParams{
		FilmID: f.FilmID,
	}
}

func (f *FilmRateRequest) GetFilm() *models.Film {
	return &models.Film{
		ID: f.FilmID,
	}
}
