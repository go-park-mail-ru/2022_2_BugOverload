package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	// Common delivery
	ErrBadBodyRequest       = stdErrors.New("bad body request")
	ErrJSONUnexpectedEnd    = stdErrors.New("unexpected end of JSON input")
	ErrContentTypeUndefined = stdErrors.New("content-type undefined")
	ErrUnsupportedMediaType = stdErrors.New("unsupported media type")
	ErrEmptyBody            = stdErrors.New("empty body")
	ErrConvertQuery         = stdErrors.New("bad input query")
	ErrQueryRequiredEmpty   = stdErrors.New("miss query params")
	ErrQueryBad             = stdErrors.New("bad query params")
	ErrEmptyField           = stdErrors.New("empty field")
	ErrEmptyRequiredFields  = stdErrors.New("bad params, empty")
	ErrBadRequestParams     = stdErrors.New("bad params, impossible value")

	// Common repository
	ErrNotFoundInDB     = stdErrors.New("not found")
	ErrPostgresRequest  = stdErrors.New("error sql")
	ErrGetParamsConvert = stdErrors.New("err get sql params")

	// Auth delivery
	ErrNoCookie        = stdErrors.New("no such cookie")
	ErrSessionNotExist = stdErrors.New("no such session")

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
	ErrImageNotFound = stdErrors.New("no such image")
	ErrImage         = stdErrors.New("service picture not work")

	// User delivery
	ErrGetUserRequest     = stdErrors.New("fatal getting user")
	ErrWrongValidPassword = stdErrors.New("bad pass")

	// Middleware
	ErrBigRequest    = stdErrors.New("big request")
	ErrConvertLength = stdErrors.New("getting content-length failed")

	// Security
	ErrCsrfTokenCreate        = stdErrors.New("csrf token create error")
	ErrCsrfTokenCheck         = stdErrors.New("csrf token check error")
	ErrCsrfTokenCheckInternal = stdErrors.New("csrf token check internal error")
	ErrCsrfTokenExpired       = stdErrors.New("csrf token expired")
	ErrCsrfTokenInvalid       = stdErrors.New("invalid csrf token")
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
	res[ErrConvertQuery] = http.StatusBadRequest
	res[ErrQueryRequiredEmpty] = http.StatusBadRequest
	res[ErrQueryBad] = http.StatusBadRequest
	res[ErrEmptyField] = http.StatusBadRequest
	res[ErrEmptyRequiredFields] = http.StatusBadRequest
	res[ErrBadRequestParams] = http.StatusBadRequest

	// Common repository
	res[ErrNotFoundInDB] = http.StatusNotFound
	res[ErrPostgresRequest] = http.StatusInternalServerError
	res[ErrGetParamsConvert] = http.StatusInternalServerError

	// Auth delivery
	res[ErrNoCookie] = http.StatusNotFound
	res[ErrSessionNotExist] = http.StatusNotFound

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
	res[ErrImageNotFound] = http.StatusNotFound
	res[ErrImage] = http.StatusInternalServerError

	// User delivery
	res[ErrGetUserRequest] = http.StatusInternalServerError
	res[ErrWrongValidPassword] = http.StatusForbidden

	// Middleware
	res[ErrBigRequest] = http.StatusBadRequest
	res[ErrConvertLength] = http.StatusBadRequest

	// Security
	res[ErrCsrfTokenCreate] = http.StatusInternalServerError
	res[ErrCsrfTokenCheck] = http.StatusForbidden
	res[ErrCsrfTokenCheckInternal] = http.StatusInternalServerError
	res[ErrCsrfTokenExpired] = http.StatusForbidden
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
