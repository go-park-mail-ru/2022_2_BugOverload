package repository

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmRepository interface {
	GetFilmByID(ctx context.Context, film *models.Film) (models.Film, error)
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

func (f *filmPostgres) GetFilmByID(ctx context.Context, film *models.Film) (models.Film, error) {
	response := NewFilmSQL()

	// Film - Main
	errTX := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowPerson := tx.QueryRowContext(ctx, getFilmByID, film.ID)
		if rowPerson.Err() != nil {
			return rowPerson.Err()
		}

		err := rowPerson.Scan(
			&response.Name,
			&response.OriginalName,
			&response.ProdYear,
			&response.Slogan,
			&response.Description,
			&response.AgeLimit,
			&response.Duration,
			&response.PosterVer,
			&response.Budget,
			&response.BoxOffice,
			&response.CurrencyBudget,
			&response.CountSeasons,
			&response.EndYear,
			&response.Type,
			&response.Rating,
			&response.CountActors,
			&response.CountScores,
			&response.CountNegativeReviews,
			&response.CountNeutralReviews,
			&response.CountPositiveReviews)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errTX, sql.ErrNoRows) {
		return models.Film{}, errors.ErrNotFoundInDB
	}

	if errTX != nil {
		return models.Film{}, errors.ErrPostgresRequest
	}

	wg := sync.WaitGroup{}

	// Parts
	// Genres
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{film.ID}

			var err error

			response.Genres, err = sqltools.GetSimpleAttr(ctx, tx, getFilmGenres, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	// Companies
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{film.ID}

			var err error

			response.ProdCompanies, err = sqltools.GetSimpleAttr(ctx, tx, getFilmCompanies, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	// Countries
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{film.ID}

			var err error

			response.ProdCountries, err = sqltools.GetSimpleAttr(ctx, tx, getFilmCountries, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	// Tags
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			values := []interface{}{film.ID}

			var err error

			response.Tags, err = sqltools.GetSimpleAttr(ctx, tx, getFilmTags, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	// Images
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			rowFilmImages := tx.QueryRowContext(ctx, getFilmImages, film.ID)
			if rowFilmImages.Err() != nil {
				return rowFilmImages.Err()
			}

			var images sql.NullString

			err := rowFilmImages.Scan(&images)
			if err != nil {
				return err
			}

			response.Images = strings.Split(images.String, "_")

			return nil
		})
	}()

	wg.Wait()

	return response.Convert(), nil
}
