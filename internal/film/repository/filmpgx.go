package repository

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type FilmRepository interface {
	GetRecommendation(ctx context.Context) (models.Film, error)
	GetFilmByID(ctx context.Context, film *models.Film, params *innerPKG.GetFilmParams) (models.Film, error)
	GetReviewsByFilmID(ctx context.Context, params *innerPKG.GetReviewsFilmParams) ([]models.Review, error)
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
	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowFilm := conn.QueryRowContext(ctx, getFilmByID, film.ID)
		if rowFilm.Err() != nil {
			return rowFilm.Err()
		}

		err := rowFilm.Scan(
			&response.Name,
			&response.OriginalName,
			&response.ProdDate,
			&response.Slogan,
			&response.Description,
			&response.ShortDescription,
			&response.AgeLimit,
			&response.Duration,
			&response.PosterHor,
			&response.Budget,
			&response.BoxOffice,
			&response.CurrencyBudget,
			&response.Type,
			&response.Rating,
			&response.CountActors,
			&response.CountRatings,
			&response.CountNegativeReviews,
			&response.CountNeutralReviews,
			&response.CountPositiveReviews)
		if err != nil {
			return err
		}

		if !response.PosterHor.Valid {
			response.PosterHor.String = innerPKG.DefFilmPosterHor
		}

		if response.Type.String != innerPKG.DefTypeSerial {
			return nil
		}

		rowSerial := conn.QueryRowContext(ctx, getSerialByID, film.ID)
		if rowSerial.Err() != nil {
			return rowSerial.Err()
		}

		err = rowSerial.Scan(
			&response.CountSeasons,
			&response.EndYear)
		if err != nil {
			return err
		}

		return nil
	})
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrFilmNotFound,
			"Film main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmByID, film.ID, errMain)
	}

	if errMain != nil {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Film main info Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmByID, film.ID, errMain)
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, film.ID)
	if errQuery != nil && !stdErrors.Is(stdErrors.Cause(errQuery), sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Companies
	response.ProdCompanies, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCompanies, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Countries
	response.ProdCountries, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmCountries, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Tags
	response.Tags, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmTags, film.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	//  Images
	response.Images, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmImages, film.ID, params.CountImages)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	// Actors
	errQuery = sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsFilmActors, err := conn.QueryContext(ctx, getFilmActors, film.ID)
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
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Actors Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmActors, film.ID, errQuery)
	}

	// Persons
	errQuery = sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsFilmPersons, err := conn.QueryContext(ctx, getFilmPersons, film.ID)
		if err != nil {
			return err
		}
		defer rowsFilmPersons.Close()

		for rowsFilmPersons.Next() {
			var person FilmPersonSQL
			var professionID int

			err = rowsFilmPersons.Scan(
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
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Persons Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getFilmPersons, film.ID, errQuery)
	}

	return response.Convert(), nil
}

func (f *filmPostgres) GetRecommendation(ctx context.Context) (models.Film, error) {
	response := NewFilmSQL()

	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowFilm := conn.QueryRowContext(ctx, getFilmRecommendation)
		if stdErrors.Is(rowFilm.Err(), sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err: params input: query - [%s]. Special Error [%s]",
				getFilmRecommendation, rowFilm.Err())
		}
		if rowFilm.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s]. Special Error [%s]",
				getFilmRecommendation, rowFilm.Err())
		}

		err := rowFilm.Scan(
			&response.ID,
			&response.Name,
			&response.ProdDate,
			&response.Type,
			&response.PosterHor,
			&response.ShortDescription,
			&response.Rating)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s]. Special Error [%s]",
				getFilmRecommendation, err)
		}

		if response.Type.String != innerPKG.DefTypeSerial {
			return nil
		}

		rowSerial := conn.QueryRowContext(ctx, getShortSerialByID, response.ID)
		if rowSerial.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				getShortSerialByID, response.ID, rowSerial.Err())
		}

		err = rowSerial.Scan(&response.EndYear)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err Scan: params input: query - [%s], values - [%d]. Special Error [%s]",
				getShortSerialByID, response.ID, rowSerial.Err())
		}

		return nil
	})

	if errMain != nil {
		return models.Film{}, errMain
	}

	var errQuery error

	// Parts
	// Genres
	response.Genres, errQuery = sqltools.GetSimpleAttrOnConn(ctx, f.database.Connection, getFilmGenres, response.ID)
	if errQuery != nil && !stdErrors.Is(errQuery, sql.ErrNoRows) {
		return models.Film{}, stdErrors.WithMessage(errors.ErrWorkDatabase, errQuery.Error())
	}

	return response.Convert(), nil
}

func (f *filmPostgres) GetReviewsByFilmID(ctx context.Context, params *innerPKG.GetReviewsFilmParams) ([]models.Review, error) {
	response := make([]ReviewSQL, 0)

	// Reviews - Main
	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsReviews, err := conn.QueryContext(ctx, getReviewsByFilmID, params.FilmID, params.CountReviews, params.Offset)
		if err != nil {
			return err
		}
		defer rowsReviews.Close()

		for rowsReviews.Next() {
			var review ReviewSQL

			err = rowsReviews.Scan(
				&review.Name,
				&review.Type,
				&review.Body,
				&review.CountLikes,
				&review.CreateTime,
				&review.Author.ID,
				&review.Author.Nickname,
				&review.Author.Avatar,
				&review.Author.CountReviews)
			if err != nil {
				return err
			}

			if !review.Author.Avatar.Valid {
				review.Author.Avatar.String = innerPKG.DefPersonAvatar
			}

			response = append(response, review)
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) || len(response) == 0 {
		return []models.Review{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Err: params input: query - [%s], valies - [%d, %d, %d]. Special Error [%s]",
			getReviewsByFilmID, params.FilmID, params.CountReviews, params.Offset, sql.ErrNoRows)
	}

	if errMain != nil {
		return []models.Review{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], valies - [%d, %d, %d]. Special Error [%s]",
			getReviewsByFilmID, params.FilmID, params.CountReviews, params.Offset, errMain)
	}

	res := make([]models.Review, len(response))

	for idx, value := range response {
		res[idx] = value.Convert()
	}

	return res, nil
}
