package recommendationfilmhandler

import (
	"net/http"

	stdErrors "github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	filmModels "go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/models"
	collectionInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// handler is the structure that handles the request for
// getting recommendation film for authorized ot unauthorized user.
type handler struct {
	filmService collectionInterface.FilmsService
	authService authInterface.AuthService
}

// NewHandler is constructor for handler in this pkg - recommendation film.
func NewHandler(fs collectionInterface.FilmsService, as authInterface.AuthService) interfaces.Handler {
	return &handler{
		fs,
		as,
	}
}

// TODO: возможно нужно раздить на рекоммендацию авторизованного и неавторизованного пользователя через middleware

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	//  cookieStr := r.Header.Get("Cookie")
	//
	//  var user models.User
	//  var err error
	//
	//  if cookieStr != "" {
	//	ctx := context.WithValue(r.Context(), params.CookieKey, cookieStr)
	//
	//	user, err = h.authService.GetUserBySession(ctx)
	//	if err != nil {
	//		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
	//		return
	//	}
	//  }

	filmRecommendation, err := h.filmService.GerRecommendation(r.Context())
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	response := filmModels.NewRecommendFilmRequest(filmRecommendation)

	httpwrapper.Response(w, http.StatusOK, response.ToPublic())
}
