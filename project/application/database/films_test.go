package database_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

func TestFilmStorage(t *testing.T) {
	fs := database.NewFilmStorage()

	if fs.GetStorageLen() != 30 {
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

	ok := fs.CheckExist(testFilm.ID)
	if !ok {
		t.Errorf("Err: [%t], expected: true", ok)
	}
}

func TestFilmStorageAdd(t *testing.T) {
	fs := database.NewFilmStorage()

	newFilm := structs.Film{
		ID:   30,
		Name: "TestFilm",
	}

	fs.AddFilm(newFilm)

	if ok := fs.CheckExist(newFilm.ID); !ok {
		t.Errorf("Err: [%t], expected: true", ok)
	}

	getNewFilm, err := fs.GetFilm(30)

	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	if !cmp.Equal(newFilm, getNewFilm) {
		t.Errorf("Error getFilm, expected [%v], found [%v]", newFilm, getNewFilm)
	}
}

func TestFilmStorageGet(t *testing.T) {
	fs := database.NewFilmStorage()

	_, err := fs.GetFilm(125)

	if err != errorshandlers.ErrFilmNotFound {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}
}
