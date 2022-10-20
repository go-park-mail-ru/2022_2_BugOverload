package models

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UploadImageRequest struct {
	Key    string `json:"key" example:"1"`
	Object string `json:"object" example:"film_hor"`
	Bytes  []byte `json:"-"`
}

func NewUploadImageRequest() *UploadImageRequest {
	return &UploadImageRequest{}
}

func (i *UploadImageRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") != pkg.ContentTypeJPEG {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	i.Key = r.FormValue("key")
	i.Object = r.FormValue("object")

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

	if len(body) == 0 {
		return errors.NewErrValidation(errors.ErrEmptyBody)
	}

	i.Bytes = body

	return nil
}

func (i *UploadImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
		Bytes:  i.Bytes,
	}
}
