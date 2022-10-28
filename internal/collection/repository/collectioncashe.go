package repository

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetPopular(ctx context.Context) (models.Collection, error)
	GetInCinema(ctx context.Context) (models.Collection, error)
}

// collectionCache is implementation repository of collection
// in memory corresponding to the CollectionService interface.
type collectionCache struct {
	Popular  models.Collection
	InCinema models.Collection
	mu       *sync.RWMutex
}

// NewCollectionCache is constructor for collectionCache. Accepts paths to data collection.
func NewCollectionCache(pathPopular string, pathInCinema string) CollectionRepository {
	res := &collectionCache{
		mu: &sync.RWMutex{},
	}

	res.FillRepo(pathPopular, "popular")
	res.FillRepo(pathInCinema, "in_cinema")

	return res
}

// GetPopular it gives away popular movies from the repository.
func (c *collectionCache) GetPopular(ctx context.Context) (models.Collection, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.Popular.Films) == 0 {
		return models.Collection{}, errors.ErrFilmNotFound
	}

	return c.Popular, nil
}

// GetInCinema it gives away movies in cinema from the repository.
func (c *collectionCache) GetInCinema(ctx context.Context) (models.Collection, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.InCinema.Films) == 0 {
		return models.Collection{}, errors.ErrFilmsNotFound
	}

	return c.InCinema, nil
}

// FillRepo for filling repository from file by path.
func (c *collectionCache) FillRepo(path string, storage string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("FillRepoCollection: can't get data from file", err)
	}

	var collection models.Collection

	err = json.Unmarshal(file, &collection)
	if err != nil {
		logrus.Error("FillRepoCollection: can't Unmarshal data from file", err)
	}

	if storage == "popular" {
		c.Popular = collection

		return
	}

	if storage == "in_cinema" {
		c.InCinema = collection

		return
	}
}
