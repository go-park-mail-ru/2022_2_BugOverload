package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate mockgen -source collectionservice.go -destination mocks/mockcollectionservice.go -package mockCollectionService

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
	switch params.Tag {
	case pkg.TagFromPopular:
		params.Tag = pkg.TagInPopular
	case pkg.TagFromInCinema:
		params.Tag = pkg.TagInInCinema
	default:
		return models.Collection{}, stdErrors.Wrap(errors.ErrNotFoundInDB, "GetCollectionByTag")
	}

	inCinemaCollection, err := c.collectionRepo.GetCollectionByTag(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionByTag")
	}

	return inCinemaCollection, nil
}
