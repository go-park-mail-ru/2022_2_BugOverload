package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"
	collectionInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type collectionService struct {
	collectionRepo collectionInterface.CollectionRepository
	contextTimeout time.Duration
}

func NewCollectionService(cr collectionInterface.CollectionRepository, timeout time.Duration) collectionInterface.CollectionService {
	return &collectionService{
		collectionRepo: cr,
		contextTimeout: timeout,
	}
}

func (c collectionService) GetPopular(ctx context.Context) ([]models.Film, error) {
	inCinemaCollection, err := c.collectionRepo.GetPopular(ctx)
	if err != nil {
		return []models.Film{}, stdErrors.Wrap(err, "GetPopular")
	}

	return inCinemaCollection, nil
}

func (c collectionService) GetInCinema(ctx context.Context) ([]models.Film, error) {
	popularCollection, err := c.collectionRepo.GetInCinema(ctx)
	if err != nil {
		return []models.Film{}, stdErrors.Wrap(err, "GetInCinema")
	}

	return popularCollection, nil
}
