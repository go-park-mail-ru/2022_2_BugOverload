package memory

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

type FilmStorage struct {
	storage []models.Film
	mu      *sync.Mutex
}

func NewFilmStorage() *FilmStorage {
	res := &FilmStorage{
		mu: &sync.Mutex{},
	}

	return res
}

func (fs *FilmStorage) FillStorage(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Error("can't get data from file")
	}

	var films []models.Film

	err = json.Unmarshal(file, &films)
	if err != nil {
		logrus.Error("can't Unmarshal data from file")
	}

	fs.storage = films
}

func (fs *FilmStorage) CheckExist(filmID uint) bool {
	return filmID <= uint(fs.GetStorageLen())
}

func (fs *FilmStorage) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

func (fs *FilmStorage) GetFilm(filmID uint) (models.Film, error) {
	if !fs.CheckExist(filmID) {
		return models.Film{}, errors.ErrFilmNotFound
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	return fs.storage[filmID], nil
}

func (fs *FilmStorage) GetStorageLen() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return len(fs.storage)
}
