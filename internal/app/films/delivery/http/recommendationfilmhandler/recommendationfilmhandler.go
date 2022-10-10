package recommendationfilmhandler

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/httpwrapper"
)

// FilmRecommendationHandler is structure for API films requests processing
type FilmRecommendationHandler struct {
	storage *memory.FilmStorage
}

// NewHandlerRecommendationFilm is constructor for NewHandlerRecommendationFilm
func NewHandlerRecommendationFilm(fs *memory.FilmStorage) interfaces.Handler {
	return &FilmRecommendationHandler{
		fs,
	}
}

// tmp const
const countFilmPreview = 4

// Action is handle film to poster request
func (hf *FilmRecommendationHandler) Action(w http.ResponseWriter, r *http.Request) {
	var recommendFilmRequest models.RecommendFilmRequest

	max := hf.storage.GetStorageLen()
	min := max - countFilmPreview

	if max == 0 {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))
		return
	}

	film, err := hf.storage.GetFilm(uint(utils.Rand(max-min) + min))
	if err != nil {
		httpwrapper.DefaultHandlerError(w, errors.NewErrFilms(errors.ErrFilmNotFound))
		return
	}

	recommendFilmRequest.SetFilm(film)

	response := recommendFilmRequest.CreateResponse()

	httpwrapper.Response(w, http.StatusOK, response)
}
