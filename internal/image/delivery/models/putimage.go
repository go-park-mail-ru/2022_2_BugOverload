package models

import (
	"bytes"
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

	file, multipartFileHeader, err := r.FormFile("object")
	if err != nil {
		return errors.ErrEmptyBody
	}

	fileHeader := make([]byte, multipartFileHeader.Size)
	if _, errRead := file.Read(fileHeader); errRead != nil {
		return errors.ErrBadBodyRequest
	}

	if _, errSeek := file.Seek(0, io.SeekStart); errSeek != nil {
		return errors.ErrBadBodyRequest
	}

	contentType := http.DetectContentType(fileHeader)
	if !(contentType == pkg.ContentTypeJPEG || contentType == pkg.ContentTypeWEBP) {
		return errors.ErrContentTypeUndefined
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
