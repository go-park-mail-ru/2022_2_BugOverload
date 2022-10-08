package recommendation

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	httpwrapper2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
	"math/rand"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/models"
)

// FilmRecommendationHandler is structure for API films requests processing
type FilmRecommendationHandler struct {
	storage *memory.FilmStorage
}

// NewHandlerRecommendationFilm is constructor for NewHandlerRecommendationFilm
func NewHandlerRecommendationFilm(fs *memory.FilmStorage) *FilmRecommendationHandler {
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
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrFilms(errors2.ErrFilmNotFound))
		return
	}

	film, err := hf.storage.GetFilm(uint(rand.Intn(max-min) + min))
	if err != nil {
		httpwrapper2.DefaultHandlerError(w, errors2.NewErrFilms(errors2.ErrFilmNotFound))
		return
	}

	recommendFilmRequest.SetFilm(film)

	response := recommendFilmRequest.CreateResponse()

	httpwrapper2.Response(w, http.StatusOK, response)
}
