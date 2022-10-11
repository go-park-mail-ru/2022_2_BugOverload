package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type filmService struct {
	filmsRepo      interfaces.FilmsRepository
	contextTimeout time.Duration
}

func NewFilmService(cr interfaces.FilmsRepository, timeout time.Duration) interfaces.FilmsService {
	return &filmService{
		filmsRepo:      cr,
		contextTimeout: timeout,
	}
}

func (c filmService) GerRecommendation(ctx context.Context, user *models.User) (models.Film, error) {
	inCinemaCollection, err := c.filmsRepo.GerRecommendation(ctx, user)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GerRecommendation")
	}

	return inCinemaCollection, nil
}
