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
	GetCollectionByTag(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByGenre(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error)
	GetUserCollections(ctx context.Context, user *models.User, params *innerPKG.GetUserCollectionsParams) ([]models.Collection, error)
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
func (c *collectionPostgres) GetCollectionByTag(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error
	var query string
	var values []interface{}

	switch params.SortParam {
	case innerPKG.CollectionSortParamDate:
		query = getFilmsByTagDate

		values = []interface{}{params.Key, params.CountFilms, params.Delimiter}
	case innerPKG.CollectionSortParamFilmRating:
		query = getFilmsByTagRating

		var delimiter float64

		delimiter, err = strconv.ParseFloat(params.Delimiter, 32)
		if err != nil {
			return models.Collection{}, stdErrors.WithMessagef(errors.ErrGetParamsConvert,
				"Get Delimeter Err: params input:[%s]",
				params.Delimiter)
		}

		values = []interface{}{params.Key, delimiter, params.CountFilms}
	default:
		return models.Collection{}, errors.ErrUnsupportedSortParameter
	}

	//  Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response.Films, err = repository.GetShortFilmsBatch(ctx, conn, query, values...)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}

		response.Name = params.Key

		//  Genres
		response.Films, err = repository.GetGenresBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Collection{}, errMain
	}

	return response.Convert(), nil
}

// GetUserCollections it gives away movies by genre from the repository.
func (c *collectionPostgres) GetUserCollections(ctx context.Context, user *models.User, params *innerPKG.GetUserCollectionsParams) ([]models.Collection, error) {
	response := NewCollectionSQL()

	var err error
	var query string

	if params.Delimiter == "now" {
		params.Delimiter = "NOW()"
	}

	values := []interface{}{user.ID, params.Delimiter, params.CountCollections}

	switch params.SortParam {
	case innerPKG.UserCollectionSortParamDate:
		query = getFilmsByGenreDate
	case innerPKG.UserCollectionSortParamFilmRating:
		query = getFilmsByGenreRating
	default:
		return []models.Collection{}, errors.ErrUnsupportedSortParameter
	}

	//  Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response.Films, err = repository.GetShortFilmsBatch(ctx, conn, query, values...)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}

		return nil
	})

	if errMain != nil {
		return []models.Collection{}, errMain
	}

	return response.Convert(), nil
}

// GetCollectionByGenre it gives away movies by genre from the repository.
func (c *collectionPostgres) GetCollectionByGenre(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error
	var query string
	var values []interface{}

	switch params.SortParam {
	case innerPKG.CollectionSortParamDate:
		query = getFilmsByGenreDate

		values = []interface{}{params.Key, params.CountFilms, params.Delimiter}
	case innerPKG.CollectionSortParamFilmRating:
		query = getFilmsByGenreRating

		var delimiter float64

		delimiter, err = strconv.ParseFloat(params.Delimiter, 32)
		if err != nil {
			return models.Collection{}, stdErrors.WithMessagef(errors.ErrGetParamsConvert,
				"Get Delimeter Err: params input:[%s]",
				params.Delimiter)
		}

		values = []interface{}{params.Key, delimiter, params.CountFilms}

	default:
		return models.Collection{}, errors.ErrUnsupportedSortParameter
	}

	//  Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response.Films, err = repository.GetShortFilmsBatch(ctx, conn, query, values...)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}

		response.Name = params.Key

		//  Genres
		response.Films, err = repository.GetGenresBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Collection{}, errMain
	}

	return response.Convert(), nil
}
