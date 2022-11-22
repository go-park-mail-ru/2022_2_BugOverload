package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	// Common delivery
	ErrBadBodyRequest                      = stdErrors.New("bad body request")
	ErrJSONUnexpectedEnd                   = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined                = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType                = stdErrors.New("unsupported media type")
	ErrEmptyBody                           = stdErrors.New("empty body")
	ErrConvertQueryType                    = stdErrors.New("bad input query")
	ErrQueryRequiredEmpty                  = stdErrors.New("miss query params")
	ErrBadRequestParams                    = stdErrors.New("bad query params")
	ErrBadRequestParamsEmptyRequiredFields = stdErrors.New("bad params, empty required field")

	// Common repository
	ErrNotFoundInDB     = stdErrors.New("not found")
	ErrWorkDatabase     = stdErrors.New("error sql")
	ErrGetParamsConvert = stdErrors.New("err get sql params")

	// Collection service
	ErrNotFindSuchTarget = stdErrors.New("not found such target")

	// Auth delivery
	ErrNoCookie = stdErrors.New("no such cookie")

	// Auth repository
	ErrUserExist    = stdErrors.New("such user exists")
	ErrUserNotExist = stdErrors.New("no such user")

	// Auth service
	ErrInvalidNickname   = stdErrors.New("invalid nickname")
	ErrInvalidEmail      = stdErrors.New("invalid email")
	ErrInvalidPassword   = stdErrors.New("invalid password")
	ErrIncorrectPassword = stdErrors.New("incorrect password")

	// Image delivery
	ErrBigImage     = stdErrors.New("big image")
	ErrBadImageType = stdErrors.New("bad image type")

	// Image repository
	ErrImage = stdErrors.New("service picture not work")

	// User delivery
	ErrGetUserRequest     = stdErrors.New("fatal getting user")
	ErrWrongValidPassword = stdErrors.New("bad pass")

	// User service
	ErrFilmRatingNotExist = stdErrors.New("film rating not exist")

	// Middleware
	ErrBigRequest    = stdErrors.New("big request")
	ErrConvertLength = stdErrors.New("getting content-length failed")

	// Security
	ErrCsrfTokenCreate        = stdErrors.New("csrf token create error")
	ErrCsrfTokenCheck         = stdErrors.New("csrf token check error")
	ErrCsrfTokenCheckInternal = stdErrors.New("csrf token check internal error")
	ErrCsrfTokenExpired       = stdErrors.New("csrf token expired")
	ErrCsrfTokenInvalid       = stdErrors.New("invalid csrf token")

	// Not Found
	ErrGenreNotFount   = stdErrors.New("genre not fount")
	ErrTagNotFount     = stdErrors.New("tag not fount")
	ErrFilmNotFount    = stdErrors.New("film not fount")
	ErrPersonNotFount  = stdErrors.New("person not fount")
	ErrImageNotFound   = stdErrors.New("image not found")
	ErrSessionNotFound = stdErrors.New("session not found")
	ErrUserNotFound    = stdErrors.New("user not found")
)

type ErrClassifier struct {
	table map[error]int
}

func NewErrClassifier() ErrClassifier {
	res := make(map[error]int)

	// Common delivery
	res[ErrBadBodyRequest] = http.StatusBadRequest
	res[ErrJSONUnexpectedEnd] = http.StatusBadRequest
	res[ErrContentTypeUndefined] = http.StatusBadRequest
	res[ErrUnsupportedMediaType] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody] = http.StatusBadRequest
	res[ErrConvertQueryType] = http.StatusBadRequest
	res[ErrQueryRequiredEmpty] = http.StatusBadRequest
	res[ErrBadRequestParams] = http.StatusBadRequest
	res[ErrBadRequestParamsEmptyRequiredFields] = http.StatusBadRequest
	res[ErrBadRequestParams] = http.StatusBadRequest

	// Common repository
	res[ErrNotFoundInDB] = http.StatusNotFound
	res[ErrWorkDatabase] = http.StatusInternalServerError
	res[ErrGetParamsConvert] = http.StatusInternalServerError

	// Collection service
	res[ErrNotFindSuchTarget] = http.StatusNotFound

	// Auth delivery
	res[ErrNoCookie] = http.StatusNotFound

	// Auth repository
	res[ErrUserExist] = http.StatusBadRequest
	res[ErrUserNotExist] = http.StatusNotFound

	// Auth service
	res[ErrInvalidNickname] = http.StatusBadRequest
	res[ErrInvalidEmail] = http.StatusBadRequest
	res[ErrInvalidPassword] = http.StatusBadRequest
	res[ErrIncorrectPassword] = http.StatusForbidden

	// Image delivery
	res[ErrBigImage] = http.StatusBadRequest
	res[ErrBadImageType] = http.StatusBadRequest

	// Image repository
	res[ErrImage] = http.StatusInternalServerError

	// User delivery
	res[ErrGetUserRequest] = http.StatusInternalServerError
	res[ErrWrongValidPassword] = http.StatusForbidden

	// User service
	res[ErrFilmRatingNotExist] = http.StatusNotFound

	// Middleware
	res[ErrBigRequest] = http.StatusBadRequest
	res[ErrConvertLength] = http.StatusBadRequest

	// Security
	res[ErrCsrfTokenCreate] = http.StatusInternalServerError
	res[ErrCsrfTokenCheck] = http.StatusForbidden
	res[ErrCsrfTokenCheckInternal] = http.StatusInternalServerError
	res[ErrCsrfTokenExpired] = http.StatusForbidden
	res[ErrCsrfTokenInvalid] = http.StatusForbidden

	// Not found
	res[ErrGenreNotFount] = http.StatusNotFound
	res[ErrTagNotFount] = http.StatusNotFound
	res[ErrPersonNotFount] = http.StatusNotFound
	res[ErrFilmNotFount] = http.StatusNotFound
	res[ErrSessionNotFound] = http.StatusNotFound
	res[ErrUserNotFound] = http.StatusNotFound
	res[ErrImageNotFound] = http.StatusNotFound

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
