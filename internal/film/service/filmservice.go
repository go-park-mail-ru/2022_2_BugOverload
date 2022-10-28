package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// FilmsService provides universal service for work with film.
type FilmsService interface {
	GerRecommendation(ctx context.Context) (models.Film, error)
}

// filmService is implementation for auth service corresponding to the FilmsService interface.
type filmService struct {
	filmsRepo repository.FilmsRepository
}

// NewFilmService is constructor for filmService. Accepts FilmsRepository interfaces.
func NewFilmService(cr repository.FilmsRepository) FilmsService {
	return &filmService{
		filmsRepo: cr,
	}
}

func (c *filmService) GerRecommendation(ctx context.Context) (models.Film, error) {
	inCinemaCollection, err := c.filmsRepo.GerRecommendation(ctx)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GerRecommendation")
	}

	return inCinemaCollection, nil
}
