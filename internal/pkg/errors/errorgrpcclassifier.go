package errors

import (
	"google.golang.org/grpc/codes"
)

type ErrGRPCClassifier struct {
	table map[string]codes.Code
}

func NewErrGRPCClassifier() ErrGRPCClassifier {
	res := make(map[string]codes.Code)

	// Common delivery
	res[ErrBadBodyRequest.Error()] = codes.InvalidArgument
	res[ErrJSONUnexpectedEnd.Error()] = codes.InvalidArgument
	res[ErrContentTypeUndefined.Error()] = codes.InvalidArgument
	res[ErrUnsupportedMediaType.Error()] = codes.InvalidArgument
	res[ErrEmptyBody.Error()] = codes.InvalidArgument
	res[ErrConvertQueryType.Error()] = codes.InvalidArgument
	res[ErrQueryRequiredEmpty.Error()] = codes.InvalidArgument
	res[ErrBadRequestParams.Error()] = codes.InvalidArgument
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = codes.InvalidArgument
	res[ErrBadRequestParams.Error()] = codes.InvalidArgument
	res[ErrGetEasyJSON.Error()] = codes.Internal

	// Common repository
	res[ErrNotFoundInDB.Error()] = codes.NotFound
	res[ErrWorkDatabase.Error()] = codes.Internal
	res[ErrGetParamsConvert.Error()] = codes.Internal
	res[ErrUnsupportedSortParameter.Error()] = codes.InvalidArgument

	// Collection service
	res[ErrCollectionIsNotPublic.Error()] = codes.PermissionDenied
	res[ErrNotFindSuchTarget.Error()] = codes.NotFound

	// Auth delivery
	res[ErrNoCookie.Error()] = codes.NotFound

	// Auth repository
	res[ErrUserExist.Error()] = codes.InvalidArgument
	res[ErrUserNotExist.Error()] = codes.NotFound
	res[ErrCreateSession.Error()] = codes.Internal

	// Auth service
	res[ErrInvalidNickname.Error()] = codes.InvalidArgument
	res[ErrInvalidEmail.Error()] = codes.InvalidArgument
	res[ErrInvalidPassword.Error()] = codes.InvalidArgument
	res[ErrIncorrectPassword.Error()] = codes.PermissionDenied

	// Image delivery
	res[ErrBigImage.Error()] = codes.InvalidArgument
	res[ErrBadImageType.Error()] = codes.InvalidArgument

	// Image repository
	res[ErrImage.Error()] = codes.Internal

	// User delivery
	res[ErrGetUserRequest.Error()] = codes.Internal
	res[ErrWrongValidPassword.Error()] = codes.PermissionDenied

	// User service
	res[ErrFilmExistInCollection.Error()] = codes.AlreadyExists
	res[ErrFilmNotExistInCollection.Error()] = codes.NotFound
	res[ErrBadUserCollectionID.Error()] = codes.PermissionDenied
	res[ErrFilmRatingNotExist.Error()] = codes.NotFound

	// Middleware
	res[ErrBigRequest.Error()] = codes.InvalidArgument
	res[ErrConvertLength.Error()] = codes.InvalidArgument

	// Security
	res[ErrCsrfTokenCreate.Error()] = codes.Internal
	res[ErrCsrfTokenCheck.Error()] = codes.PermissionDenied
	res[ErrCsrfTokenCheckInternal.Error()] = codes.Internal
	res[ErrCsrfTokenExpired.Error()] = codes.PermissionDenied
	res[ErrCsrfTokenInvalid.Error()] = codes.PermissionDenied

	// Not found
	res[ErrGenreNotFound.Error()] = codes.NotFound
	res[ErrTagNotFound.Error()] = codes.NotFound
	res[ErrPersonNotFound.Error()] = codes.NotFound
	res[ErrFilmNotFound.Error()] = codes.NotFound
	res[ErrSessionNotFound.Error()] = codes.NotFound
	res[ErrUserNotFound.Error()] = codes.NotFound
	res[ErrImageNotFound.Error()] = codes.NotFound
	res[ErrFilmsNotFound.Error()] = codes.NotFound
	res[ErrCollectionsNotFound.Error()] = codes.NotFound
	res[ErrCollectionNotFound.Error()] = codes.NotFound

	return ErrGRPCClassifier{
		table: res,
	}
}

var errGRPCCsf = NewErrGRPCClassifier()

func GetCodeGRPC(message string) codes.Code {
	code, exist := errGRPCCsf.table[message]
	if !exist {
		return codes.Internal
	}

	return code
}
