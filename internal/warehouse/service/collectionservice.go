package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
)

//go:generate mockgen -source collectionservice.go -destination mocks/mockcollectionservice.go -package mockWarehouseService

// CollectionService provides universal service for work with collection.
type CollectionService interface {
	GetStdCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetPremieresCollection(ctx context.Context, params *constparams.GetPremiersCollectionParams) (models.Collection, error)
	GetCollectionAuthorized(ctx context.Context, user *models.User, params *constparams.CollectionGetFilmsRequestParams) (models.Collection, error)
	GetCollectionNotAuthorized(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (models.Collection, error)
}

// collectionService is implementation for collection service corresponding to the CollectionService interface.
type collectionService struct {
	collectionRepo collection.Repository
}

// NewCollectionService is constructor for collectionService.
// Accepts Repository interfaces.
func NewCollectionService(cr collection.Repository) CollectionService {
	return &collectionService{
		collectionRepo: cr,
	}
}

// GetCollectionByTag is the service that accesses the interface Repository
func (c *collectionService) GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	var ok bool

	params.Key, ok = constparams.TagsMap[params.Key]
	if !ok {
		return models.Collection{}, stdErrors.Wrap(errors.ErrTagNotFound, "GetCollectionByTag")
	}

	collection, err := c.collectionRepo.GetCollectionByTag(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionByTag")
	}

	return collection, nil
}

// GetCollectionByGenre is the service that accesses the interface Repository
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

// GetStdCollection is the service that accesses the interface Repository
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

// GetPremieresCollection is the service that accesses the interface Repository
func (c *collectionService) GetPremieresCollection(ctx context.Context, params *constparams.GetPremiersCollectionParams) (models.Collection, error) {
	collection, err := c.collectionRepo.GetPremieresCollection(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetPremieresCollection")
	}

	collection.Name = "Премьеры"

	return collection, nil
}

func (c *collectionService) GetCollectionAuthorized(ctx context.Context, user *models.User, params *constparams.CollectionGetFilmsRequestParams) (models.Collection, error) {
	var collection models.Collection

	isAuthor, err := c.collectionRepo.CheckUserIsAuthor(ctx, user, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionAuthorized")
	}

	if !isAuthor {
		var isPublic bool
		isPublic, err = c.collectionRepo.CheckCollectionIsPublic(ctx, params)
		if err != nil {
			return models.Collection{}, stdErrors.Wrap(err, "GetCollectionAuthorized")
		}
		if !isPublic {
			return models.Collection{}, stdErrors.Wrap(errors.ErrCollectionIsNotPublic, "GetCollectionAuthorized")
		}
	}

	collection, err = c.collectionRepo.GetCollection(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionAuthorized")
	}
	if !isAuthor {
		collection.Author.ID = 0
	}

	return collection, nil
}

func (c *collectionService) GetCollectionNotAuthorized(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (models.Collection, error) {
	var collection models.Collection

	isPublic, err := c.collectionRepo.CheckCollectionIsPublic(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionNotAuthorized")
	}
	if !isPublic {
		return models.Collection{}, stdErrors.Wrap(errors.ErrCollectionIsNotPublic, "GetCollectionNotAuthorized")
	}

	collection, err = c.collectionRepo.GetCollection(ctx, params)
	if err != nil {
		return models.Collection{}, stdErrors.Wrap(err, "GetCollectionNotAuthorized")
	}

	collection.Author.ID = 0

	return collection, nil
}
