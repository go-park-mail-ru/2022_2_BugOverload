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

// FilmsRepository provides the versatility of films repositories.
type FilmsRepository interface {
	GerRecommendation(ctx context.Context) (models.Film, error)
}

// filmsCash is implementation repository of films in memory corresponding to the FilmsRepository interface.
type filmsCash struct {
	storage []models.Film
	mu      *sync.Mutex
}

// NewFilmCash is constructor for filmsCash. Accepts mutex and path to data films.
func NewFilmCash(path string) FilmsRepository {
	res := &filmsCash{
		mu: &sync.Mutex{},
	}

	res.FillRepo(path)

	return res
}

// FillRepo for filling repository from file by path.
func (fs *filmsCash) FillRepo(path string) {
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

// CheckExist is a check for the existence of such a film by ID.
func (fs *filmsCash) CheckExist(filmID uint) bool {
	return filmID <= uint(fs.GetStorageCapacity())
}

func (fs *filmsCash) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

func (fs *filmsCash) GetStorageCapacity() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return len(fs.storage)
}

// GerRecommendation it gives away recommendation film from the repository for unauthorized users.
func (fs *filmsCash) GerRecommendation(ctx context.Context) (models.Film, error) {
	randIndex := pkg.Rand(fs.GetStorageCapacity())

	filmRecommendation := fs.storage[randIndex]

	return filmRecommendation, nil
}
