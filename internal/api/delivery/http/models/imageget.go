package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate easyjson  -disallow_unknown_fields imageget.go

//easyjson:json
type GetImageRequest struct {
	Key    string `json:"key" example:"1"`
	Object string `json:"object" example:"film_poster_hor"`
}

func NewGetImageRequest() *GetImageRequest {
	return &GetImageRequest{}
}

func (i *GetImageRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		return errors.ErrUnsupportedMediaType
	}

	i.Key = r.FormValue("key")
	i.Object = r.FormValue("object")

	return nil
}

func (i *GetImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
	}
}
