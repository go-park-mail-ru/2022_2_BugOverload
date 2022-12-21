package collection_test

import (
	"context"
	"database/sql"
	"fmt"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
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

func TestCollection_GetPremieresCollection_DirectorsFail(t *testing.T) {
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
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetPremieresCollection(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetPremieresCollection_CountriesFail(t *testing.T) {
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
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetPremieresCollection(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetPremieresCollection_FilmsGenresFail(t *testing.T) {
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

	// Settings mock
	query := film.GetGenresFilmBatchBegin + strconv.Itoa(expectedFilms[0].ID) + film.GetGenresFilmBatchEnd

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetPremieresCollection(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetPremieresCollection_FilmsNoRows(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetPremiersCollectionParams{
		CountFilms: 2,
		Delimiter:  0,
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
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetPremieresCollection(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

// Sequence
// query by sort param (GetFilmsByGenreDate, GetFilmsByGenreRating)
// GetGenresButch (GetGenresFilmBatchBegin + setID + GetGenresFilmBatchEnd)

func TestCollection_GetCollectionByGenre_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "0",
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

	expected := modelsGlobal.Collection{
		Name:  "комедия",
		Films: expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByGenreDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

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

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetCollectionByGenre(ctx, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}

func TestCollection_GetCollectionByGenre_GenresFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "0",
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

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByGenreDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

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
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByGenre(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByGenre_FilmsFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "0",
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByGenreDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByGenre(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByGenre_FilmsNoRows(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "0",
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByGenreDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByGenre(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByGenre_SortRatingOK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "10",
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

	expected := modelsGlobal.Collection{
		Name:  "комедия",
		Films: expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByGenreRating)).
		WithArgs(inputParams.Key, 10.0, inputParams.CountFilms). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

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

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetCollectionByGenre(ctx, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}

func TestCollection_GetCollectionByGenre_InvalidSortParam(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "invalid parameter",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "10",
	}

	// Input global
	ctx := context.TODO()

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByGenre(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByGenre_InvalidDelimeter(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "zxcasd",
	}

	// Input global
	ctx := context.TODO()

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByGenre(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

// Sequence
// query by sort param (GetFilmsByGenreDate, GetFilmsByGenreRating)
// GetGenresButch (GetGenresFilmBatchBegin + setID + GetGenresFilmBatchEnd)

func TestCollection_GetCollectionByTag_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "0",
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

	expected := modelsGlobal.Collection{
		Name:        "популярное",
		Description: "test description",
		Films:       expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

	// Tags description

	rowsDescription := sqlmock.NewRows([]string{"description"})
	rowsDescription.AddRow("test description")

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.GetTagDescription)).
		WithArgs(inputParams.Key).
		WillReturnRows(rowsDescription)

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

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetCollectionByTag(ctx, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}

func TestCollection_GetCollectionByTag_GenresFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "0",
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

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

	// Tags description

	rowsDescription := sqlmock.NewRows([]string{"description"})
	rowsDescription.AddRow("test description")

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.GetTagDescription)).
		WithArgs(inputParams.Key).
		WillReturnRows(rowsDescription)

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
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_FilmsFail(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "0",
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_FilmsNoRows(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "0",
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagDate)).
		WithArgs(inputParams.Key, inputParams.CountFilms, inputParams.Delimiter). // Values in query
		WillReturnError(sql.ErrNoRows)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_SortRatingOK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "10",
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

	expected := modelsGlobal.Collection{
		Name:        "популярное",
		Description: "test description",
		Films:       expectedFilms,
	}

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagRating)).
		WithArgs(inputParams.Key, 10.0, inputParams.CountFilms). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

	// Tags description

	rowsDescription := sqlmock.NewRows([]string{"description"})
	rowsDescription.AddRow("test description")

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.GetTagDescription)).
		WithArgs(inputParams.Key).
		WillReturnRows(rowsDescription)

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

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.GetCollectionByTag(ctx, inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.Equal(t, expected, actual)
}

func TestCollection_GetCollectionByTag_InvalidSortParam(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "invalid parameter",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "10",
	}

	// Input global
	ctx := context.TODO()

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_InvalidDelimeter(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "комедия",
		Delimiter:  "zxcasd",
	}

	// Input global
	ctx := context.TODO()

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_TagDescriptionScanErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "10",
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

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagRating)).
		WithArgs(inputParams.Key, 10.0, inputParams.CountFilms). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

	// Tags description

	rowsDescription := sqlmock.NewRows([]string{})
	rowsDescription.AddRow()

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.GetTagDescription)).
		WithArgs(inputParams.Key).
		WillReturnRows(rowsDescription)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_GetCollectionByTag_TagDescriptionErr(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	// Global input output
	inputParams := &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 2,
		Key:        "популярное",
		Delimiter:  "10",
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

	// Input global
	ctx := context.TODO()

	// Films Main
	// Data
	outputFilms := []film.ModelSQL{{
		ID:           1,
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
		ExpectQuery(regexp.QuoteMeta(collection.GetFilmsByTagRating)).
		WithArgs(inputParams.Key, 10.0, inputParams.CountFilms). // Values in query
		WillReturnRows(rowsFilms)

	// Serials
	// Data
	rowsSerials := sqlmock.NewRows([]string{"end_year"})

	endYear, _ := time.Parse(constparams.OnlyDate, "2019")

	rowsSerials = rowsSerials.AddRow(endYear)

	mock.
		ExpectQuery(regexp.QuoteMeta(film.GetShortSerialByID)).
		WithArgs(expectedFilms[0].ID). // Values in query
		WillReturnRows(rowsSerials)

	// Tags description
	mock.
		ExpectQuery(regexp.QuoteMeta(collection.GetTagDescription)).
		WithArgs(inputParams.Key).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.GetCollectionByTag(ctx, inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_CheckUserIsAuthor_OK(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputUser := modelsGlobal.User{
		ID: 1,
	}

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
	}

	// Input global
	ctx := context.TODO()

	rowsResult := sqlmock.NewRows([]string{"exist"})
	rowsResult.AddRow(true)

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.CheckUserIsCollectionAuthor)).
		WithArgs(inputParams.CollectionID, inputUser.ID).
		WillReturnRows(rowsResult)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	actual, err := repo.CheckUserIsAuthor(ctx, &inputUser, &inputParams)
	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))

	// Check actual
	require.True(t, actual)
}

func TestCollection_CheckUserIsAuthor_ScanError(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputUser := modelsGlobal.User{
		ID: 1,
	}

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
	}

	// Input global
	ctx := context.TODO()

	rowsResult := sqlmock.NewRows([]string{})
	rowsResult.AddRow()

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.CheckUserIsCollectionAuthor)).
		WithArgs(inputParams.CollectionID, inputUser.ID).
		WillReturnRows(rowsResult)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.CheckUserIsAuthor(ctx, &inputUser, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}

func TestCollection_CheckUserIsAuthor_Error(t *testing.T) {
	// Init sqlmock
	db, mock, err := sqlmock.New()
	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
	defer db.Close()

	inputUser := modelsGlobal.User{
		ID: 1,
	}

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
	}

	// Input global
	ctx := context.TODO()

	mock.
		ExpectQuery(regexp.QuoteMeta(collection.CheckUserIsCollectionAuthor)).
		WithArgs(inputParams.CollectionID, inputUser.ID).
		WillReturnError(errors.ErrWorkDatabase)

	// Init
	repo := collection.NewCollectionPostgres(&sqltools.Database{Connection: db})

	// Check result
	_, err = repo.CheckUserIsAuthor(ctx, &inputUser, &inputParams)
	require.NotNil(t, err, fmt.Errorf("unexpected err: %s", err))

	err = mock.ExpectationsWereMet()
	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
}
