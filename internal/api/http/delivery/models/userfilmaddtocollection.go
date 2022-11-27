package models

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type AddFilmToUserCollectionRequest struct {
	FilmID       int
	CollectionID int `json:"idCollection,omitempty" example:"4"`
}

func NewAddFilmRequest() *AddFilmToUserCollectionRequest {
	return &AddFilmToUserCollectionRequest{}
}

func (p *AddFilmToUserCollectionRequest) Bind(r *http.Request) error {
	var err error

	vars := mux.Vars(r)

	p.FilmID, err = strconv.Atoi(vars["id"])
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

	err = json.Unmarshal(body, p)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	return nil
}

func (p *AddFilmToUserCollectionRequest) GetParams() *constparams.UserCollectionFilmsUpdateParams {
	return &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       p.FilmID,
		CollectionID: p.CollectionID,
	}
}
