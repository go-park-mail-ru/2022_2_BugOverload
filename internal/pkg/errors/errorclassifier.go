package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	// Auth
	ErrEmptyFieldAuth           = stdErrors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = stdErrors.New("no such combination of login and password")
	ErrBadBodyRequest           = stdErrors.New("bad body request")

	ErrUserExist       = stdErrors.New("such user exist")
	ErrUserNotExist    = stdErrors.New("no such user")
	ErrSignupUserExist = stdErrors.New("such a login exists")

	ErrNoCookie        = stdErrors.New("request has no cookies")
	ErrCookieNotExist  = stdErrors.New("no such cookie")
	ErrSessionNotExist = stdErrors.New("no such session")

	ErrGetUserRequest = stdErrors.New("fatal getting user")
	ErrWrongPassword  = stdErrors.New("bad pass")

	// Access
	ErrNoAccess = stdErrors.New("not enough rights")

	// Auth Validaton
	ErrInvalidEmail    = stdErrors.New("invalid email, try another one")
	ErrInvalidPassword = stdErrors.New("invalid password, try another one")

	// Films
	ErrFilmNotFound = stdErrors.New("no such film")

	// Images
	ErrImageNotFound   = stdErrors.New("no such image")
	ErrBadImageType    = stdErrors.New("bad image type")
	ErrGetImageStorage = stdErrors.New("err get data from storage")
	ErrReadImage       = stdErrors.New("err read bin data")
	ErrImage           = stdErrors.New("service picture not work")

	// Def validation
	ErrJSONUnexpectedEnd    = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
	ErrEmptyBody            = stdErrors.New("empty body")
	ErrBigRequest           = stdErrors.New("big request")
	ErrConvertLength        = stdErrors.New("getting content-length failed")
	ErrBigImage             = stdErrors.New("big image")
	ErrConvertQuery         = stdErrors.New("bad input query")
	ErrQueryRequiredEmpty   = stdErrors.New("miss query params")
	ErrQueryBad             = stdErrors.New("bad query params")
	ErrEmptyRequiredFields  = stdErrors.New("bad params, empty")
	ErrBadRequestParams     = stdErrors.New("bad params, impossible value")

	// DB
	ErrPostgresRequest  = stdErrors.New("error sql")
	ErrNotFoundInDB     = stdErrors.New("not found")
	ErrGetParamsConvert = stdErrors.New("err get sql params")

	// Security
	ErrCsrfTokenCreate        = stdErrors.New("csrf token create error")
	ErrCsrfTokenCheck         = stdErrors.New("csrf token check error")
	ErrCsrfTokenCheckInternal = stdErrors.New("csrf token check internal error")
	ErrCsrfTokenExpired       = stdErrors.New("csrf token expired")
	ErrCsrfTokenNotFound      = stdErrors.New("csrf token not found")
	ErrCsrfTokenInvalid       = stdErrors.New("invalid csrf token")
)

type ErrClassifier struct {
	table map[error]int
}

func NewErrClassifier() ErrClassifier {
	res := make(map[error]int)

	// Auth
	// Delivery
	res[ErrEmptyFieldAuth] = http.StatusBadRequest
	res[ErrBadBodyRequest] = http.StatusBadRequest
	res[ErrNoCookie] = http.StatusUnauthorized
	res[ErrQueryRequiredEmpty] = http.StatusBadRequest

	// Service
	res[ErrInvalidEmail] = http.StatusBadRequest
	res[ErrInvalidPassword] = http.StatusBadRequest
	res[ErrLoginCombinationNotFound] = http.StatusForbidden
	res[ErrWrongPassword] = http.StatusForbidden

	// Repository
	res[ErrUserExist] = http.StatusBadRequest
	res[ErrUserNotExist] = http.StatusNotFound
	res[ErrSignupUserExist] = http.StatusBadRequest
	res[ErrCookieNotExist] = http.StatusNotFound
	res[ErrSessionNotExist] = http.StatusNotFound

	// Access
	// Service
	res[ErrNoAccess] = http.StatusForbidden

	// Films
	// Repository
	res[ErrFilmNotFound] = http.StatusNotFound

	// Images
	// Delivery
	res[ErrReadImage] = http.StatusBadRequest
	res[ErrBadImageType] = http.StatusBadRequest

	// Repository
	res[ErrImageNotFound] = http.StatusNotFound
	res[ErrGetImageStorage] = http.StatusBadRequest
	res[ErrImage] = http.StatusInternalServerError

	// Global
	// Def Validation
	res[ErrJSONUnexpectedEnd] = http.StatusBadRequest
	res[ErrContentTypeUndefined] = http.StatusBadRequest
	res[ErrUnsupportedMediaType] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody] = http.StatusBadRequest
	res[ErrBigRequest] = http.StatusBadRequest
	res[ErrConvertLength] = http.StatusBadRequest
	res[ErrConvertQuery] = http.StatusBadRequest
	res[ErrQueryBad] = http.StatusBadRequest
	res[ErrEmptyRequiredFields] = http.StatusBadRequest
	res[ErrBadRequestParams] = http.StatusBadRequest

	// DB
	res[ErrPostgresRequest] = http.StatusInternalServerError
	res[ErrNotFoundInDB] = http.StatusNotFound
	res[ErrGetParamsConvert] = http.StatusInternalServerError

	// Security
	res[ErrCsrfTokenCreate] = http.StatusInternalServerError
	res[ErrCsrfTokenCheck] = http.StatusForbidden
	res[ErrCsrfTokenCheckInternal] = http.StatusInternalServerError
	res[ErrCsrfTokenExpired] = http.StatusForbidden
	res[ErrCsrfTokenNotFound] = http.StatusForbidden
	res[ErrCsrfTokenInvalid] = http.StatusForbidden

	return ErrClassifier{
		table: res,
	}
}

var errCsf = NewErrClassifier()

func GetCode(err error) int {
	code, exist := errCsf.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}
