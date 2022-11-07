package repository

import (
	"context"
	"database/sql"
	"strconv"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetCollectionByTag(ctx context.Context, params *innerPKG.GetCollectionTagParams) (models.Collection, error)
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
func (c *collectionPostgres) GetCollectionByTag(ctx context.Context, params *innerPKG.GetCollectionTagParams) (models.Collection, error) {
	response := NewCollectionSQL()

	//  Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		delimiter, err := strconv.ParseFloat(params.Delimiter, 32)
		if err != nil {
			return errors.ErrGetParamsConvert
		}

		response.Films, err = repository.GetShortFilmsBatch(ctx, conn, getFilmsByTag, params.Tag, delimiter, params.CountFilms)
		if err != nil {
			return err
		}

		response.Name = params.Tag

		//  Genres
		response.Films, err = repository.GetGenresBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Collection{}, errors.ErrNotFoundInDB
	}

	// execution error
	if errMain != nil {
		return models.Collection{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}
