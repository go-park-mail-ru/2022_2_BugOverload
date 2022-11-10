package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"

	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// CollectionService provides universal service for work with collection.
type CollectionService interface {
	GetCollectionByTag(ctx context.Context, params *pkg.GetCollectionTagParams) (models.Collection, error)
}

// collectionService is implementation for collection service corresponding to the CollectionService interface.
type collectionService struct {
	collectionRepo repository.CollectionRepository
}

// NewCollectionService is constructor for collectionService.
// Accepts CollectionRepository interfaces.
func NewCollectionService(cr repository.CollectionRepository) CollectionService {
	return &collectionService{
		collectionRepo: cr,
	}
}

// GetCollectionByTag is the service that accesses the interface CollectionRepository
func (c *collectionService) GetCollectionByTag(ctx context.Context, params *pkg.GetCollectionTagParams) (models.Collection, error) {
	inCinemaCollection, err := c.collectionRepo.GetCollectionByTag(ctx, params)
	if err != nil {
		return models.Collection{}, errors.Wrap(err, "GetCollectionByTag")
	}

	return inCinemaCollection, nil
}
