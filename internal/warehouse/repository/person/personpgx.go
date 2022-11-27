package person

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
)

type Repository interface {
	GetPersonByID(ctx context.Context, person *models.Person, params *constparams.GetPersonParams) (models.Person, error)
}

// personPostgres is implementation repository of Postgres corresponding to the Repository interface.
type personPostgres struct {
	database *sqltools.Database
}

// NewPersonPostgres is constructor for personPostgres.
func NewPersonPostgres(database *sqltools.Database) Repository {
	return &personPostgres{
		database,
	}
}

func (p *personPostgres) GetPersonByID(ctx context.Context, person *models.Person, params *constparams.GetPersonParams) (models.Person, error) {
	response := NewPersonSQL()

	// Person - Main
	errMain := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowPerson := conn.QueryRowContext(ctx, getPersonByID, person.ID)
		if rowPerson.Err() != nil {
			return rowPerson.Err()
		}

		err := rowPerson.Scan(
			&response.Name,
			&response.Birthday,
			&response.Growth,
			&response.OriginalName,
			&response.Avatar,
			&response.Death,
			&response.Gender,
			&response.CountFilms)
		if err != nil {
			return err
		}

		if !response.Avatar.Valid {
			response.Avatar.String = constparams.DefPersonAvatar
		}

		return nil
	})
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrPersonNotFound,
			"Person main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonByID, person.ID, errMain)
	}

	if errMain != nil {
		return models.Person{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Person main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPersonByID, person.ID, errMain)
	}

	// Parts
	// Films + GenresFilms
	errQuery := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		var err error

		response.BestFilms, err = film.GetShortFilmsBatch(ctx, conn, getPersonBestFilms, person.ID, params.CountFilms)
		if err != nil && !stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				getPersonBestFilms, person.ID, params.CountFilms, err)
		}

		response.BestFilms, err = film.GetGenresBatch(ctx, response.BestFilms, conn)
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
		return models.Person{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Professions
	response.Professions, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonProfessions, person.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonGenres, person.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	return response.Convert(), nil
}
