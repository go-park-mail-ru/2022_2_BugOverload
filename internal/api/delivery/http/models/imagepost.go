package models

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
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

	if !(strings.Contains(r.Header.Get("Content-Type"), constparams.ContentTypeMultipartFormData)) {
		return errors.ErrUnsupportedMediaType
	}

	i.Key = r.FormValue("key")
	i.Object = r.FormValue("object")

	errParse := r.ParseMultipartForm(constparams.BufSizeImage)
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

	// contentType := http.DetectContentType(fileHeader) more complicated way (ignoring Headers multipart value)
	contentType := multipartFileHeader.Header.Get("Content-type")
	if !(contentType == constparams.ContentTypeJPEG || contentType == constparams.ContentTypeWEBP || contentType == constparams.ContentTypePNG) {
		return errors.ErrContentTypeUndefined
	}

	body := bytes.NewBuffer(nil)

	_, errCopy := io.Copy(body, file)
	if errCopy != nil {
		return errors.ErrBadBodyRequest
	}

	i.Bytes = body.Bytes()

	return nil
}

func (i *PostImageRequest) GetImage() *models.Image {
	return &models.Image{
		Object: i.Object,
		Key:    i.Key,
		Bytes:  i.Bytes,
	}
}
