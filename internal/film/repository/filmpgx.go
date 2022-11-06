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
	GetRecommendation(ctx context.Context) (models.Film, error)
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
		rowFilm := tx.QueryRowContext(ctx, getFilmByID, film.ID)
		if rowFilm.Err() != nil {
			return rowFilm.Err()
		}

		err := rowFilm.Scan(
			&response.Name,
			&response.OriginalName,
			&response.ProdYear,
			&response.Slogan,
			&response.Description,
			&response.ShortDescription,
			&response.AgeLimit,
			&response.Duration,
			&response.PosterHor,
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

		if !response.PosterHor.Valid {
			response.PosterHor.String = innerPKG.DefFilmPosterHor
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

			imagesSet := strings.Split(images.String, "_")

			params, _ := ctx.Value(innerPKG.GetFilmParamsKey).(innerPKG.GetFilmParamsCtx)

			if params.CountImages > len(imagesSet) {
				params.CountImages = len(imagesSet)
			}

			response.Images = imagesSet[:params.CountImages]

			return nil
		})
	}()

	// Actors
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			rowsFilmActors, err := tx.QueryContext(ctx, getFilmActors, film.ID)
			if err != nil {
				return err
			}
			defer rowsFilmActors.Close()

			for rowsFilmActors.Next() {
				var actor FilmActorSQL

				err = rowsFilmActors.Scan(
					&actor.ID,
					&actor.Name,
					&actor.Avatar,
					&actor.Character)
				if err != nil {
					return err
				}

				response.Actors = append(response.Actors, actor)
			}

			return nil
		})
	}()

	// Persons
	wg.Add(1)
	go func() {
		defer wg.Done()

		errTX = sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
			rowsFilmActors, err := tx.QueryContext(ctx, getFilmPersons, film.ID)
			if err != nil {
				return err
			}
			defer rowsFilmActors.Close()

			for rowsFilmActors.Next() {
				var person FilmPersonSQL
				var professionID int

				err = rowsFilmActors.Scan(
					&person.ID,
					&person.Name,
					&professionID)
				if err != nil {
					return err
				}

				switch professionID {
				case innerPKG.Artist:
					response.Artists = append(response.Artists, person)
				case innerPKG.Director:
					response.Directors = append(response.Directors, person)
				case innerPKG.Writer:
					response.Writers = append(response.Writers, person)
				case innerPKG.Producer:
					response.Producers = append(response.Producers, person)
				case innerPKG.Operator:
					response.Operators = append(response.Operators, person)
				case innerPKG.Montage:
					response.Montage = append(response.Montage, person)
				case innerPKG.Composer:
					response.Composers = append(response.Composers, person)
				}
			}

			return nil
		})
	}()

	wg.Wait()

	return response.Convert(), nil
}

func (f *filmPostgres) GetRecommendation(ctx context.Context) (models.Film, error) {
	response := NewFilmSQL()

	errTX := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowFilm := tx.QueryRowContext(ctx, getFilmRecommendation)
		if rowFilm.Err() != nil {
			return rowFilm.Err()
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
			values := []interface{}{response.ID}

			var err error

			response.Genres, err = sqltools.GetSimpleAttr(ctx, tx, getFilmGenres, values)
			if err != nil {
				return err
			}

			return nil
		})
	}()

	wg.Wait()

	return response.Convert(), nil
}
