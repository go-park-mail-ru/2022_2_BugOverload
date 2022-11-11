package models

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"io"
	"net/http"
	"strings"
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
	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if !(strings.Contains(r.Header.Get("Content-Type"), pkg.ContentTypeMultipartFormData)) {
		return errors.ErrUnsupportedMediaType
	}

	i.Key = r.FormValue("key")
	i.Object = r.FormValue("object")

	errParse := r.ParseMultipartForm(pkg.BufSizeImage)
	if errParse != nil {
		return errors.ErrBigImage
	}

	file, _, err := r.FormFile("object")
	logrus.Info(file)
	if err != nil {
		return errors.ErrEmptyBody
	}

	body := bytes.NewBuffer(nil)
	if _, errCopy := io.Copy(body, file); errCopy != nil {
		return errors.ErrBadBodyRequest
	}

	i.Bytes = body.Bytes()

	return nil
}

func (i *PutImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
		Bytes:  i.Bytes,
	}
}
