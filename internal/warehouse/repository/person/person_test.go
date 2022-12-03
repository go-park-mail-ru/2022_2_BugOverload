package person_test

import (
	"context"
	"fmt"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/person"
)

// Sequence
// GetPersonByID
// GetPersonBestFilms
// Array GetShortSerialByID
// GetGenresButch (GetGenresFilmBatchBegin + setID + GetGenresFilmBatchEnd)
// GetPersonImages
// GetPersonProfessions
// GetPersonGenres

func TestPerson_GetByID_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputPerson := &modelsGlobal.Person{
		ID: 1,
	}

	// Images, films which return can be less than in request params
	inputParams := &constparams.GetPersonParams{
		CountFilms:  2,
		CountImages: 3,
	}

	expectedFilms := []modelsGlobal.Film{{
		ID:           1,
		Name:         "Игра престолов",
		OriginalName: "Game of thrones",
		ProdDate:     "2013",
		EndYear:      "2019",
		Type:         "serial",
		PosterVer:    "12",
		Rating:       9.2,
		Genres:       []string{"фантастика", "драма"},
		Actors:       []modelsGlobal.FilmActor{},
		Artists:      []modelsGlobal.FilmPerson{},
		Directors:    []modelsGlobal.FilmPerson{},
		Writers:      []modelsGlobal.FilmPerson{},
		Producers:    []modelsGlobal.FilmPerson{},
		Operators:    []modelsGlobal.FilmPerson{},
		Montage:      []modelsGlobal.FilmPerson{},
		Composers:    []modelsGlobal.FilmPerson{},
	}}

	expected := modelsGlobal.Person{
		Name:         "Вин дизель",
		Birthday:     "1967.12.12",
		GrowthMeters: 1.92,
		OriginalName: "Vin Dizel",
		Avatar:       "23",
		Images:       []string{"1", "2"},
		Genres:       []string{"боевик", "триллер"},
		Professions:  []string{"режисер", "продюсер"},
		CountFilms:   1,
		Gender:       "male",
		BestFilms:    expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Person Main
	// Data

	birthday, _ := time.Parse(constparams.DateFormat, "1967.12.12")

	outputMain := person.ModelSQL{
		Name:         "Вин дизель",
		Birthday:     birthday,
		GrowthMeters: 1.92,
		OriginalName: sqltools.NewSQLNullString("Vin Dizel"),
		Avatar:       sqltools.NewSQLNullString("23"),
		Death:        sqltools.NewSQLNNullDate("", constparams.DateFormat),
		Gender:       sqltools.NewSQLNullString("male"),
		CountFilms:   sqltools.NewSQLNullInt32(1),
	}

	// Create required setup for handling
	rowMain := sqlmock.NewRows([]string{"name", "birthday", "growth_meters", "original_name", "avatar", "death", "gender", "count_films"})

	rowMain = rowMain.AddRow(
		outputMain.Name,
		outputMain.Birthday,
		outputMain.GrowthMeters,
		outputMain.OriginalName,
		outputMain.Avatar,
		outputMain.Death,
		outputMain.Gender,
		outputMain.CountFilms)

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(person.GetPersonByID)).
		WithArgs(inputPerson.ID). // Values in query
		WillReturnRows(rowMain)

	// Parts
	// Films
	// Data
	filmID := 1

	outputFilms := []film.ModelSQL{{
		ID:           filmID,
		Name:         "Игра престолов",
		OriginalName: sqltools.NewSQLNullString("Game of thrones"),
		ProdDate:     "2013",
		EndYear:      sqltools.NewSQLNNullDate("", ""),
		PosterVer:    sqltools.NewSQLNullString("12"),
		Rating:       sqltools.NewSQLNullFloat64(9.2),
		Type:         sqltools.NewSQLNullString("serial"),
	}}

	// Create required setup for handling
	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original_name", "prod_date", "poster_ver", "type", "rating"})

	for _, val := range outputFilms {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			val.OriginalName,
			val.ProdDate,
			val.PosterVer,
			val.Type,
			val.Rating)
	}

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(person.GetPersonBestFilms)).
		WithArgs(inputPerson.ID, inputParams.CountFilms). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(filmID). // Values in query
		WillReturnRows(rowsSerials)

	// FilmsGenres
	// Create required setup for handling
	rowsFilmsGenres := sqlmock.NewRows([]string{"film_id", "genre"})

	rowsFilmsGenres = rowsFilmsGenres.AddRow(filmID, "фантастика")
	rowsFilmsGenres = rowsFilmsGenres.AddRow(filmID, "драма")

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(filmID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsGenres)

	// Images
	// Data
	// Create required setup for handling
	rowsPersonImages := sqlmock.NewRows([]string{"image_key"})

	rowsPersonImages = rowsPersonImages.AddRow("1")
	rowsPersonImages = rowsPersonImages.AddRow("2")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(person.GetPersonImages)).
		WithArgs(inputPerson.ID, inputParams.CountImages). // Values in query
		WillReturnRows(rowsPersonImages)

	// Professions
	// Data
	// Create required setup for handling
	rowsPersonProfessions := sqlmock.NewRows([]string{"name"})

	rowsPersonProfessions = rowsPersonProfessions.AddRow("режисер")
	rowsPersonProfessions = rowsPersonProfessions.AddRow("продюсер")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(person.GetPersonProfessions)).
		WithArgs(inputPerson.ID). // Values in query
		WillReturnRows(rowsPersonProfessions)

	// Genres
	// Data
	// Create required setup for handling
	rowsPersonGenres := sqlmock.NewRows([]string{"name"})

	rowsPersonGenres = rowsPersonGenres.AddRow("боевик")
	rowsPersonGenres = rowsPersonGenres.AddRow("триллер")

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(person.GetPersonGenres)).
		WithArgs(inputPerson.ID). // Values in query
		WillReturnRows(rowsPersonGenres)

	// Init
	repo := person.NewPersonPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetPersonByID(ctx, inputPerson, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}
