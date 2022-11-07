package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/review/repository"
)

// ReviewService provides universal service for work with reviews.
type ReviewService interface {
	GetReviewsByFilmID(ctx context.Context, params *pkg.GetReviewsFilmParams) ([]models.Review, error)
}

// reviewService is implementation for users service corresponding to the ReviewService interface.
type reviewService struct {
	reviewRepo repository.ReviewRepository
}

// NewReviewService is constructor for reviewService.
func NewReviewService(pr repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepo: pr,
	}
}

func (r *reviewService) GetReviewsByFilmID(ctx context.Context, params *pkg.GetReviewsFilmParams) ([]models.Review, error) {
	reviews, err := r.reviewRepo.GetReviewsByFilmID(ctx, params)
	if err != nil {
		return []models.Review{}, stdErrors.Wrap(err, "GetReviewsByFilmID")
	}

	return reviews, nil
}
