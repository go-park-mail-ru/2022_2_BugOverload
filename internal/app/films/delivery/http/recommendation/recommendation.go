package recommendation

import (
	"math/rand"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/OLD/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/models"
)

// FilmRecommendationHandler is structure for API films requests processing
type FilmRecommendationHandler struct {
	storage *database.FilmStorage
}

// NewHandlerRecommendationFilm is constructor for NewHandlerRecommendationFilm
func NewHandlerRecommendationFilm(fs *database.FilmStorage) *FilmRecommendationHandler {
	return &FilmRecommendationHandler{
		fs,
	}
}

// GetRecommendedFilm is handle film to poster request
func (hf *FilmRecommendationHandler) GetRecommendedFilm(w http.ResponseWriter, r *http.Request) {
	var recommendFilmRequest models.RecommendFilmRequest

	max := hf.storage.GetStorageLen()
	min := max - 4

	if max == 0 {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))
		return
	}

	film, err := hf.storage.GetFilm(uint(rand.Intn(max-min) + min))
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))
		return
	}

	recommendFilmRequest.SetFilm(film)

	response := recommendFilmRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
