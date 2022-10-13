package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// filmService is implementation for auth service corresponding to the FilmsService interface.
type filmService struct {
	filmsRepo      interfaces.FilmsRepository
	contextTimeout time.Duration
}

// NewFilmService is constructor for filmService. Accepts FilmsRepository interfaces and context timeout.
func NewFilmService(cr interfaces.FilmsRepository, timeout time.Duration) interfaces.FilmsService {
	return &filmService{
		filmsRepo:      cr,
		contextTimeout: timeout,
	}
}

func (c filmService) GerRecommendation(ctx context.Context) (models.Film, error) {
	inCinemaCollection, err := c.filmsRepo.GerRecommendation(ctx)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GerRecommendation")
	}

	return inCinemaCollection, nil
}
