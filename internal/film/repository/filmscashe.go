package repository

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// FilmsRepository provides the versatility of film repositories.
type FilmsRepository interface {
	GerRecommendation(ctx context.Context) (models.Film, error)
}

// filmsCache is implementation repository of film in memory corresponding to the FilmsRepository interface.
type filmsCache struct {
	storage []models.Film
	mu      *sync.RWMutex
}

// NewFilmCache is constructor for filmsCache. Accepts mutex and path to data film.
func NewFilmCache(path string) FilmsRepository {
	res := &filmsCache{
		mu: &sync.RWMutex{},
	}

	res.FillRepo(path)

	return res
}

// FillRepo for filling repository from file by path.
func (fs *filmsCache) FillRepo(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("FillRepoFilms: can't get data from file")
	}

	var films []models.Film

	err = json.Unmarshal(file, &films)
	if err != nil {
		logrus.Error("FillRepoFilms: can't Unmarshal data from file")
	}

	fs.storage = films
}

// CheckExist is a check for the existence of such a film by Key.
func (fs *filmsCache) CheckExist(filmID int) bool {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	return filmID <= fs.GetStorageCapacity()
}

func (fs *filmsCache) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

func (fs *filmsCache) GetStorageCapacity() int {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	return len(fs.storage)
}

// GerRecommendation it gives away recommendation film from the repository for unauthorized users.
func (fs *filmsCache) GerRecommendation(ctx context.Context) (models.Film, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	randIndex := pkg.Rand(fs.GetStorageCapacity())

	filmRecommendation := fs.storage[randIndex]

	return filmRecommendation, nil
}
