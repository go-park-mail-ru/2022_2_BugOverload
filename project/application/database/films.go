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
	return &FilmStorage{storage: map[uint]structs.Film{
		0: {
			ID:        0,
			Name:      "Дюна",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/dune_poster.jpg",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		1: {
			ID:        1,
			Name:      "Убить Билла",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/8.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		2: {
			ID:        2,
			Name:      "Головокружение",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/9.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		3: {
			ID:        3,
			Name:      "Доказательство смерти",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/5.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		4: {
			ID:        4,
			Name:      "Чунгингский экспресс",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/7.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		5: {
			ID:        5,
			Name:      "Девушка с татуировкой дракона",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/6.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		6: {
			ID:        6,
			Name:      "Дюна",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/dune_poster.jpg",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		7: {
			ID:        7,
			Name:      "Человек",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/1.png",
			Genres:    []string{"Документальный", "Смотрю и плачу"},
		},
		8: {
			ID:        8,
			Name:      "Люси",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/2.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		9: {
			ID:        9,
			Name:      "Властелин колец. Братство кольца",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/3.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		10: {
			ID:        10,
			Name:      "Дом, который построил Джек",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/4.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		11: {
			ID:        11,
			Name:      "Доказательство смерти",
			YearProd:  "2021",
			Rating:    "7.1",
			PosterVer: "asserts/img/posters/5.png",
			Genres:    []string{"Фэнтези", "Приключения"},
		},
		12: {
			ID:               12,
			Name:             "Звёздные войны. Эпизод IV: Новая надежда",
			ShortDescription: "Может хватит бухтеть и дестабилизировать ситуацию в стране? Световой меч делает вжух-вжух",
			YearProd:         "2021",
			Rating:           "7.1",
			PosterHor:        "asserts/img/StarWars.jpeg",
			Genres:           []string{"Фэнтези", "Приключения"},
		},
		13: {
			ID:               13,
			Name:             "Дюна",
			ShortDescription: "Ну типо по пустыням ходят, а ещё черви там всякие делают уууу",
			YearProd:         "2021",
			Rating:           "7.1",
			PosterHor:        "asserts/img/dune.jpg",
			Genres:           []string{"Фэнтези", "Приключения"},
		},
	}}
}

// CheckExist is method to check the existence of such a film in the database
func (fs *FilmStorage) CheckExist(filmID uint) error {
	_, ok := fs.storage[filmID]
	if !ok {
		return errorshandlers.ErrFilmNotFound
	}

	return nil
}

// AddFilm is method for creating a film in database
func (fs *FilmStorage) AddFilm(f structs.Film) {
	err := fs.CheckExist(f.ID)

	if err != nil {
		fs.storage[f.ID] = f
	}
}

// GetFilm return film using film_id (primary key)
func (fs *FilmStorage) GetFilm(filmID uint) (structs.Film, error) {
	if err := fs.CheckExist(filmID); err != nil {
		return structs.Film{}, err
	}
	return fs.storage[filmID], nil
}

// GetStorageLen return films count in storage
func (fs *FilmStorage) GetStorageLen() int {
	return len(fs.storage)
}
