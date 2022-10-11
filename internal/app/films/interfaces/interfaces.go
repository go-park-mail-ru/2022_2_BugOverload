package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type FilmsService interface {
	GerRecommendation(ctx context.Context, user *models.User) (models.Film, error)
}

type FilmsRepository interface {
	GerRecommendation(ctx context.Context, user *models.User) (models.Film, error)
}
