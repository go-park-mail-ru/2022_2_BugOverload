package memory

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

// FilmStorage is TMP impl database for films, where key = film_id
type FilmStorage struct {
	storage []models.Film
	mu      *sync.Mutex
}

// NewFilmStorage is constructor for FilmStorage
func NewFilmStorage() *FilmStorage {
	res := &FilmStorage{
		mu: &sync.Mutex{},
	}

	return res
}

// FillStorage for getting data from file
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

	logrus.Info(len(films))

	fs.storage = films
}

// CheckExist is method to check the existence of such a film in the database
func (fs *FilmStorage) CheckExist(filmID uint) bool {
	return filmID <= uint(fs.GetStorageLen())
}

// AddFilm is method for creating a film in database
func (fs *FilmStorage) AddFilm(f models.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

// GetFilm return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (models.Film, error) {
	if !fs.CheckExist(filmID) {
		return models.Film{}, errors.ErrFilmNotFound
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	return fs.storage[filmID], nil
}

// GetStorageLen return films count in storage
func (fs *FilmStorage) GetStorageLen() int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return len(fs.storage)
}
