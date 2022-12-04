package collection_test

import (
	"context"
	"fmt"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
)

// Sequence
// GetNewFilms
// GetGenresButch (GetGenresFilmBatchBegin + setID + GetGenresFilmBatchEnd)
// GetProdCountriesButch (GetProdCountriesFilmBatchBegin + setID + GetProdCountriesBatchEnd)
// GetDirectorsButch (GetDirectorsFilmBatchBegin + setIDRes + GetDirectorsFilmBatchEnd)

func TestCollection_GetPremieresCollection_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetPremiersCollectionParams{
		CountFilms: 2,
		Delimiter:  0,
	}

	expectedFilms := []modelsGlobal.Film{{
		ID:              1,
		Name:            "Игра престолов",
		ProdDate:        "2013",
		PosterVer:       "12",
		Rating:          9.2,
		DurationMinutes: 110,
		Description:     "some Description",
		Genres:          []string{"фантастика", "драма"},
		ProdCountries:   []string{"США", "Канада"},
		Actors:          []modelsGlobal.FilmActor{},
		Artists:         []modelsGlobal.FilmPerson{},
		Directors: []modelsGlobal.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []modelsGlobal.FilmPerson{},
		Producers: []modelsGlobal.FilmPerson{},
		Operators: []modelsGlobal.FilmPerson{},
		Montage:   []modelsGlobal.FilmPerson{},
		Composers: []modelsGlobal.FilmPerson{},
	}}

	expected := modelsGlobal.Collection{
		Films: expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data

	outputMain := []film.ModelSQL{{
		ID:              1,
		Name:            "Игра престолов",
		ProdDate:        "2013",
		PosterVer:       sqltools.NewSQLNullString("12"),
		Rating:          sqltools.NewSQLNullFloat64(9.2),
		Description:     "some Description",
		DurationMinutes: 110,
	}}

	rowsMain := sqlmock.NewRows([]string{"film_id", "name", "prod_date", "poster_ver", "rating", "duration_minutes", "description"})

	for _, val := range outputMain {
		rowsMain = rowsMain.AddRow(
			val.ID,
			val.Name,
			val.ProdDate,
			val.PosterVer,
			val.Rating,
			val.DurationMinutes,
			val.Description)
	}

	// Settings mock
	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetNewFilms)).
		WithArgs(inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnRows(rowsMain)

	// FilmsGenres
	// Create required setup for handling
	rowsFilmsGenres := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expectedFilms {
		for _, genre := range film.Genres {
			rowsFilmsGenres = rowsFilmsGenres.AddRow(film.ID, genre)
		}
	}

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expectedFilms[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsGenres)

	// FilmsProdCountries
	// Create required setup for handling
	rowsFilmsProdCountries := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expectedFilms {
		for _, country := range film.ProdCountries {
			rowsFilmsProdCountries = rowsFilmsProdCountries.AddRow(film.ID, country)
		}
	}

	// Settings mock
	query = film.GetProdCountriesFilmBatchBegin + strconv.Itoa(expectedFilms[0].ID) + film.GetProdCountriesBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsProdCountries)

	// FilmsDirectors
	// Create required setup for handling
	rowsFilmsDirectors := sqlmock.NewRows([]string{"film_id", "person_id", "name"})

	for _, film := range expectedFilms {
		for _, director := range film.Directors {
			rowsFilmsDirectors = rowsFilmsDirectors.AddRow(film.ID, director.ID, director.Name)
		}
	}

	// Settings mock
	query = film.GetDirectorsFilmBatchBegin + strconv.Itoa(expectedFilms[0].ID) + film.GetDirectorsFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsDirectors)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetPremieresCollection(ctx, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}
