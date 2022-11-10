package models

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type PutImageRequest struct {
	Key    string `json:"key" example:"23"`
	Object string `json:"object" example:"user_avatar"`
	Bytes  []byte `json:"-"`
}

func NewPutImageRequest() *PutImageRequest {
	return &PutImageRequest{}
}

func (i *PutImageRequest) Bind(r *http.Request) error {
	/*	if r.Header.Get("Content-Type") == "" {
			return errors.ErrContentTypeUndefined
		}

		if !(r.Header.Get("Content-Type") == pkg.ContentTypeWEBP || r.Header.Get("Content-Type") == pkg.ContentTypeJPEG) {
			return errors.ErrUnsupportedMediaType
		}*/

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

	if len(body) > pkg.BufSizeImage {
		return errors.ErrBigImage
	}

	i.Bytes = body

	return nil
}

func (i *PutImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
		Bytes:  i.Bytes,
	}
}
