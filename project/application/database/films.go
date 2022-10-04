package database

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
	"sync"
)

// FilmStorage is TMP impl database for films, where key = film_id
type FilmStorage struct {
	storage map[uint]structs.Film
	mu      *sync.Mutex
}

// NewFilmStorage is constructor for FilmStorage
func NewFilmStorage() *FilmStorage {
	return &FilmStorage{
		make(map[uint]structs.Film),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a film in the database
func (fs *FilmStorage) CheckExist(filmID uint) bool {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	_, ok := fs.storage[filmID]
	return ok
}

// AddFilm is method for creating a film in database
func (fs *FilmStorage) AddFilm(f structs.Film) {
	if !fs.CheckExist(f.ID) {
		fs.mu.Lock()
		defer fs.mu.Unlock()

		fs.storage[f.ID] = f
	}
}

// GetFilm return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (structs.Film, error) {
	if !fs.CheckExist(filmID) {
		return structs.Film{}, errorshandlers.ErrFilmNotFound
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

// FillFilmStoragePartOne is temporary function, filling local storage
func (fs *FilmStorage) FillFilmStoragePartOne() {
	// First collection
	var currentID uint
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "9.9",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Убить Билла",
		YearProd:  "2021",
		Rating:    "2.0",
		PosterVer: "asserts/img/posters/8.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Головокружение",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/9.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "3.3",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Чунгингский экспресс",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/7.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Девушка с татуировкой дракона",
		YearProd:  "2021",
		Rating:    "5.7",
		PosterVer: "asserts/img/posters/6.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
}

// FillFilmStoragePartTwo is temporary function, filling local storage
func (fs *FilmStorage) FillFilmStoragePartTwo() {
	var currentID uint = 6
	// Second collection
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "9.9",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Человек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/1.png",
		Genres:    []string{"Документальный", "Смотрю и плачу"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Люси",
		YearProd:  "2021",
		Rating:    "8.9",
		PosterVer: "asserts/img/posters/2.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Властелин колец. Братство кольца",
		YearProd:  "2021",
		Rating:    "8.4",
		PosterVer: "asserts/img/posters/3.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дом, который построил Джек",
		YearProd:  "2021",
		Rating:    "7.2",
		PosterVer: "asserts/img/posters/4.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "4.1",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	// Duple
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "9.9",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Человек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/1.png",
		Genres:    []string{"Документальный", "Смотрю и плачу"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Люси",
		YearProd:  "2021",
		Rating:    "8.9",
		PosterVer: "asserts/img/posters/2.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Властелин колец. Братство кольца",
		YearProd:  "2021",
		Rating:    "8.4",
		PosterVer: "asserts/img/posters/3.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дом, который построил Джек",
		YearProd:  "2021",
		Rating:    "7.2",
		PosterVer: "asserts/img/posters/4.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "4.1",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	// Third collection (poster)
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Звёздные войны. Эпизод IV: Новая надежда",
		ShortDescription: "Может хватит бухтеть и дестабилизировать ситуацию в стране? Световой меч делает вжух-вжух",
		YearProd:         "2021",
		PosterHor:        "asserts/img/StarWars.jpeg",
		Genres:           []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Дюна",
		ShortDescription: "Ну типо по пустыням ходят, а ещё черви там всякие делают уууу",
		YearProd:         "2021",
		PosterHor:        "asserts/img/dune.jpg",
		Genres:           []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "Джокер",
		ShortDescription: "Псих вышел погулять",
		YearProd:         "2021",
		PosterHor:        "asserts/img/joker_hor.jpg",
		Genres:           []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:               currentID,
		Name:             "2001: Космическая одисея",
		ShortDescription: "Псих вышел погулять",
		YearProd:         "2021",
		PosterHor:        "asserts/img/space_odyssey_hor.jpg",
		Genres:           []string{"Фэнтези", "Приключения"},
	})
}
