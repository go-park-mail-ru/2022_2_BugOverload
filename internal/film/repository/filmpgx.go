package repository

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmRepository interface {
	GetRecommendation(ctx context.Context) (models.Film, error)
	GetFilmByID(ctx context.Context, film *models.Film, params *innerPKG.GetFilmParams) (models.Film, error)
}

// filmPostgres is implementation repository of Postgres corresponding to the FilmRepository interface.
type filmPostgres struct {
	database *sqltools.Database
}

// NewFilmPostgres is constructor for filmPostgres.
func NewFilmPostgres(database *sqltools.Database) FilmRepository {
	return &filmPostgres{
		database,
	}
}

func (f *filmPostgres) GetFilmByID(ctx context.Context, film *models.Film, params *innerPKG.GetFilmParams) (models.Film, error) {
	response := NewFilmSQL()

	// Film - Main
	errMain := response.GetMainInfo(ctx, f.database.Connection, getFilmByID, film.ID)
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Film main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmByID, film.ID, errMain)
	}

	if errMain != nil {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Film main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmByID, film.ID, errMain)
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Genres Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmGenres, film.ID, errQuery)
	}

	// Companies
	response.ProdCompanies, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCompanies, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Companies Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmCompanies, film.ID, errQuery)
	}

	// Countries
	response.ProdCountries, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCountries, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Countries Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmCountries, film.ID, errQuery)
	}

	// Tags
	response.Tags, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmTags, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Tags Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmTags, film.ID, errQuery)
	}

	//  Images
	errQuery = response.GetActors(ctx, f.database.Connection, getFilmImages, film.ID, params.CountImages)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmImages, film.ID, errQuery)
	}

	// Actors
	errQuery = response.GetActors(ctx, f.database.Connection, getFilmActors, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Actors Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmActors, film.ID, errQuery)
	}

	// Persons
	errQuery = response.GetPersons(ctx, f.database.Connection, getFilmPersons, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Persons Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmPersons, film.ID, errQuery)
	}

	return response.Convert(), nil
}

func (f *filmPostgres) GetRecommendation(ctx context.Context) (models.Film, error) {
	response := NewFilmSQL()

	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowFilm := conn.QueryRowContext(ctx, getFilmRecommendation)
		if stdErrors.Is(rowFilm.Err(), sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err: params input: query - [%s]. Special Error [%s]",
				getFilmRecommendation, rowFilm.Err())
		}
		if rowFilm.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s]. Special Error [%s]",
				getFilmRecommendation, rowFilm.Err())
		}

		err := rowFilm.Scan(
			&response.ID,
			&response.Name,
			&response.ProdYear,
			&response.EndYear,
			&response.PosterHor,
			&response.ShortDescription,
			&response.Rating)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Film{}, errMain
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, response.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Genres Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmGenres, response.ID, errQuery)
	}

	return response.Convert(), nil
}
