package repository

import (
	"context"
	"database/sql"
	"strings"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type PersonRepository interface {
	GetPersonByID(ctx context.Context, person *models.Person, params *innerPKG.GetPersonParams) (models.Person, error)
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

func (p *personPostgres) GetPersonByID(ctx context.Context, person *models.Person, params *innerPKG.GetPersonParams) (models.Person, error) {
	response := NewPersonSQL()

	// Person - Main
	errMain := response.GetMainInfo(ctx, p.database.Connection, getPersonByID, person.ID)
	if errMain != nil {
		return models.Person{}, stdErrors.Wrap(errMain, "GetMainInfo")
	}

	// Parts
	// Films + GenresFilms
	errQuery := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		var err error

		response.BestFilms, err = repository.GetShortFilmsBatch(ctx, conn, getPersonBestFilms, person.ID, params.CountFilms)
		if err != nil {
			return err
		}

		response.BestFilms, err = repository.GetGenresBatch(ctx, response.BestFilms, conn)
		if err != nil {
			return err
		}

		return nil
	})
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, errors.ErrPostgresRequest
	}

	//  Images
	errQuery = sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowPersonImages := conn.QueryRowContext(ctx, getPersonImages, person.ID)
		if rowPersonImages.Err() != nil {
			return rowPersonImages.Err()
		}

		var images sql.NullString

		err := rowPersonImages.Scan(&images)
		if err != nil {
			return err
		}

		response.Images = strings.Split(images.String, "_")

		imagesSet := strings.Split(images.String, "_")

		if params.CountImages > len(imagesSet) {
			params.CountImages = len(imagesSet)
		}

		response.Images = imagesSet[:params.CountImages]

		return nil
	})
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Person{}, errors.ErrPostgresRequest
	}

	// Professions
	response.Professions, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonProfessions, person.ID)
	if errQuery != nil {
		return models.Person{}, stdErrors.Wrap(errMain, "Professions")
	}

	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, p.database.Connection, getPersonGenres, person.ID)
	if errQuery != nil {
		return models.Person{}, stdErrors.Wrap(errMain, "Genres")
	}

	return response.Convert(), nil
}
