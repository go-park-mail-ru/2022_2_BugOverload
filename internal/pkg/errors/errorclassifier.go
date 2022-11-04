package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	// Auth
	ErrEmptyFieldAuth           = stdErrors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = stdErrors.New("no such combination of login and password")

	ErrUserExist       = stdErrors.New("such user exist")
	ErrUserNotExist    = stdErrors.New("no such user")
	ErrSignupUserExist = stdErrors.New("such a login exists")

	ErrNoCookie        = stdErrors.New("request has no cookies")
	ErrSessionNotExist = stdErrors.New("no such cookie")

	// Films
	ErrFilmNotFound = stdErrors.New("no such film")

	// Images
	ErrImageNotFound   = stdErrors.New("no such image")
	ErrGetImageStorage = stdErrors.New("err get data from storage")
	ErrReadImage       = stdErrors.New("err read bin data")

	// Def validation
	ErrCJSONUnexpectedEnd   = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
	ErrEmptyBody            = stdErrors.New("empty body")
	ErrBigRequest           = stdErrors.New("big request")
	ErrConvertLength        = stdErrors.New("getting content-length failed")
	ErrBigImage             = stdErrors.New("big image")
	ErrConvertQuery         = stdErrors.New("bad input query")
	ErrQueryRequiredEmpty   = stdErrors.New("miss query params")
	ErrQueryBad             = stdErrors.New("bad query params")

	// DB
	ErrPostgresRequest = stdErrors.New("error sql")
	ErrNotFoundInDB    = stdErrors.New("not fount")
)

type ErrClassifier struct {
	table map[error]int
}

func NewErrClassifier() ErrClassifier {
	res := make(map[error]int)

	// Auth
	res[ErrEmptyFieldAuth] = http.StatusBadRequest
	res[ErrUserExist] = http.StatusBadRequest
	res[ErrUserNotExist] = http.StatusNotFound
	res[ErrSignupUserExist] = http.StatusBadRequest

	res[ErrLoginCombinationNotFound] = http.StatusForbidden
	res[ErrNoCookie] = http.StatusUnauthorized
	res[ErrSessionNotExist] = http.StatusNotFound

	// Films
	res[ErrFilmNotFound] = http.StatusNotFound

	// Images
	res[ErrImageNotFound] = http.StatusNotFound

	res[ErrGetImageStorage] = http.StatusBadRequest
	res[ErrReadImage] = http.StatusBadRequest

	// Def Validation
	res[ErrCJSONUnexpectedEnd] = http.StatusBadRequest
	res[ErrContentTypeUndefined] = http.StatusBadRequest
	res[ErrUnsupportedMediaType] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody] = http.StatusBadRequest
	res[ErrBigRequest] = http.StatusBadRequest
	res[ErrConvertLength] = http.StatusBadRequest
	res[ErrConvertQuery] = http.StatusBadRequest
	res[ErrQueryBad] = http.StatusBadRequest

	// DB
	res[ErrPostgresRequest] = http.StatusInternalServerError
	res[ErrNotFoundInDB] = http.StatusNotFound

	return ErrClassifier{
		table: res,
	}
}

func (ec *ErrClassifier) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsf = NewErrClassifier()
