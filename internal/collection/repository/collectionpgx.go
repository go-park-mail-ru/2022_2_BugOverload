package repository

import (
	"context"
	"database/sql"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"strconv"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

// CollectionRepository provides the versatility of collection repositories.
type CollectionRepository interface {
	GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error)
	GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error)
	GetPremieresCollection(ctx context.Context, params *constparams.PremiersCollectionParams) (models.Collection, error)

	CheckExistFilmInCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) (bool, error)
	AddFilmToCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) error
	DropFilmFromCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) error
}

// collectionPostgres is implementation repository of collection
// Postgres DB corresponding to the CollectionService interface.
type collectionPostgres struct {
	database *sqltools.Database
}

// NewCollectionPostgres is constructor for collectionPostgres.
func NewCollectionPostgres(database *sqltools.Database) CollectionRepository {
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
func (c *collectionPostgres) GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error) {
	response := make([]CollectionSQL, 0)

	var query string

	if params.Delimiter == constparams.UserCollectionsDelimiter {
		params.Delimiter = time.Now().Format(constparams.DateFormat + " " + constparams.TimeFormat)
	}

	values := []interface{}{user.ID, params.Delimiter, params.CountCollections}

	switch params.SortParam {
	case constparams.UserCollectionsSortParamCreateDate:
		query = getUserCollectionByCreateDate
	case constparams.UserCollectionsSortParamUpdateDate:
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
func (c *collectionPostgres) GetPremieresCollection(ctx context.Context, params *constparams.PremiersCollectionParams) (models.Collection, error) {
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

// CheckExistFilmInCollection return true if film exist in collection, false otherwise
func (c *collectionPostgres) CheckExistFilmInCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) (bool, error) {
	var response bool

	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkFilmExistInCollection, params.CollectionID, params.FilmID)
		if row.Err() != nil {
			return row.Err()
		}

		err := row.Scan(&response)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return false, errMain
	}

	return response, nil
}

// AddFilmToCollection return nil if film added successfully, error otherwise
func (c *collectionPostgres) AddFilmToCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) error {
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, addFilmToCollection, params.CollectionID, params.FilmID)
		if err != nil {
			return err
		}

		return nil
	})
	return errMain
}

// DropFilmFromCollection return nil if film removed successfully, error otherwise
func (c *collectionPostgres) DropFilmFromCollection(ctx context.Context, params *constparams.CollectionFilmsUpdateParams) error {
	errMain := sqltools.RunQuery(ctx, c.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		_, err := conn.ExecContext(ctx, dropFilmFromCollection, params.CollectionID, params.FilmID)
		if err != nil {
			return err
		}

		return nil
	})
	return errMain
}
