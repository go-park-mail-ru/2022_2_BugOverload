package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// FilmsService provides universal service for work with films.
type FilmsService interface {
	GerRecommendation(ctx context.Context) (models.Film, error)
}

// FilmsRepository provides the versatility of films repositories.
type FilmsRepository interface {
	GerRecommendation(ctx context.Context) (models.Film, error)
}
