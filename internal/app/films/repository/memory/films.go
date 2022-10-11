package memory

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
)

type filmsRepo struct {
	storage []models.Film
	mu      *sync.Mutex
}

func NewFilmRepo(mu *sync.Mutex, path string) interfaces.FilmsRepository {
	res := &filmsRepo{
		mu: mu,
	}

	res.FillRepo(path)

	return res
}

func (fs *filmsRepo) FillRepo(path string) {
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

func (fs *filmsRepo) CheckExist(filmID uint) bool {
	return filmID <= uint(fs.GetStorageCapacity())
}

func (fs *filmsRepo) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

func (fs *filmsRepo) GetStorageCapacity() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return len(fs.storage)
}

func (fs *filmsRepo) GerRecommendation(ctx context.Context, user *models.User) (models.Film, error) {
	randIndex := utils.Rand(fs.GetStorageCapacity())

	filmRecommendation := fs.storage[randIndex]

	return filmRecommendation, nil
}
