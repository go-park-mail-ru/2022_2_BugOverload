package database_test

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilmStorage(t *testing.T) {
	fs := database.NewFilmStorage()

	if fs.GetStorageLen() != 14 {
		t.Errorf("Invalid film storage len, [%d], expected: 14", fs.GetStorageLen())
	}

	testFilm := structs.Film{
		ID:        0,
		Name:      "Дюна",
		YearProd:  "2021",
		Rating:    "7.1",
		PosterVer: "asserts/img/posters/dune_poster.jpg",
		Genres:    []string{"Фэнтези", "Приключения"},
	}

	err := fs.CheckExist(testFilm.ID)
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}
}

func TestFilmStorageAdd(t *testing.T) {
	fs := database.NewFilmStorage()

	newFilm := structs.Film{
		ID:   14,
		Name: "TestFilm",
	}

	fs.AddFilm(newFilm)

	if err := fs.CheckExist(newFilm.ID); err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	getNewFilm, err := fs.GetFilm(14)

	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	if !cmp.Equal(newFilm, getNewFilm) {
		t.Errorf("Error getFilm, expected [%v], found [%v]", newFilm, getNewFilm)
	}
}

func TestFilmStorageGet(t *testing.T) {
	fs := database.NewFilmStorage()

	_, err := fs.GetFilm(15)

	if err != errorshandlers.ErrFilmNotFound {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}
}
