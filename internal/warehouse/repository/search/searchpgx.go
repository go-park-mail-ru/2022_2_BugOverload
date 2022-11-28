package search

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/person"
)

type Repository interface {
	SearchFilms(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error)
	SearchSeries(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error)
	SearchPersons(ctx context.Context, params *constparams.SearchParams) ([]models.Person, error)
}

// searchPostgres is implementation repository of Postgres corresponding to the Repository interface.
type searchPostgres struct {
	database *sqltools.Database
}

// NewSearchPostgres is constructor for searchPostgres.
func NewSearchPostgres(database *sqltools.Database) Repository {
	return &searchPostgres{
		database,
	}
}

func (s *searchPostgres) SearchFilms(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error) {
	var responseDB []film.ModelSQL
	var err error

	errMain := sqltools.RunQuery(ctx, s.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		responseDB, err = film.GetShortFilmsBatch(ctx, conn, searchFilmsByName, params.Query)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: values - [%+v]. Special Error [%s]",
				params, err)
		}

		responseDB, err = film.GetGenresBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		responseDB, err = film.GetProdCountriesBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		responseDB, err = film.GetDirectorsBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return []models.Film{}, nil
	}
	if errMain != nil {
		return []models.Film{}, errMain
	}

	response := make([]models.Film, len(responseDB))
	for idx := range responseDB {
		response[idx] = responseDB[idx].Convert()
	}
	return response, nil
}

func (s *searchPostgres) SearchSeries(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error) {
	var responseDB []film.ModelSQL
	var err error

	errMain := sqltools.RunQuery(ctx, s.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		responseDB, err = film.GetShortFilmsBatch(ctx, conn, searchSeriesByName, params.Query)
		if stdErrors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info Err: params input: values - [%+v]. Special Error [%s]",
				params, err)
		}

		responseDB, err = film.GetGenresBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		responseDB, err = film.GetProdCountriesBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		responseDB, err = film.GetDirectorsBatch(ctx, responseDB, conn)
		if err != nil {
			return err
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return []models.Film{}, nil
	}
	if errMain != nil {
		return []models.Film{}, errMain
	}

	response := make([]models.Film, len(responseDB))
	for idx := range responseDB {
		response[idx] = responseDB[idx].Convert()
	}
	return response, nil
}

func (s *searchPostgres) SearchPersons(ctx context.Context, params *constparams.SearchParams) ([]models.Person, error) {
	var responseDB []person.ModelSQL

	errMain := sqltools.RunQuery(ctx, s.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsPersons, err := conn.QueryContext(ctx, searchPersonsByName, params.Query)
		if err != nil {
			logrus.Info("NeededCondition ", err)
			return err
		}

		for rowsPersons.Next() {
			var currentPerson person.ModelSQL
			err = rowsPersons.Scan(
				&currentPerson.ID,
				&currentPerson.Name,
				&currentPerson.Birthday,
				&currentPerson.OriginalName,
				&currentPerson.Avatar,
				&currentPerson.CountFilms)

			if err != nil {
				return errors.ErrWorkDatabase
			}

			if !currentPerson.Avatar.Valid {
				currentPerson.Avatar.String = constparams.DefPersonAvatar
			}

			// Professions
			var errQuery error
			currentPerson.Professions, errQuery = sqltools.GetSimpleAttrOnConn(ctx, s.database.Connection, getPersonProfessions, currentPerson.ID)
			if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
				return errors.ErrWorkDatabase
			}

			responseDB = append(responseDB, currentPerson)
		}

		//  Это какой то треш, запрос не отдает sql.ErrNoRows
		if len(responseDB) == 0 {
			logrus.Info("BadCondition")
			return sql.ErrNoRows
		}

		rowsPersons.Close()

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return []models.Person{}, nil
	}
	if errMain != nil {
		return []models.Person{}, errMain
	}

	response := make([]models.Person, len(responseDB))
	for idx := range responseDB {
		response[idx] = responseDB[idx].Convert()
	}
	return response, nil
}
