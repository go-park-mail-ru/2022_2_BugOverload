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

	// Auth
	res[ErrEmptyFieldAuth] = infoLogLevel
	res[ErrUserExist] = infoLogLevel
	res[ErrUserNotExist] = infoLogLevel
	res[ErrSignupUserExist] = infoLogLevel
	res[ErrBadBodyRequest] = infoLogLevel

	res[ErrLoginCombinationNotFound] = infoLogLevel
	res[ErrNoCookie] = infoLogLevel
	res[ErrCookieNotExist] = infoLogLevel
	res[ErrSessionNotExist] = infoLogLevel
	res[ErrQueryRequiredEmpty] = infoLogLevel
	res[ErrWrongPassword] = infoLogLevel

	// Access
	res[ErrNoAccess] = infoLogLevel

	// Auth Validation
	res[ErrInvalidEmail] = infoLogLevel
	res[ErrInvalidPassword] = infoLogLevel

	// Films
	res[ErrFilmNotFound] = infoLogLevel

	// Images
	res[ErrImageNotFound] = infoLogLevel

	res[ErrGetImageStorage] = infoLogLevel
	res[ErrReadImage] = infoLogLevel

	res[ErrImage] = infoLogLevel
	res[ErrBadImageType] = infoLogLevel

	// Def Validation
	res[ErrJSONUnexpectedEnd] = infoLogLevel
	res[ErrContentTypeUndefined] = infoLogLevel
	res[ErrUnsupportedMediaType] = infoLogLevel
	res[ErrEmptyBody] = infoLogLevel
	res[ErrBigRequest] = infoLogLevel
	res[ErrConvertLength] = infoLogLevel
	res[ErrConvertQuery] = infoLogLevel
	res[ErrQueryBad] = infoLogLevel
	res[ErrEmptyRequiredFields] = infoLogLevel
	res[ErrBadRequestParams] = infoLogLevel

	// DB
	res[ErrPostgresRequest] = errLogLevel
	res[ErrNotFoundInDB] = infoLogLevel
	res[ErrGetParamsConvert] = infoLogLevel

	// Security
	res[ErrCsrfTokenCreate] = infoLogLevel
	res[ErrCsrfTokenCheck] = infoLogLevel
	res[ErrCsrfTokenCheckInternal] = infoLogLevel
	res[ErrCsrfTokenExpired] = infoLogLevel
	res[ErrCsrfTokenNotFound] = infoLogLevel
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
