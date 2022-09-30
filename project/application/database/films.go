package database

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// FilmStorage is TMP impl database for films, where key = film_id
type FilmStorage struct {
	storage map[uint]structs.Film
}

// NewFilmStorage is constructor for FilmStorage
func NewFilmStorage() *FilmStorage {
	return &FilmStorage{make(map[uint]structs.Film)}
}

// CheckExist is method to check the existence of such a film in the database
func (fs *FilmStorage) CheckExist(filmID uint) error {
	_, ok := fs.storage[filmID]
	if !ok {
		return errorshandlers.ErrFilmNotFound
	}

	return nil
}

// Create is method for creating a film in database
func (fs *FilmStorage) AddFilm(f structs.Film) {
	err := fs.CheckExist(f.ID)

	if err != nil {
		fs.storage[f.ID] = f
	}
}

// Return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (structs.Film, error) {
	if err := fs.CheckExist(filmID); err != nil {
		return structs.Film{}, err
	}
	return fs.storage[filmID], nil
}

func (fs *FilmStorage) GetStorageLen() int {
	return len(fs.storage)
}

// Temporary function, filling local storage
func (fs *FilmStorage) FillFilmStoragePartOne() {
	// First collection
	var currentID uint
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Убить Билла",
		YearProd:  "2021",
		Rating:    "7.1",
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
		Rating:    "7.1",
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
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/6.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
}

func (fs *FilmStorage) FillFilmStoragePartTwo() {
	var currentID uint = 6
	// Second collection
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "7.1",
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
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/2.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Властелин колец. Братство кольца",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/3.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Дом, который построил Джек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/4.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:        currentID,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	currentID++
	// Third collection (poster)
	fs.AddFilm(structs.Film{
		ID:          currentID,
		Name:        "Звёздные войны. Эпизод IV: Новая надежда",
		Description: "Может хватит бухтеть и дестабилизировать ситуацию в стране? Световой меч делает вжух-вжух",
		YearProd:    "2021",
		Rating:      "7.1",
		PosterVer:   "asserts/img/StarWars.jpeg",
		Genres:      []string{"Фэнтези", "Приключения"},
	})
	currentID++
	fs.AddFilm(structs.Film{
		ID:          currentID,
		Name:        "Дюна",
		Description: "Ну типо по пустыням ходят, а ещё черви там всякие делают уууу",
		YearProd:    "2021",
		Rating:      "7.1",
		PosterVer:   "asserts/img/dune.jpg",
		Genres:      []string{"Фэнтези", "Приключения"},
	})
}
