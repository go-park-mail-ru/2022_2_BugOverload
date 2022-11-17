package errors

import (
	stdErrors "github.com/pkg/errors"
)

type ErrLogClassifier struct {
	table map[error]string
}

const (
	infoLogLevel  = "info"
	errLogLevel   = "error"
	debugLogLevel = "debug"
)

func NewErrLogClassifier() ErrLogClassifier {
	res := make(map[error]string)

	// Common delivery
	res[ErrBadBodyRequest] = infoLogLevel
	res[ErrJSONUnexpectedEnd] = infoLogLevel
	res[ErrContentTypeUndefined] = infoLogLevel
	res[ErrUnsupportedMediaType] = infoLogLevel
	res[ErrEmptyBody] = infoLogLevel
	res[ErrConvertQueryType] = infoLogLevel
	res[ErrQueryRequiredEmpty] = infoLogLevel
	res[ErrBadQueryParams] = infoLogLevel
	res[ErrEmptyField] = infoLogLevel
	res[ErrEmptyRequiredFields] = infoLogLevel
	res[ErrBadRequestParams] = infoLogLevel

	// Common repository
	res[ErrNotFoundInDB] = errLogLevel
	res[ErrWorkDatabase] = errLogLevel
	res[ErrGetParamsConvert] = errLogLevel

	// Auth delivery
	res[ErrNoCookie] = infoLogLevel
	res[ErrSessionNotExist] = infoLogLevel

	// Auth repository
	res[ErrUserExist] = errLogLevel
	res[ErrUserNotExist] = errLogLevel

	// Auth service
	res[ErrInvalidNickname] = infoLogLevel
	res[ErrInvalidEmail] = infoLogLevel
	res[ErrInvalidPassword] = infoLogLevel
	res[ErrIncorrectPassword] = infoLogLevel

	// Image delivery
	res[ErrBigImage] = infoLogLevel
	res[ErrBadImageType] = infoLogLevel

	// Image repository
	res[ErrImageNotFound] = errLogLevel
	res[ErrImage] = errLogLevel

	// User delivery
	res[ErrGetUserRequest] = infoLogLevel
	res[ErrWrongValidPassword] = infoLogLevel

	// User service
	res[ErrFilmRatingNotExist] = errLogLevel

	// Middleware
	res[ErrBigRequest] = infoLogLevel
	res[ErrConvertLength] = infoLogLevel

	// Security
	res[ErrCsrfTokenCreate] = errLogLevel
	res[ErrCsrfTokenCheck] = errLogLevel
	res[ErrCsrfTokenCheckInternal] = errLogLevel
	res[ErrCsrfTokenExpired] = errLogLevel
	res[ErrCsrfTokenInvalid] = errLogLevel

	return ErrLogClassifier{
		table: res,
	}
}

func GetLogLevelErr(err error) (string, error) {
	level, exist := errLogCsf.table[err]
	if !exist {
		return "", stdErrors.New("error not found")
	}

	return level, nil
}

var errLogCsf = NewErrLogClassifier()
