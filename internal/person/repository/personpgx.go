package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"sync"

	stdErrors "github.com/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type PersonRepository interface {
	GetPersonByID(ctx context.Context, person *models.Person) (models.Person, error)
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

func (u personPostgres) GetPersonByID(ctx context.Context, person *models.Person) (models.Person, error) {
	response := NewPersonSQL()

	err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(tx *sql.Tx) error {
		// Person
		rowPerson := tx.QueryRowContext(ctx, getPerson, person.ID)
		if stdErrors.Is(rowPerson.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}

		if rowPerson.Err() != nil {
			return rowPerson.Err()
		}

		err := rowPerson.Scan(
			&response.Name,
			&response.Birthday,
			&response.Growth,
			&response.OriginalName,
			&response.Avatar,
			&response.Death,
			&response.Gender,
			&response.CountFilms)
		if err != nil {
			return err
		}

		//  Films
		countFilms := 5

		params, ok := ctx.Value(innerPKG.GetReviewsParamsKey).(innerPKG.GetPersonParamsCtx)
		if !ok {
			countFilms = params.CountFilms
		}

		rowsBestFilms, err := tx.QueryContext(ctx, getPersonBestFilms, person.ID, countFilms)
		if err != nil {
			return err
		}
		defer rowsBestFilms.Close()

		IDSet := make([]string, 0)

		for rowsBestFilms.Next() {
			film := repository.NewFilmSQL()

			err = rowsBestFilms.Scan(
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

			IDSet = append(IDSet, strconv.Itoa(film.ID))

			response.BestFilms = append(response.BestFilms, film)
		}

		IDSetResult := strings.Join(IDSet, ",")

		wg := &sync.WaitGroup{}

		// GenresFilms
		wg.Add(1)
		go func() {
			defer wg.Done()

			var rowsFilmsGenres *sql.Rows

			rowsFilmsGenres, err = tx.QueryContext(ctx, getGenresFilmBatchBegin+IDSetResult+getGenresFilmBatchEnd)
			if err != nil {
				return
			}

			defer rowsFilmsGenres.Close()

			counter := 0

			for rowsFilmsGenres.Next() {
				var filmID int
				var genre sql.NullString

				err = rowsFilmsGenres.Scan(
					&filmID,
					&genre)
				if err != nil {
					return
				}

				if filmID != response.BestFilms[counter].ID {
					counter++
				}

				response.BestFilms[counter].Genres = append(response.BestFilms[counter].Genres, genre.String)
			}
		}()

		//  Images
		wg.Add(1)
		go func() {
			defer wg.Done()

			rowPersonImages := tx.QueryRowContext(ctx, getPersonImages, person.ID)
			if rowPerson.Err() != nil {
				err = rowPerson.Err()
			}

			var images sql.NullString

			err = rowPersonImages.Scan(&images)
			if err != nil {
				return
			}

			response.Images = strings.Split(images.String, "_")
		}()

		wg.Wait()

		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(err, errors.ErrNotFoundInDB) {
		return models.Person{}, errors.ErrNotFoundInDB
	}

	// the main entity is found, its components are not found
	if stdErrors.Is(err, sql.ErrNoRows) {
		return response.Convert(), nil
	}

	// execution error
	if err != nil {
		return models.Person{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}
