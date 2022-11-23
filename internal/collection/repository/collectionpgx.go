package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

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
	GetPremieresCollection(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error)
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
			return stdErrors.WithMessagef(errors.ErrFilmsNotFound,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}

		// Tag Description
		rowTag := conn.QueryRowContext(ctx, getTagDescription, params.Key)
		if rowTag.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Get Tag desctiption Err: params input: query - [%s], values - [%s]. Special Error [%s]",
				getTagDescription, params.Key, err)
		}

		err = rowTag.Scan(&response.Description)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Get Tag desctiption Scan Err: params input: query - [%s], values - [%s]. Special Error [%s]",
				getTagDescription, params.Key, err)
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
	response := make([]CollectionSQL, 0)

	var query string

	if params.Delimiter == innerPKG.UserCollectionsDelimiter {
		params.Delimiter = time.Now().Format(innerPKG.DateFormat + " " + innerPKG.TimeFormat)
	}

	values := []interface{}{user.ID, params.Delimiter, params.CountCollections}

	switch params.SortParam {
	case innerPKG.UserCollectionsSortParamCreateDate:
		query = getUserCollectionByCreateDate
	case innerPKG.UserCollectionsSortParamUpdateDate:
		query = getUserCollectionByUpdateDate
	default:
		return []models.Collection{}, errors.ErrUnsupportedSortParameter
	}

	//  Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsCollections, err := conn.QueryContext(ctx, query, values...)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrCollectionsNotFound,
				"Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		defer rowsCollections.Close()

		for rowsCollections.Next() {
			collection := NewCollectionSQL()

			err = rowsCollections.Scan(
				&collection.ID,
				&collection.Name,
				&collection.Poster,
				&collection.CountFilms,
				&collection.CountLikes,
				&collection.UpdateTime,
				&collection.CreateTime)
			if err != nil {
				return stdErrors.WithMessagef(errors.ErrWorkDatabase,
					"Err Scan: params input: query - [%s], values - [%+v]. Special Error [%s]",
					query, values, err)
			}

			response = append(response, collection)
		}

		if len(response) == 0 {
			return stdErrors.WithMessagef(errors.ErrFilmsNotFound,
				"Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}

		return nil
	})

	if errMain != nil {
		return []models.Collection{}, errMain
	}

	res := make([]models.Collection, len(response))

	for idx, value := range response {
		res[idx] = value.Convert()
	}

	return res, nil
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
			return stdErrors.WithMessagef(errors.ErrFilmsNotFound,
				"Film main info Err: params input: query - [%s], values - [%+v]. Special Error [%s]",
				query, values, err)
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
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

// GetPremieresCollection it gives away only movies with prod_date > current from the repository.
func (c *collectionPostgres) GetPremieresCollection(ctx context.Context, params *innerPKG.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error

	// Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response.Films, err = repository.GetNewFilmsBatch(ctx, conn, params.CountFilms, params.Delimiter)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: values - [%+v]. Special Error [%s]",
				params, err)
		}

		response.Films, err = repository.GetGenresBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		response.Films, err = repository.GetProdCountriesBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		response.Films, err = repository.GetDirectorsBatch(ctx, response.Films, conn)
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
