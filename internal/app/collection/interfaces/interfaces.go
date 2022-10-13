package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// CollectionService provides universal service for work with collection.
type CollectionService interface {
	GetPopular(ctx context.Context) ([]models.Film, error)
	GetInCinema(ctx context.Context) ([]models.Film, error)
}

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetPopular(ctx context.Context) ([]models.Film, error)
	GetInCinema(ctx context.Context) ([]models.Film, error)
}
