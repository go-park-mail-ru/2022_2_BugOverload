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

func (u personPostgres) GetPersonByID(ctx context.Context, person *models.Person) (models.Person, error) {
	response := NewPersonSQL()

	err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(tx *sql.Tx) error {
		// Person
		rowPerson := tx.QueryRowContext(ctx, getPerson, person.ID)
		if stdErrors.Is(rowPerson.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}

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

		//  Films
		params, _ := ctx.Value(innerPKG.GetReviewsParamsKey).(innerPKG.GetPersonParamsCtx)

		values := make([]interface{}, 0)
		values = append(values, person.ID, params.CountFilms)

		response.BestFilms, err = repository.GetShortFilmsBatch(ctx, tx, getPersonBestFilms, values)
		if err != nil {
			return err
		}

		wg := &sync.WaitGroup{}

		// GenresFilms
		wg.Add(1)
		go func() {
			defer wg.Done()

			response.BestFilms, err = repository.GetGenresBatch(ctx, response.BestFilms, tx)
			if err != nil {
				return
			}
		}()

		//  Images
		wg.Add(1)
		go func() {
			defer wg.Done()

			rowPersonImages := tx.QueryRowContext(ctx, getPersonImages, person.ID)
			if rowPerson.Err() != nil {
				err = rowPerson.Err()
			}

			var images sql.NullString

			err = rowPersonImages.Scan(&images)
			if err != nil {
				return
			}

			response.Images = strings.Split(images.String, "_")
		}()

		wg.Wait()

		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(err, errors.ErrNotFoundInDB) {
		return models.Person{}, errors.ErrNotFoundInDB
	}

	// the main entity is found, its components are not found
	if stdErrors.Is(err, sql.ErrNoRows) {
		return response.Convert(), nil
	}

	// execution error
	if err != nil {
		return models.Person{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}
