package models

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

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

	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != innerPKG.ContentTypeJSON {
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

	err = json.Unmarshal(body, f)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	return nil
}

func (f *FilmRateRequest) GetParams() *innerPKG.FilmRateParams {
	return &innerPKG.FilmRateParams{
		FilmID: f.FilmID,
		Score:  f.Score,
	}
}
