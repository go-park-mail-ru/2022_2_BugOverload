package memory

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"os"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type collectionRepo struct {
	storagePopular  []models.Film
	storageInCinema []models.Film
	mu              *sync.Mutex
}

func NewCollectionRepo(mu *sync.Mutex, pathPopular string, pathInCinema string) interfaces.CollectionRepository {
	res := &collectionRepo{
		make([]models.Film, 0),
		make([]models.Film, 0),
		mu,
	}

	res.FillRepo(pathPopular, "popular")
	res.FillRepo(pathInCinema, "in_cinema")

	return res
}

func (c *collectionRepo) GetPopular(ctx context.Context) ([]models.Film, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.storagePopular) == 0 {
		return []models.Film{}, errors.ErrFilmNotFound
	}

	return c.storagePopular, nil
}

func (c *collectionRepo) GetInCinema(ctx context.Context) ([]models.Film, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.storageInCinema) == 0 {
		return []models.Film{}, errors.ErrFilmsNotFound
	}

	return c.storageInCinema, nil
}

func (c *collectionRepo) FillRepo(path string, storage string) {
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
		c.storagePopular = append(c.storagePopular, films...)

		return
	}

	if storage == "in_cinema" {
		c.storageInCinema = append(c.storageInCinema, films...)

		return
	}
}
