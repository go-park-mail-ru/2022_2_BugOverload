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
func (fs *FilmStorage) FillFilmStorage() {
	// First collection
	fs.AddFilm(structs.Film{
		ID:        0,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        1,
		Name:      "Убить Билла",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/8.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        2,
		Name:      "Головокружение",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/9.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        3,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        4,
		Name:      "Чунгингский экспресс",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/7.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        5,
		Name:      "Девушка с татуировкой дракона",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/6.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	// Second collection
	fs.AddFilm(structs.Film{
		ID:        6,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        7,
		Name:      "Человек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/1.png",
		Genres:    []string{"Документальный", "Смотрю и плачу"},
	})
	fs.AddFilm(structs.Film{
		ID:        8,
		Name:      "Люси",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/2.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        9,
		Name:      "Властелин колец. Братство кольца",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/3.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        10,
		Name:      "Дом, который построил Джек",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/4.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:        11,
		Name:      "Доказательство смерти",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/5.png",
		Genres:    []string{"Фэнтези", "Приключения"},
	})
	// Third collection (poster)
	fs.AddFilm(structs.Film{
		ID:          12,
		Name:        "Звёздные войны. Эпизод IV: Новая надежда",
		Description: "Может хватит бухтеть и дестабилизировать ситуацию в стране? Световой меч делает вжух-вжух",
		YearProd:    "2021",
		Rating:      "7.1",
		PosterVer:   "asserts/img/StarWars.jpeg",
		Genres:      []string{"Фэнтези", "Приключения"},
	})
	fs.AddFilm(structs.Film{
		ID:          13,
		Name:        "Дюна",
		Description: "Ну типо по пустыням ходят, а ещё черви там всякие делают уууу",
		YearProd:    "2021",
		Rating:      "7.1",
		PosterVer:   "asserts/img/dune.jpg",
		Genres:      []string{"Фэнтези", "Приключения"},
	})
}
