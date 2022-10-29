package errors

import (
	"fmt"

	stdErrors "github.com/pkg/errors"

	"net/http"
)

var (
	ErrImageNotFound   = stdErrors.New("no such image")
	ErrGetImageStorage = stdErrors.New("err get data from storage")
	ErrReadImage       = stdErrors.New("err read bin data")
)

type errClassifierImages struct {
	table map[error]int
}

func NewErrClassifierImages() errClassifier {
	res := make(map[error]int)

	res[ErrImageNotFound] = http.StatusNotFound

	res[ErrGetImageStorage] = http.StatusBadRequest
	res[ErrReadImage] = http.StatusBadRequest

	return &errClassifierImages{
		table: res,
	}
}

func (ec *errClassifierImages) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfImages = NewErrClassifierImages()

type ImagesError struct {
	Reason string
	Code   int
}

func (e ImagesError) Error() string {
	return fmt.Sprintf("Image: [%s]", e.Reason)
}

func NewErrImages(err error) ImagesError {
	return ImagesError{
		Reason: err.Error(),
		Code:   ErrCsfImages.GetCode(err),
	}
}
