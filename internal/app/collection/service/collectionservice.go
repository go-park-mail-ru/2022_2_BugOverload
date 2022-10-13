package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// collectionService is implementation for collection service corresponding to the CollectionService interface.
type collectionService struct {
	collectionRepo interfaces.CollectionRepository
	contextTimeout time.Duration
}

// NewCollectionService is constructor for collectionService.
// Accepts CollectionRepository interfaces and context timeout.
func NewCollectionService(cr interfaces.CollectionRepository, timeout time.Duration) interfaces.CollectionService {
	return &collectionService{
		collectionRepo: cr,
		contextTimeout: timeout,
	}
}

// GetPopular is the service that accesses the interface CollectionRepository
func (c *collectionService) GetPopular(ctx context.Context) ([]models.Film, error) {
	inCinemaCollection, err := c.collectionRepo.GetPopular(ctx)
	if err != nil {
		return []models.Film{}, stdErrors.Wrap(err, "GetPopular")
	}

	return inCinemaCollection, nil
}

// GetInCinema is the service that accesses the interface CollectionRepository
func (c *collectionService) GetInCinema(ctx context.Context) ([]models.Film, error) {
	popularCollection, err := c.collectionRepo.GetInCinema(ctx)
	if err != nil {
		return []models.Film{}, stdErrors.Wrap(err, "GetInCinema")
	}

	return popularCollection, nil
}
