package httpwrapper

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// DefaultHandlerError is error handler that detects the type of error and gives an error response.
func DefaultHandlerError(w http.ResponseWriter, err error) {
	errResp := ErrResponse{
		err.Error(),
	}

	var errHTTP errors.DefaultValidationError
	if ok := stdErrors.As(err, &errHTTP); ok {
		Response(w, errHTTP.Code, errResp)
		return
	}

	var errAuth errors.AuthError
	if ok := stdErrors.As(err, &errAuth); ok {
		Response(w, errAuth.Code, errResp)
		return
	}

	var errFilms errors.FilmError
	if ok := stdErrors.As(err, &errFilms); ok {
		Response(w, errFilms.Code, errResp)
		return
	}

	var errImages errors.ImagesError
	if ok := stdErrors.As(err, &errImages); ok {
		Response(w, errImages.Code, errResp)
		return
	}

	var profileError errors.ProfileError
	if ok := stdErrors.As(err, &profileError); ok {
		Response(w, profileError.Code, errResp)
		return
	}

	var errCollection errors.CollectionError
	if ok := stdErrors.As(err, &errCollection); ok {
		Response(w, errCollection.Code, errResp)
		return
	}

	var errProfile errors.ProfileError
	if ok := stdErrors.As(err, &errProfile); ok {
		Response(w, errProfile.Code, errResp)
		return
	}

	var errReview errors.ReviewError
	if ok := stdErrors.As(err, &errReview); ok {
		Response(w, errReview.Code, errResp)
		return
	}

	Response(w, http.StatusInternalServerError, errResp)
}
