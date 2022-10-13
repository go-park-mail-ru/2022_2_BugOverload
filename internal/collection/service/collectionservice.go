package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// CollectionService provides universal service for work with collection.
type CollectionService interface {
	GetPopular(ctx context.Context) ([]models.Film, error)
	GetInCinema(ctx context.Context) ([]models.Film, error)
}

// collectionService is implementation for collection service corresponding to the CollectionService interface.
type collectionService struct {
	collectionRepo repository.CollectionRepository
	contextTimeout time.Duration
}

// NewCollectionService is constructor for collectionService.
// Accepts CollectionRepository interfaces and context timeout.
func NewCollectionService(cr repository.CollectionRepository, timeout time.Duration) CollectionService {
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
