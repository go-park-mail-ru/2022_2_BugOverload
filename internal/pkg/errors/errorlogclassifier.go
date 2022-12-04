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
	res[ErrBadBodyRequest.Error()] = errLogLevel
	res[ErrJSONUnexpectedEnd.Error()] = errLogLevel
	res[ErrContentTypeUndefined.Error()] = errLogLevel
	res[ErrUnsupportedMediaType.Error()] = errLogLevel
	res[ErrEmptyBody.Error()] = errLogLevel
	res[ErrConvertQueryType.Error()] = errLogLevel
	res[ErrQueryRequiredEmpty.Error()] = errLogLevel
	res[ErrBadRequestParams.Error()] = errLogLevel
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = errLogLevel
	res[ErrBadRequestParams.Error()] = errLogLevel

	// Common repository
	res[ErrNotFoundInDB.Error()] = errLogLevel
	res[ErrWorkDatabase.Error()] = errLogLevel
	res[ErrGetParamsConvert.Error()] = errLogLevel

	// Collection service
	res[ErrCollectionIsNotPublic.Error()] = errLogLevel

	// Auth delivery
	res[ErrNoCookie.Error()] = errLogLevel

	// Auth repository
	res[ErrUserExist.Error()] = errLogLevel
	res[ErrUserNotExist.Error()] = errLogLevel
	res[ErrCreateSession.Error()] = errLogLevel

	// Auth service
	res[ErrInvalidNickname.Error()] = errLogLevel
	res[ErrInvalidEmail.Error()] = errLogLevel
	res[ErrInvalidPassword.Error()] = errLogLevel
	res[ErrIncorrectPassword.Error()] = errLogLevel

	// Image delivery
	res[ErrBigImage.Error()] = errLogLevel
	res[ErrBadImageType.Error()] = errLogLevel

	// Image repository
	res[ErrImage.Error()] = errLogLevel

	// User delivery
	res[ErrGetUserRequest.Error()] = errLogLevel
	res[ErrWrongValidPassword.Error()] = errLogLevel

	// User service
	res[ErrFilmExistInCollection.Error()] = errLogLevel
	res[ErrFilmNotExistInCollection.Error()] = errLogLevel
	res[ErrBadUserCollectionID.Error()] = errLogLevel
	res[ErrFilmRatingNotExist.Error()] = errLogLevel

	// Middleware
	res[ErrBigRequest.Error()] = errLogLevel
	res[ErrConvertLength.Error()] = errLogLevel

	// Security
	res[ErrCsrfTokenCreate.Error()] = errLogLevel
	res[ErrCsrfTokenCheck.Error()] = errLogLevel
	res[ErrCsrfTokenCheckInternal.Error()] = errLogLevel
	res[ErrCsrfTokenExpired.Error()] = errLogLevel
	res[ErrCsrfTokenInvalid.Error()] = errLogLevel

	// Not found
	res[ErrGenreNotFound.Error()] = errLogLevel
	res[ErrTagNotFound.Error()] = errLogLevel
	res[ErrPersonNotFound.Error()] = errLogLevel
	res[ErrFilmNotFound.Error()] = errLogLevel
	res[ErrSessionNotFound.Error()] = errLogLevel
	res[ErrUserNotFound.Error()] = errLogLevel
	res[ErrImageNotFound.Error()] = errLogLevel
	res[ErrFilmsNotFound.Error()] = errLogLevel
	res[ErrCollectionsNotFound.Error()] = errLogLevel
	res[ErrCollectionNotFound.Error()] = errLogLevel

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
