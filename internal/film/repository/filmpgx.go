package repository

import (
	"context"
	"database/sql"
	stdErrors "github.com/pkg/errors"
	"strings"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmRepository interface {
	GetRecommendation(ctx context.Context) (models.Film, error)
	GetFilmByID(ctx context.Context, film *models.Film, params *innerPKG.GetFilmParams) (models.Film, error)
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

func (f *filmPostgres) GetFilmByID(ctx context.Context, film *models.Film, params *innerPKG.GetFilmParams) (models.Film, error) {
	response := NewFilmSQL()

	// Film - Main
	errMain := response.GetMainInfo(ctx, f.database.Connection, getFilmByID, film.ID)
	if errMain != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "GetMainInfo")
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "Genres")
	}

	// Companies
	response.ProdCompanies, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCompanies, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "Companies")
	}

	// Countries
	response.ProdCountries, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCountries, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "Countries")
	}

	// Tags
	response.Tags, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmTags, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "Tags")
	}

	//  Images
	errQuery = sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowPersonImages := conn.QueryRowContext(ctx, getFilmImages, film.ID)
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

	// Actors
	errQuery = response.GetActors(ctx, f.database.Connection, getFilmActors, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "GetActors")
	}

	// Persons
	errQuery = response.GetPersons(ctx, f.database.Connection, getFilmPersons, film.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "GetPersons")
	}

	return response.Convert(), nil
}

func (f *filmPostgres) GetRecommendation(ctx context.Context) (models.Film, error) {
	response := NewFilmSQL()

	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowFilm := conn.QueryRowContext(ctx, getFilmRecommendation)
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
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Film{}, errors.ErrNotFoundInDB
	}

	if errMain != nil {
		return models.Film{}, errors.ErrPostgresRequest
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, response.ID)
	if errQuery != nil {
		return models.Film{}, stdErrors.Wrap(errMain, "Genres")
	}

	return response.Convert(), nil
}
