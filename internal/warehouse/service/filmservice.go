package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
)

//go:generate mockgen -source filmservice.go -destination mocks/mockfilmservice.go -package mockWarehouseService

// FilmService provides universal service for work with film.
type FilmService interface {
	GetRecommendation(ctx context.Context) (models.Film, error)
	GetFilmByID(ctx context.Context, film *models.Film, params *constparams.GetFilmParams) (models.Film, error)
	GetReviewsByFilmID(ctx context.Context, params *constparams.GetFilmReviewsParams) ([]models.Review, error)
}

// filmService is implementation for auth service corresponding to the FilmService interface.
type filmService struct {
	filmsRepo film.Repository
}

// NewFilmService is constructor for filmService. Accepts FilmsRepository interfaces.
func NewFilmService(cr film.Repository) FilmService {
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

func (c *filmService) GetFilmByID(ctx context.Context, film *models.Film, params *constparams.GetFilmParams) (models.Film, error) {
	filmRepo, err := c.filmsRepo.GetFilmByID(ctx, film, params)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "GetFilmByID")
	}

	return filmRepo, nil
}

func (c *filmService) GetReviewsByFilmID(ctx context.Context, params *constparams.GetFilmReviewsParams) ([]models.Review, error) {
	reviews, err := c.filmsRepo.GetReviewsByFilmID(ctx, params)
	if err != nil {
		return []models.Review{}, stdErrors.Wrap(err, "GetReviewsByFilmID")
	}

	return reviews, nil
}
