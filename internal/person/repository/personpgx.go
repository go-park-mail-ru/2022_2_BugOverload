package repository

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type PersonRepository interface {
	GetPersonByID(ctx context.Context, person *models.Person) (models.Person, error)
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

func (u *personPostgres) GetPersonByID(ctx context.Context, person *models.Person) (models.Person, error) {
	response := NewPersonSQL()

	// Person - Main
	errTX := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowPerson := tx.QueryRowContext(ctx, getPersonByID, person.ID)
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

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errTX, sql.ErrNoRows) {
		return models.Person{}, errors.ErrNotFoundInDB
	}

	if errTX != nil {
		return models.Person{}, errors.ErrPostgresRequest
	}

	wg := sync.WaitGroup{}

	// Parts
	// Films + GenresFilms
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			params, _ := ctx.Value(innerPKG.GetPersonParamsKey).(innerPKG.GetPersonParamsCtx)

			values := []interface{}{person.ID, params.CountFilms}

			var err error

			response.BestFilms, err = repository.GetShortFilmsBatch(ctx, tx, getPersonBestFilms, values)
			if err != nil {
				return err
			}

			response.BestFilms, err = repository.GetGenresBatch(ctx, response.BestFilms, tx)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	//  Images
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			rowPersonImages := tx.QueryRowContext(ctx, getPersonImages, person.ID)
			if rowPersonImages.Err() != nil {
				return rowPersonImages.Err()
			}

			var images sql.NullString

			err := rowPersonImages.Scan(&images)
			if err != nil {
				return err
			}

			response.Images = strings.Split(images.String, "_")

			return nil
		})
	}()

	//  Professions
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{person.ID}

			var err error

			response.Professions, err = sqltools.GetSimpleAttr(ctx, tx, getPersonProfessions, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	//  Genres
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{person.ID}

			var err error

			response.Genres, err = sqltools.GetSimpleAttr(ctx, tx, getPersonGenres, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	wg.Wait()

	return response.Convert(), nil
}
