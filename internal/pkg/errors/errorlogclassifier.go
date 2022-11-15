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
	res[ErrConvertQuery] = infoLogLevel
	res[ErrQueryRequiredEmpty] = infoLogLevel
	res[ErrQueryBad] = infoLogLevel
	res[ErrEmptyField] = infoLogLevel
	res[ErrEmptyRequiredFields] = infoLogLevel
	res[ErrBadRequestParams] = infoLogLevel

	// Common repository
	res[ErrNotFoundInDB] = infoLogLevel
	res[ErrPostgresRequest] = infoLogLevel
	res[ErrGetParamsConvert] = infoLogLevel

	// Auth delivery
	res[ErrNoCookie] = infoLogLevel
	res[ErrSessionNotExist] = infoLogLevel

	// Auth repository
	res[ErrSignupUserExist] = infoLogLevel
	res[ErrUserNotExist] = infoLogLevel

	// Auth service
	res[ErrInvalidNickname] = infoLogLevel
	res[ErrInvalidEmail] = infoLogLevel
	res[ErrInvalidPassword] = infoLogLevel
	res[ErrIncorrectPassword] = infoLogLevel

	// Image delivery
	res[ErrBigImage] = infoLogLevel
	res[ErrBadImageType] = infoLogLevel

	// Image repository
	res[ErrImageNotFound] = infoLogLevel
	res[ErrImage] = infoLogLevel

	// User delivery
	res[ErrGetUserRequest] = infoLogLevel
	res[ErrWrongPassword] = infoLogLevel

	// Middleware
	res[ErrBigRequest] = infoLogLevel
	res[ErrConvertLength] = infoLogLevel

	// Security
	res[ErrCsrfTokenCreate] = infoLogLevel
	res[ErrCsrfTokenCheck] = infoLogLevel
	res[ErrCsrfTokenCheckInternal] = infoLogLevel
	res[ErrCsrfTokenExpired] = infoLogLevel
	res[ErrCsrfTokenInvalid] = infoLogLevel

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
