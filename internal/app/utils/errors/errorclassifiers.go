package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

type errClassifier interface {
	GetCode(error) int
}

type errClassifierDefaultValidation struct {
	table map[error]int
}

var (
	ErrCJSONUnexpectedEnd   = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
)

func NewErrClassifierValidation() errClassifier {
	res := make(map[error]int)

	res[ErrCJSONUnexpectedEnd] = http.StatusBadRequest
	res[ErrContentTypeUndefined] = http.StatusBadRequest
	res[ErrUnsupportedMediaType] = http.StatusUnsupportedMediaType

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

var (
	ErrEmptyFieldAuth           = stdErrors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = stdErrors.New("no such combination of login and password")
	ErrUserExist                = stdErrors.New("such user exist")
	ErrUserNotExist             = stdErrors.New("such user doesn't exist")
	ErrSignupUserExist          = stdErrors.New("such a login exists")
	ErrNoCookie                 = stdErrors.New("request has no cookies")
	ErrCookieNotExist           = stdErrors.New("no such cookie")
)

type errClassifierAuth struct {
	table map[error]int
}

func NewErrClassifierAuth() errClassifier {
	res := make(map[error]int)

	res[ErrEmptyFieldAuth] = http.StatusBadRequest
	res[ErrUserExist] = http.StatusBadRequest
	res[ErrUserNotExist] = http.StatusBadRequest
	res[ErrSignupUserExist] = http.StatusBadRequest

	res[ErrLoginCombinationNotFound] = http.StatusUnauthorized
	res[ErrNoCookie] = http.StatusUnauthorized
	res[ErrCookieNotExist] = http.StatusUnauthorized

	return &errClassifierAuth{
		table: res,
	}
}

func (ec *errClassifierAuth) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var (
	ErrFilmNotFound  = stdErrors.New("no such film")
	ErrFilmsNotFound = stdErrors.New("no such films")
)

type errClassifierFilms struct {
	table map[error]int
}

func NewErrClassifierFilms() errClassifier {
	res := make(map[error]int)

	res[ErrFilmNotFound] = http.StatusNotFound
	res[ErrFilmsNotFound] = http.StatusNotFound

	return &errClassifierFilms{
		table: res,
	}
}

func (ec *errClassifierFilms) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfValid = NewErrClassifierValidation()
var ErrCsfAuth = NewErrClassifierAuth()
var ErrCsfFilms = NewErrClassifierFilms()
