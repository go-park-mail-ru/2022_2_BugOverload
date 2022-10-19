package models

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type GetImageRequest struct {
	Bucket string `json:"bucket" example:"films/"`
	Item   string `json:"item" example:"posters/hor/1.jpg"`
}

func NewGetImageRequest() *GetImageRequest {
	return &GetImageRequest{}
}

func (i *GetImageRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.NewErrValidation(errors.ErrContentTypeUndefined)
	}

	if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	err = json.Unmarshal(body, i)
	if err != nil {
		return errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd)
	}

	return nil
}

func (i *GetImageRequest) GetImage() *models.Image {
	return &models.Image{
		Bucket: i.Bucket,
		Item:   i.Item,
	}
}
