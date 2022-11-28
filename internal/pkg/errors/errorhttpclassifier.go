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
	ErrNotFoundInDB             = stdErrors.New("not found")
	ErrWorkDatabase             = stdErrors.New("error sql")
	ErrGetParamsConvert         = stdErrors.New("err get sql params")
	ErrUnsupportedSortParameter = stdErrors.New("unsupported sort parameter")

	// Collection service
	ErrNotFindSuchTarget     = stdErrors.New("not found such target")
	ErrCollectionIsNotPublic = stdErrors.New("this collection is not public")

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
	ErrBadUserCollectionID      = stdErrors.New("this collection doesn't belong to current user")
	ErrFilmExistInCollection    = stdErrors.New("such film exist in collection")
	ErrFilmNotExistInCollection = stdErrors.New("such film not found in collection")
	ErrFilmRatingNotExist       = stdErrors.New("film rating not exist")

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
	ErrGenreNotFound       = stdErrors.New("genre not found")
	ErrTagNotFound         = stdErrors.New("tag not found")
	ErrFilmNotFound        = stdErrors.New("film not found")
	ErrPersonNotFound      = stdErrors.New("person not found")
	ErrImageNotFound       = stdErrors.New("image not found")
	ErrSessionNotFound     = stdErrors.New("session not found")
	ErrUserNotFound        = stdErrors.New("user not found")
	ErrFilmsNotFound       = stdErrors.New("films not found")
	ErrCollectionsNotFound = stdErrors.New("collections not found")
)

type ErrHTTPClassifier struct {
	table map[string]int
}

func NewErrHTTPClassifier() ErrHTTPClassifier {
	res := make(map[string]int)

	// Common delivery
	res[ErrBadBodyRequest.Error()] = http.StatusBadRequest
	res[ErrJSONUnexpectedEnd.Error()] = http.StatusBadRequest
	res[ErrContentTypeUndefined.Error()] = http.StatusBadRequest
	res[ErrUnsupportedMediaType.Error()] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody.Error()] = http.StatusBadRequest
	res[ErrConvertQueryType.Error()] = http.StatusBadRequest
	res[ErrQueryRequiredEmpty.Error()] = http.StatusBadRequest
	res[ErrBadRequestParams.Error()] = http.StatusBadRequest
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = http.StatusBadRequest
	res[ErrBadRequestParams.Error()] = http.StatusBadRequest

	// Common repository
	res[ErrNotFoundInDB.Error()] = http.StatusNotFound
	res[ErrWorkDatabase.Error()] = http.StatusInternalServerError
	res[ErrGetParamsConvert.Error()] = http.StatusInternalServerError
	res[ErrUnsupportedSortParameter.Error()] = http.StatusBadRequest

	// Collection service
	res[ErrNotFindSuchTarget.Error()] = http.StatusNotFound
	res[ErrCollectionIsNotPublic.Error()] = http.StatusForbidden

	// Auth delivery
	res[ErrNoCookie.Error()] = http.StatusNotFound

	// Auth repository
	res[ErrUserExist.Error()] = http.StatusBadRequest
	res[ErrUserNotExist.Error()] = http.StatusNotFound

	// Auth service
	res[ErrInvalidNickname.Error()] = http.StatusBadRequest
	res[ErrInvalidEmail.Error()] = http.StatusBadRequest
	res[ErrInvalidPassword.Error()] = http.StatusBadRequest
	res[ErrIncorrectPassword.Error()] = http.StatusForbidden

	// Image delivery
	res[ErrBigImage.Error()] = http.StatusBadRequest
	res[ErrBadImageType.Error()] = http.StatusBadRequest

	// Image repository
	res[ErrImage.Error()] = http.StatusInternalServerError

	// User delivery
	res[ErrGetUserRequest.Error()] = http.StatusInternalServerError
	res[ErrWrongValidPassword.Error()] = http.StatusForbidden

	// User service
	res[ErrBadUserCollectionID.Error()] = http.StatusForbidden
	res[ErrFilmExistInCollection.Error()] = http.StatusBadRequest
	res[ErrFilmNotExistInCollection.Error()] = http.StatusNotFound
	res[ErrFilmRatingNotExist.Error()] = http.StatusNotFound

	// Middleware
	res[ErrBigRequest.Error()] = http.StatusBadRequest
	res[ErrConvertLength.Error()] = http.StatusBadRequest

	// Security
	res[ErrCsrfTokenCreate.Error()] = http.StatusInternalServerError
	res[ErrCsrfTokenCheck.Error()] = http.StatusForbidden
	res[ErrCsrfTokenCheckInternal.Error()] = http.StatusInternalServerError
	res[ErrCsrfTokenExpired.Error()] = http.StatusForbidden
	res[ErrCsrfTokenInvalid.Error()] = http.StatusForbidden

	// Not found
	res[ErrGenreNotFound.Error()] = http.StatusNotFound
	res[ErrTagNotFound.Error()] = http.StatusNotFound
	res[ErrPersonNotFound.Error()] = http.StatusNotFound
	res[ErrFilmNotFound.Error()] = http.StatusNotFound
	res[ErrSessionNotFound.Error()] = http.StatusNotFound
	res[ErrUserNotFound.Error()] = http.StatusNotFound
	res[ErrImageNotFound.Error()] = http.StatusNotFound
	res[ErrFilmsNotFound.Error()] = http.StatusNotFound
	res[ErrCollectionsNotFound.Error()] = http.StatusNotFound

	return ErrHTTPClassifier{
		table: res,
	}
}

var errHTTPCsf = NewErrHTTPClassifier()

func GetErrorCodeHTTP(err error) (int, bool) {
	code, exist := errHTTPCsf.table[err.Error()]
	if !exist {
		return http.StatusInternalServerError, exist
	}

	return code, exist
}
