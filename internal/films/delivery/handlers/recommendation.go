package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	stdErrors "github.com/pkg/errors"

	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/films/delivery/models"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/films/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

// recommendationFilmHandler is the structure that handles the request for
// getting recommendation film for authorized ot unauthorized user.
type recommendationFilmHandler struct {
	filmService serviceFilms.FilmsService
	authService serviceAuth.AuthService
}

// NewRecommendationFilmHandler is constructor for recommendationFilmHandler in this pkg - recommendation film.
func NewRecommendationFilmHandler(fs serviceFilms.FilmsService, as serviceAuth.AuthService) pkg.Handler {
	return &recommendationFilmHandler{
		fs,
		as,
	}
}

// TODO: возможно нужно раздить на рекоммендацию авторизованного и неавторизованного пользователя через middleware

// Action is a method for initial validation of the request and data and
// delivery of the data to the service at the business logic level.
// @Summary Recommendation film
// @Description Getting a recommended movie for the user
// @tags films
// @Produce json
// @Success 200 {object} models.RecommendFilmResponse "returns recommendation film movies for auth user or not auth"
// @Failure 400 "return error"
// @Failure 405 "method not allowed"
// @Failure 500 "something unusual has happened"
// @Router /v1/recommendation_film [GET]
func (h *recommendationFilmHandler) Action(w http.ResponseWriter, r *http.Request) {
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

	response := models.NewRecommendFilmResponse()

	httpwrapper.Response(w, http.StatusOK, response.ToPublic(&filmRecommendation))
}