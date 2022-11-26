package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type PostImageRequest struct {
	Key    string `json:"key" example:"1"`
	Object string `json:"object" example:"film_poster_hor"`
	Bytes  []byte `json:"-"`
}

func NewPostImageRequest() *PostImageRequest {
	return &PostImageRequest{}
}

func (i *PostImageRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != constparams.ContentTypeWEBP || r.Header.Get("Content-Type") != constparams.ContentTypeJPEG {
		return errors.ErrUnsupportedMediaType
	}

	i.Key = r.FormValue("key")
	i.Object = r.FormValue("object")

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

	if len(body) > constparams.BufSizeImage {
		return errors.ErrBigImage
	}

	i.Bytes = body

	return nil
}

func (i *PostImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
		Bytes:  i.Bytes,
	}
}