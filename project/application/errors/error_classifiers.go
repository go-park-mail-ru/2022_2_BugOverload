package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

type ErrClassifier interface {
	GetCode(error) int
}

type ErrClassifierHTTP struct {
	table map[string]int
}

var (
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
)

func NewErrClassifierHTTP() ErrClassifierHTTP {
	res := make(map[string]int)

	res["content-type undefined"] = http.StatusBadRequest
	res["unsupported media type"] = http.StatusUnsupportedMediaType

	return ErrClassifierHTTP{
		table: res,
	}
}

func (ec *ErrClassifierHTTP) GetCode(error error) int {
	code, exist := ec.table[error.Error()]
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

type ErrClassifierAuth struct {
	table map[string]int
}

func NewErrClassifierAuth() ErrClassifierAuth {
	res := make(map[string]int)

	res["request has empty fields (nickname | email | password)"] = http.StatusBadRequest
	res["no such combination of login and password"] = http.StatusUnauthorized

	res["such user exist"] = http.StatusBadRequest
	res["such user doesn't exist"] = http.StatusBadRequest
	res["such a login exists"] = http.StatusBadRequest

	res["request has no cookies"] = http.StatusBadRequest
	res["no such cookie"] = http.StatusUnauthorized

	return ErrClassifierAuth{
		table: res,
	}
}

func (ec *ErrClassifierAuth) GetCode(error error) int {
	code, exist := ec.table[error.Error()]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var (
	ErrFilmNotFound  = stdErrors.New(`no such film`)
	ErrFilmsNotFound = stdErrors.New(`no such films`)
)

type ErrClassifierFilms struct {
	table map[string]int
}

func NewErrClassifierFilms() ErrClassifierFilms {
	res := make(map[string]int)

	res["no such film"] = http.StatusNotFound
	res["no such films"] = http.StatusNotFound

	return ErrClassifierFilms{
		table: res,
	}
}

func (ec *ErrClassifierFilms) GetCode(error error) int {
	code, exist := ec.table[error.Error()]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfHTTP = NewErrClassifierHTTP()
var ErrCsfAuth = NewErrClassifierAuth()
var ErrCsfFilms = NewErrClassifierFilms()
