package errors

import (
	stdErrors "github.com/pkg/errors"
)

type ErrLogClassifier struct {
	table map[string]string
}

const (
	infoLogLevel  = "info"
	errLogLevel   = "error"
	debugLogLevel = "api"
)

func NewErrLogClassifier() ErrLogClassifier {
	res := make(map[string]string)

	// Common delivery
	res[ErrBadBodyRequest.Error()] = infoLogLevel
	res[ErrJSONUnexpectedEnd.Error()] = infoLogLevel
	res[ErrContentTypeUndefined.Error()] = infoLogLevel
	res[ErrUnsupportedMediaType.Error()] = infoLogLevel
	res[ErrEmptyBody.Error()] = infoLogLevel
	res[ErrConvertQueryType.Error()] = infoLogLevel
	res[ErrQueryRequiredEmpty.Error()] = infoLogLevel
	res[ErrBadRequestParams.Error()] = infoLogLevel
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = infoLogLevel
	res[ErrBadRequestParams.Error()] = infoLogLevel

	// Common repository
	res[ErrNotFoundInDB.Error()] = errLogLevel
	res[ErrWorkDatabase.Error()] = errLogLevel
	res[ErrGetParamsConvert.Error()] = errLogLevel

	// Collection service
	res[ErrFilmExistInCollection.Error()] = errLogLevel
	res[ErrFilmNotExistInCollection.Error()] = errLogLevel

	// Auth delivery
	res[ErrNoCookie.Error()] = infoLogLevel

	// Auth repository
	res[ErrUserExist.Error()] = errLogLevel
	res[ErrUserNotExist.Error()] = errLogLevel

	// Auth service
	res[ErrInvalidNickname.Error()] = infoLogLevel
	res[ErrInvalidEmail.Error()] = infoLogLevel
	res[ErrInvalidPassword.Error()] = infoLogLevel
	res[ErrIncorrectPassword.Error()] = infoLogLevel

	// Image delivery
	res[ErrBigImage.Error()] = infoLogLevel
	res[ErrBadImageType.Error()] = infoLogLevel

	// Image repository
	res[ErrImage.Error()] = errLogLevel

	// User delivery
	res[ErrGetUserRequest.Error()] = infoLogLevel
	res[ErrWrongValidPassword.Error()] = infoLogLevel

	// User service
	res[ErrFilmRatingNotExist.Error()] = errLogLevel

	// Middleware
	res[ErrBigRequest.Error()] = infoLogLevel
	res[ErrConvertLength.Error()] = infoLogLevel

	// Security
	res[ErrCsrfTokenCreate.Error()] = errLogLevel
	res[ErrCsrfTokenCheck.Error()] = errLogLevel
	res[ErrCsrfTokenCheckInternal.Error()] = errLogLevel
	res[ErrCsrfTokenExpired.Error()] = errLogLevel
	res[ErrCsrfTokenInvalid.Error()] = errLogLevel

	// Not found
	res[ErrGenreNotFound.Error()] = infoLogLevel
	res[ErrTagNotFound.Error()] = infoLogLevel
	res[ErrPersonNotFound.Error()] = infoLogLevel
	res[ErrFilmNotFound.Error()] = infoLogLevel
	res[ErrSessionNotFound.Error()] = infoLogLevel
	res[ErrUserNotFound.Error()] = infoLogLevel
	res[ErrImageNotFound.Error()] = infoLogLevel
	res[ErrFilmsNotFound.Error()] = infoLogLevel
	res[ErrCollectionsNotFound.Error()] = infoLogLevel

	return ErrLogClassifier{
		table: res,
	}
}

func GetLogLevelErr(err error) (string, error) {
	level, exist := errLogCsf.table[err.Error()]
	if !exist {
		return "", stdErrors.New("error not found")
	}

	return level, nil
}

var errLogCsf = NewErrLogClassifier()
