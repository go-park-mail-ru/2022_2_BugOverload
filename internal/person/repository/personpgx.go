package repository

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type PersonRepository interface {
	GetPersonByID(ctx context.Context, person *models.Person, params *innerPKG.GetPersonParams) (models.Person, error)
}

// personPostgres is implementation repository of Postgres corresponding to the PersonRepository interface.
type personPostgres struct {
	database *sqltools.Database
}

// NewPersonPostgres is constructor for personPostgres.
func NewPersonPostgres(database *sqltools.Database) PersonRepository {
	return &personPostgres{
		database,
	}
}

func (p *personPostgres) GetPersonByID(ctx context.Context, person *models.Person, params *innerPKG.GetPersonParams) (models.Person, error) {
	response := NewPersonSQL()

	// Person - Main
	errMain := response.GetMainInfo(ctx, p.database.Connection, getPersonByID, person.ID)
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Person main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonByID, person.ID, errMain)
	}

	if errMain != nil {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Person main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonByID, person.ID, errMain)
	}

	// Parts
	// Films + GenresFilms
	errQuery := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		var err error

		response.BestFilms, err = repository.GetShortFilmsBatch(ctx, conn, getPersonBestFilms, person.ID, params.CountFilms)
		if err != nil && !stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrPostgresRequest,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				getPersonBestFilms, person.ID, params.CountFilms, err)
		}

		response.BestFilms, err = repository.GetGenresBatch(ctx, response.BestFilms, conn)
		if err != nil {
			return err
		}

		return nil
	})
	if errQuery != nil {
		return models.Person{}, errQuery
	}

	//  Images
	response.Images, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonImages, person.ID, params.CountImages)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonImages, person.ID, errQuery)
	}

	// Professions
	response.Professions, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonProfessions, person.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Professions Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonProfessions, person.ID, errQuery)
	}

	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonGenres, person.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Genres Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonGenres, person.ID, errQuery)
	}

	return response.Convert(), nil
}
