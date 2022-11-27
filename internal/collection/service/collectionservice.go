package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

//go:generate mockgen -source collectionservice.go -destination mocks/mockcollectionservice.go -package mockCollectionService

// CollectionService provides universal service for work with collection.
type CollectionService interface {
	GetStdCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error)
	GetPremieresCollection(ctx context.Context, params *constparams.PremiersCollectionParams) (models.Collection, error)
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
func (c *collectionService) GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	var ok bool

	params.Key, ok = constparams.TagsMap[params.Key]
	if !ok {
		return models.Collection{}, stdErrors.Wrap(errors.ErrTagNotFound, "GetCollectionByTag")
	}

	tagCollection, err := c.collectionRepo.GetCollectionByTag(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionByTag")
	}

	return tagCollection, nil
}

// GetCollectionByGenre is the service that accesses the interface CollectionRepository
func (c *collectionService) GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	var ok bool

	params.Key, ok = constparams.GenresMap[params.Key]
	if !ok {
		return models.Collection{}, stdErrors.Wrap(errors.ErrGenreNotFound, "GetCollectionByGenre")
	}

	collection, err := c.collectionRepo.GetCollectionByGenre(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionByGenre")
	}

	return collection, nil
}

// GetStdCollection is the service that accesses the interface CollectionRepository
func (c *collectionService) GetStdCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	var collection models.Collection
	var err error

	switch params.Target {
	case constparams.CollectionTargetTag:
		collection, err = c.GetCollectionByTag(ctx, params)
	case constparams.CollectionTargetGenre:
		collection, err = c.GetCollectionByGenre(ctx, params)
	default:
		return models.Collection{}, stdErrors.Wrap(errors.ErrNotFindSuchTarget, "GetStdCollection")
	}

	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetStdCollection")
	}

	return collection, nil
}

// GetUserCollections is the service that accesses the interface CollectionRepository
func (c *collectionService) GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error) {
	collection, err := c.collectionRepo.GetUserCollections(ctx, user, params)
	if err != nil {
		return []models.Collection{}, stdErrors.Wrap(err, "GetUserCollections")
	}

	return collection, nil
}

// GetPremieresCollection is the service that accesses the interface CollectionRepository
func (c collectionService) GetPremieresCollection(ctx context.Context, params *constparams.PremiersCollectionParams) (models.Collection, error) {
	collection, err := c.collectionRepo.GetPremieresCollection(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetPremieresCollection")
	}

	return collection, nil
}
