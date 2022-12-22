package models

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields userfilmdropfromcollection.go

//easyjson:json
type DropFilmFromUserCollectionRequest struct {
	FilmID       int `json:"-"`
	CollectionID int `json:"collection_id,omitempty" example:"4"`
}

func NewDropFilmRequest() *DropFilmFromUserCollectionRequest {
	return &DropFilmFromUserCollectionRequest{}
}

func (p *DropFilmFromUserCollectionRequest) Bind(r *http.Request) error {
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

	err = easyjson.Unmarshal(body, p)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	return nil
}

func (p *DropFilmFromUserCollectionRequest) GetParams() *constparams.UserCollectionFilmsUpdateParams {
	return &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       p.FilmID,
		CollectionID: p.CollectionID,
	}
}
