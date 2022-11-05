package repository

import (
	"context"
	"database/sql"
	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"

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

	u.database.Connection.Conn()

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

		wg := &sync.WaitGroup{}

		// Films + GenresFilms
		wg.Add(1)
		go func() {
			defer wg.Done()

			params, _ := ctx.Value(innerPKG.GetReviewsParamsKey).(innerPKG.GetPersonParamsCtx)

			valuesFilms := []interface{}{person.ID, params.CountFilms}

			response.BestFilms, err = repository.GetShortFilmsBatch(ctx, tx, getPersonBestFilms, valuesFilms)
			if err != nil {
				logrus.Info("Films + GenresFilms", err)
				return
			}

			response.BestFilms, err = repository.GetGenresBatch(ctx, response.BestFilms, tx)
			if err != nil {
				logrus.Info("Films + GenresFilms", err)
				return
			}
			logrus.Info("Films + GenresFilms")
		}()

		//  Images
		wg.Add(1)
		go func() {
			defer wg.Done()

			rowPersonImages := tx.QueryRowContext(ctx, getPersonImages, person.ID)
			if rowPerson.Err() != nil {
				logrus.Info("Images")
				err = rowPerson.Err()
			}

			var images sql.NullString

			err = rowPersonImages.Scan(&images)
			if err != nil {
				logrus.Info("Images", err)
				return
			}

			response.Images = strings.Split(images.String, "_")
			logrus.Info("Images")
		}()

		//  Professions
		wg.Add(1)
		go func() {
			defer wg.Done()

			valuesProfessions := []interface{}{person.ID}

			response.Professions, err = sqltools.GetSimpleAttr(ctx, tx, getPersonProfessions, valuesProfessions)
			if err != nil {
				logrus.Info("Professions", err)
				return
			}
			logrus.Info("Professions")
		}()

		//  Genres
		wg.Add(1)
		go func() {
			defer wg.Done()

			valuesGenres := []interface{}{person.ID}

			response.Genres, err = sqltools.GetSimpleAttr(ctx, tx, getPersonGenres, valuesGenres)
			if err != nil {
				logrus.Info("Genres", err)
				return
			}
			logrus.Info("Genres")
		}()

		wg.Wait()
		logrus.Info("ALL")

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
