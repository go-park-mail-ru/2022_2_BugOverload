package errors

import (
	"fmt"
	"net/http"

	stdErrors "github.com/pkg/errors"
)

type errClassifierDefaultValidation struct {
	table map[error]int
}

var (
	ErrCJSONUnexpectedEnd   = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
	ErrEmptyBody            = stdErrors.New("empty body")
	ErrBigRequest           = stdErrors.New("big request")
	ErrConvertLength        = stdErrors.New("getting content-length failed")
)

func NewErrClassifierValidation() errClassifier {
	res := make(map[error]int)

	res[ErrCJSONUnexpectedEnd] = http.StatusBadRequest
	res[ErrContentTypeUndefined] = http.StatusBadRequest
	res[ErrUnsupportedMediaType] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody] = http.StatusBadRequest
	res[ErrBigRequest] = http.StatusBadRequest
	res[ErrConvertLength] = http.StatusBadRequest

	return &errClassifierDefaultValidation{
		table: res,
	}
}

func (ec *errClassifierDefaultValidation) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfValid = NewErrClassifierValidation()

type DefaultValidationError struct {
	Reason string
	Code   int
}

func (e DefaultValidationError) Error() string {
	return fmt.Sprintf("Def validation: [%s]", e.Reason)
}

func NewErrValidation(err error) DefaultValidationError {
	return DefaultValidationError{
		Reason: err.Error(),
		Code:   ErrCsfValid.GetCode(err),
	}
}
