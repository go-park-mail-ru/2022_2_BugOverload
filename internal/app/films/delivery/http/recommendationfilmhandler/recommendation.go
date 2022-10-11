package recommendationfilmhandler

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"net/http"

	stdErrors "github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	filmModels "go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/models"
	collectionInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/contextparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

type handler struct {
	filmService collectionInterface.FilmsService
	authService authInterface.AuthService
}

func NewHandler(fs collectionInterface.FilmsService, as authInterface.AuthService) interfaces.Handler {
	return &handler{
		fs,
		as,
	}
}

func (h *handler) Action(w http.ResponseWriter, r *http.Request) {
	cookieStr := r.Header.Get("Cookie")

	var user models.User
	var err error

	if cookieStr != "" {
		ctx := context.WithValue(r.Context(), contextparams.CookieKey, cookieStr)

		user, err = h.authService.GetUserBySession(ctx)
		if err != nil {
			httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
			return
		}
	}

	filmRecommendation, err := h.filmService.GerRecommendation(r.Context(), &user)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(stdErrors.Cause(err)))
		return
	}

	response := filmModels.NewRecommendFilmRequest(filmRecommendation)

	httpwrapper.Response(w, http.StatusOK, response.ToPublic())
}
