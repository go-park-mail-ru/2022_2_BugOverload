package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// FilmsService provides universal service for work with film.
type FilmsService interface {
	GetRecommendation(ctx context.Context) (models.Film, error)
	GetFilmByID(ctx context.Context, film *models.Film, params *pkg.GetFilmParams) (models.Film, error)
}

// filmService is implementation for auth service corresponding to the FilmsService interface.
type filmService struct {
	filmsRepo repository.FilmRepository
}

// NewFilmService is constructor for filmService. Accepts FilmsRepository interfaces.
func NewFilmService(cr repository.FilmRepository) FilmsService {
	return &filmService{
		filmsRepo: cr,
	}
}

func (c *filmService) GetRecommendation(ctx context.Context) (models.Film, error) {
	film, err := c.filmsRepo.GetRecommendation(ctx)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GetRecommendation")
	}

	return film, nil
}

func (c *filmService) GetFilmByID(ctx context.Context, film *models.Film, params *pkg.GetFilmParams) (models.Film, error) {
	filmRepo, err := c.filmsRepo.GetFilmByID(ctx, film, params)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GetFilmByID")
	}

	return filmRepo, nil
}
