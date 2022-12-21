package search_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/search"
)

func TestSearch_SearchFilms_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	expected := []models.Film{{
		ID:            1,
		Name:          "Игра престолов",
		ProdDate:      "2013",
		PosterVer:     "12",
		Rating:        9.2,
		Genres:        []string{"фантастика", "драма"},
		ProdCountries: []string{"США", "Канада"},
		Actors:        []models.FilmActor{},
		Artists:       []models.FilmPerson{},
		Directors: []models.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []models.FilmPerson{},
		Producers: []models.FilmPerson{},
		Operators: []models.FilmPerson{},
		Montage:   []models.FilmPerson{},
		Composers: []models.FilmPerson{},
	}}

	// Input global
	ctx := context.TODO()

	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original name", "prod_date", "poster_ver", "type", "rating"})
	for _, val := range expected {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			sqltools.NewSQLNullString(val.OriginalName),
			val.ProdDate,
			sqltools.NewSQLNullString(val.PosterVer),
			sqltools.NewSQLNullString(val.Type),
			sqltools.NewSQLNullFloat64(val.Rating))
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnRows(rowsFilms)

	// FilmsGenres
	// Create required setup for handling
	rowsFilmsGenres := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expected {
		for _, genre := range film.Genres {
			rowsFilmsGenres = rowsFilmsGenres.AddRow(film.ID, genre)
		}
	}

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsGenres)

	// FilmsProdCountries
	// Create required setup for handling
	rowsFilmsProdCountries := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expected {
		for _, country := range film.ProdCountries {
			rowsFilmsProdCountries = rowsFilmsProdCountries.AddRow(film.ID, country)
		}
	}

	// Settings mock
	query = film.GetProdCountriesFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetProdCountriesBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsProdCountries)

	// FilmsDirectors
	// Create required setup for handling
	rowsFilmsDirectors := sqlmock.NewRows([]string{"film_id", "person_id", "name"})

	for _, film := range expected {
		for _, director := range film.Directors {
			rowsFilmsDirectors = rowsFilmsDirectors.AddRow(film.ID, director.ID, director.Name)
		}
	}

	// Settings mock
	query = film.GetDirectorsFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetDirectorsFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsDirectors)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.SearchFilms(ctx, &inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}

func TestSearch_SearchFilms_DirectorsFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	expected := []models.Film{{
		ID:            1,
		Name:          "Игра престолов",
		ProdDate:      "2013",
		PosterVer:     "12",
		Rating:        9.2,
		Genres:        []string{"фантастика", "драма"},
		ProdCountries: []string{"США", "Канада"},
		Actors:        []models.FilmActor{},
		Artists:       []models.FilmPerson{},
		Directors: []models.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []models.FilmPerson{},
		Producers: []models.FilmPerson{},
		Operators: []models.FilmPerson{},
		Montage:   []models.FilmPerson{},
		Composers: []models.FilmPerson{},
	}}

	// Input global
	ctx := context.TODO()

	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original name", "prod_date", "poster_ver", "type", "rating"})
	for _, val := range expected {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			sqltools.NewSQLNullString(val.OriginalName),
			val.ProdDate,
			sqltools.NewSQLNullString(val.PosterVer),
			sqltools.NewSQLNullString(val.Type),
			sqltools.NewSQLNullFloat64(val.Rating))
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnRows(rowsFilms)

	// FilmsGenres
	// Create required setup for handling
	rowsFilmsGenres := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expected {
		for _, genre := range film.Genres {
			rowsFilmsGenres = rowsFilmsGenres.AddRow(film.ID, genre)
		}
	}

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsGenres)

	// FilmsProdCountries
	// Create required setup for handling
	rowsFilmsProdCountries := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expected {
		for _, country := range film.ProdCountries {
			rowsFilmsProdCountries = rowsFilmsProdCountries.AddRow(film.ID, country)
		}
	}

	// Settings mock
	query = film.GetProdCountriesFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetProdCountriesBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsProdCountries)

	// FilmsDirectors

	// Settings mock
	query = film.GetDirectorsFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetDirectorsFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.SearchFilms(ctx, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestSearch_SearchFilms_ProdCountriesFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	expected := []models.Film{{
		ID:            1,
		Name:          "Игра престолов",
		ProdDate:      "2013",
		PosterVer:     "12",
		Rating:        9.2,
		Genres:        []string{"фантастика", "драма"},
		ProdCountries: []string{"США", "Канада"},
		Actors:        []models.FilmActor{},
		Artists:       []models.FilmPerson{},
		Directors: []models.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []models.FilmPerson{},
		Producers: []models.FilmPerson{},
		Operators: []models.FilmPerson{},
		Montage:   []models.FilmPerson{},
		Composers: []models.FilmPerson{},
	}}

	// Input global
	ctx := context.TODO()

	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original name", "prod_date", "poster_ver", "type", "rating"})
	for _, val := range expected {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			sqltools.NewSQLNullString(val.OriginalName),
			val.ProdDate,
			sqltools.NewSQLNullString(val.PosterVer),
			sqltools.NewSQLNullString(val.Type),
			sqltools.NewSQLNullFloat64(val.Rating))
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnRows(rowsFilms)

	// FilmsGenres
	// Create required setup for handling
	rowsFilmsGenres := sqlmock.NewRows([]string{"film_id", "name"})

	for _, film := range expected {
		for _, genre := range film.Genres {
			rowsFilmsGenres = rowsFilmsGenres.AddRow(film.ID, genre)
		}
	}

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rowsFilmsGenres)

	// FilmsProdCountries

	// Settings mock
	query = film.GetProdCountriesFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetProdCountriesBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.SearchFilms(ctx, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestSearch_SearchFilms_GenresFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	expected := []models.Film{{
		ID:            1,
		Name:          "Игра престолов",
		ProdDate:      "2013",
		PosterVer:     "12",
		Rating:        9.2,
		Genres:        []string{"фантастика", "драма"},
		ProdCountries: []string{"США", "Канада"},
		Actors:        []models.FilmActor{},
		Artists:       []models.FilmPerson{},
		Directors: []models.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []models.FilmPerson{},
		Producers: []models.FilmPerson{},
		Operators: []models.FilmPerson{},
		Montage:   []models.FilmPerson{},
		Composers: []models.FilmPerson{},
	}}

	// Input global
	ctx := context.TODO()

	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original name", "prod_date", "poster_ver", "type", "rating"})
	for _, val := range expected {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			sqltools.NewSQLNullString(val.OriginalName),
			val.ProdDate,
			sqltools.NewSQLNullString(val.PosterVer),
			sqltools.NewSQLNullString(val.Type),
			sqltools.NewSQLNullFloat64(val.Rating))
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnRows(rowsFilms)

	// FilmsGenres

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expected[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.SearchFilms(ctx, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestSearch_SearchFilms_FilmNoRows(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	expected := []models.Film{{
		ID:            1,
		Name:          "Игра престолов",
		ProdDate:      "2013",
		PosterVer:     "12",
		Rating:        9.2,
		Genres:        []string{"фантастика", "драма"},
		ProdCountries: []string{"США", "Канада"},
		Actors:        []models.FilmActor{},
		Artists:       []models.FilmPerson{},
		Directors: []models.FilmPerson{{
			ID:   2,
			Name: "Уеллем Дефо",
		}},
		Writers:   []models.FilmPerson{},
		Producers: []models.FilmPerson{},
		Operators: []models.FilmPerson{},
		Montage:   []models.FilmPerson{},
		Composers: []models.FilmPerson{},
	}}

	// Input global
	ctx := context.TODO()

	rowsFilms := sqlmock.NewRows([]string{"film_id", "name", "original name", "prod_date", "poster_ver", "type", "rating"})
	for _, val := range expected {
		rowsFilms = rowsFilms.AddRow(
			val.ID,
			val.Name,
			sqltools.NewSQLNullString(val.OriginalName),
			val.ProdDate,
			sqltools.NewSQLNullString(val.PosterVer),
			sqltools.NewSQLNullString(val.Type),
			sqltools.NewSQLNullFloat64(val.Rating))
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.SearchFilms(ctx, &inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestSearch_SearchFilms_FilmsFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputParams := constparams.SearchParams{
		Query: "%aboba%",
	}

	// Input global
	ctx := context.TODO()

	mock.
		ExpectQuery(regexp.QuoteMeta(search.SearchFilmsByName)).
		WithArgs(inputParams.Query).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := search.NewSearchPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.SearchFilms(ctx, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}
