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
	GetPopular(ctx context.Context) ([]models.Film, error)
	GetInCinema(ctx context.Context) ([]models.Film, error)
}

// collectionCache is implementation repository of collection
// in memory corresponding to the CollectionService interface.
type collectionCache struct {
	storagePopular  []models.Film
	storageInCinema []models.Film
	mu              *sync.RWMutex
}

// NewCollectionCache is constructor for collectionCache. Accepts paths to data collection.
func NewCollectionCache(pathPopular string, pathInCinema string) CollectionRepository {
	res := &collectionCache{
		make([]models.Film, 0),
		make([]models.Film, 0),
		&sync.RWMutex{},
	}

	res.FillRepo(pathPopular, "popular")
	res.FillRepo(pathInCinema, "in_cinema")

	return res
}

// GetPopular it gives away popular movies from the repository.
func (c *collectionCache) GetPopular(ctx context.Context) ([]models.Film, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.storagePopular) == 0 {
		return []models.Film{}, errors.ErrFilmNotFound
	}

	return c.storagePopular, nil
}

// GetInCinema it gives away movies in cinema from the repository.
func (c *collectionCache) GetInCinema(ctx context.Context) ([]models.Film, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.storageInCinema) == 0 {
		return []models.Film{}, errors.ErrFilmsNotFound
	}

	return c.storageInCinema, nil
}

// FillRepo for filling repository from file by path.
func (c *collectionCache) FillRepo(path string, storage string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("FillRepoCollection: can't get data from file")
	}

	var films []models.Film

	err = json.Unmarshal(file, &films)
	if err != nil {
		logrus.Error("FillRepoCollection: can't Unmarshal data from file")
	}

	if storage == "popular" {
		c.storagePopular = films

		return
	}

	if storage == "in_cinema" {
		c.storageInCinema = films

		return
	}
}
