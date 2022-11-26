package collection

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"strconv"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

// Repository provides the versatility of collection repositories.
type Repository interface {
	GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetPremieresCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
}

// collectionPostgres is implementation repository of collection
// Postgres DB corresponding to the CollectionService interface.
type collectionPostgres struct {
	database *sqltools.Database
}

// NewCollectionPostgres is constructor for collectionPostgres.
func NewCollectionPostgres(database *sqltools.Database) Repository {
	return &collectionPostgres{
		database,
	}
}

// GetCollectionByTag it gives away movies by tag from the repository.
func (c *collectionPostgres) GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error
	var query string
	var values []interface{}

	switch params.SortParam {
	case constparams.CollectionSortParamDate:
		query = getFilmsByTagDate

		values = []interface{}{params.Key, params.CountFilms, params.Delimiter}
	case constparams.CollectionSortParamFilmRating:
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
		response.Films, err = film.GetShortFilmsBatch(ctx, conn, query, values...)
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
		response.Films, err = film.GetGenresBatch(ctx, response.Films, conn)
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

// GetCollectionByGenre it gives away movies by genre from the repository.
func (c *collectionPostgres) GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error
	var query string
	var values []interface{}

	switch params.SortParam {
	case constparams.CollectionSortParamDate:
		query = getFilmsByGenreDate

		values = []interface{}{params.Key, params.CountFilms, params.Delimiter}
	case constparams.CollectionSortParamFilmRating:
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
		response.Films, err = film.GetShortFilmsBatch(ctx, conn, query, values...)
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
		response.Films, err = film.GetGenresBatch(ctx, response.Films, conn)
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
func (c *collectionPostgres) GetPremieresCollection(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	response := NewCollectionSQL()

	var err error

	// Films - Main
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response.Films, err = film.GetNewFilmsBatch(ctx, conn, params.CountFilms, params.Delimiter)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: values - [%+v]. Special Error [%s]",
				params, err)
		}

		response.Films, err = film.GetGenresBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		response.Films, err = film.GetProdCountriesBatch(ctx, response.Films, conn)
		if err != nil {
			return err
		}

		response.Films, err = film.GetDirectorsBatch(ctx, response.Films, conn)
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
