package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type FilmRateDropRequest struct {
	FilmID int `json:"film_id,omitempty" example:"23"`
}

func NewFilmRateDropRequest() FilmRateDropRequest {
	return FilmRateDropRequest{}
}

func (f *FilmRateDropRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	f.FilmID, err = strconv.Atoi(vars["id"])
	if err != nil {
		return errors.ErrConvertQueryType
	}

	return nil
}

func (f *FilmRateDropRequest) GetParams() *innerPKG.FilmRateDropParams {
	return &innerPKG.FilmRateDropParams{
		FilmID: f.FilmID,
	}
}

func (f *FilmRateDropRequest) GetFilm() *models.Film {
	return &models.Film{
		ID: f.FilmID,
	}
}

type FilmRateDropResponse struct {
	CountRatings int `json:"count_scores,omitempty" example:"786442"`
}

func NewFilmRateDropResponse(film *models.Film) *FilmRateDropResponse {
	return &FilmRateDropResponse{
		CountRatings: film.CountRatings,
	}
}
