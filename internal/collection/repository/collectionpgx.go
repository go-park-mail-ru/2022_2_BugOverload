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
	"strconv"
)

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetCollectionByTag(ctx context.Context) (models.Collection, error)
}

// collectionPostgres is implementation repository of collection
// Postgres DB corresponding to the CollectionService interface.
type collectionPostgres struct {
	database *sqltools.Database
}

// NewCollectionCache is constructor for collectionPostgres.
func NewCollectionCache(database *sqltools.Database) CollectionRepository {
	return &collectionPostgres{
		database,
	}
}

// GetCollectionByTag it gives away movies by tag from the repository.
func (c *collectionPostgres) GetCollectionByTag(ctx context.Context) (models.Collection, error) {
	response := NewCollectionSQL()

	err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, c.database.Connection, func(tx *sql.Tx) error {
		params, _ := ctx.Value(innerPKG.GetCollectionTagParamsKey).(innerPKG.GetCollectionTagParamsCtx)

		delimiter, err := strconv.Atoi(params.Delimiter)
		if err != nil {
			return errors.ErrGetParamsConvert
		}

		rowsFilms, err := tx.QueryContext(ctx, getFilmsByTag, params.Tag, delimiter, params.Count)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}

		if err != nil {
			return err
		}

		for rowsFilms.Next() {
			film := repository.NewFilmSQL()

			err = rowsFilms.Scan(
				&film.ID,
				&film.Name,
				&film.OriginalName,
				&film.ProdYear,
				&film.PosterVer,
				&film.EndYear,
				&film.Rating)
			if err != nil {
				return err
			}

			response.Films = append(response.Films, film)
		}

		//  Genres
		response.Films, err = repository.GetGenresBatch(ctx, response.Films, tx)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(err, errors.ErrNotFoundInDB) {
		return models.Collection{}, errors.ErrNotFoundInDB
	}

	// execution error
	if err != nil {
		return models.Collection{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}
