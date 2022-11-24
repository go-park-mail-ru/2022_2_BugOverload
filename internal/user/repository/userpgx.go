package repository

import (
	"context"
	"database/sql"

	stdErrors "github.com/pkg/errors"

	filmRepo "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type UserRepository interface {
	// Profile + settings
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)

	// FilmActivity
	GetUserActivityOnFilm(ctx context.Context, user *models.User, params *innerPKG.GetUserActivityOnFilmParams) (models.UserActivity, error)

	// ChangeInfo
	ChangeUserProfileNickname(ctx context.Context, user *models.User) error

	// Film
	FilmRatingExist(ctx context.Context, user *models.User, filmID int) (bool, error)
	FilmRateUpdate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error)
	FilmRateSet(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error)
	FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateDropParams) (models.Film, error)

	// Review
	NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *innerPKG.NewFilmReviewParams) error
}

// userPostgres is implementation repository of Postgres corresponding to the UserRepository interface.
type userPostgres struct {
	database *sqltools.Database
}

// NewUserPostgres is constructor for userPostgres.
func NewUserPostgres(database *sqltools.Database) UserRepository {
	return &userPostgres{
		database,
	}
}

func (u *userPostgres) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	response := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUserProfile := conn.QueryRowContext(ctx, getUserProfile, user.ID)
		if rowUserProfile.Err() != nil {
			return rowUserProfile.Err()
		}
		err := rowUserProfile.Scan(
			&response.Nickname,
			&response.JoinedDate,
			&response.Avatar,
			&response.CountViewsFilms,
			&response.CountCollections,
			&response.CountReviews,
			&response.CountRatings)
		if err != nil {
			return err
		}

		if !response.Avatar.Valid {
			response.Avatar.String = innerPKG.DefUserAvatar
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, stdErrors.WithMessagef(errors.ErrUserNotFound,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getUserProfile, user.ID, errMain)
	}

	// execution error
	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getUserProfile, user.ID, errMain)
	}

	return response.Convert(), nil
}

func (u *userPostgres) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	response := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowProfile := conn.QueryRowContext(ctx, getUserProfileShort, user.ID)
		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err := rowProfile.Scan(
			&response.JoinedDate,
			&response.CountViewsFilms,
			&response.CountCollections,
			&response.CountReviews,
			&response.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, stdErrors.WithMessagef(errors.ErrUserNotFound,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getUserProfileShort, user.ID, errMain)
	}

	// execution error
	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getUserProfileShort, user.ID, errMain)
	}

	return response.Convert(), nil
}

func (u *userPostgres) ChangeUserProfileNickname(ctx context.Context, user *models.User) error {
	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserSettingsNickname, user.Nickname, user.ID)
		if err != nil {
			return err
		}
		return nil
	})

	// execution error
	if errMain != nil {
		return stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%s, %d]. Special Error [%s]",
			updateUserSettingsNickname, user.Nickname, user.ID, errMain)
	}

	return nil
}

func (u *userPostgres) FilmRatingExist(ctx context.Context, user *models.User, filmID int) (bool, error) {
	response := false

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowExist := tx.QueryRowContext(ctx, checkUserRateExist, user.ID, filmID)
		if rowExist.Err() != nil {
			return rowExist.Err()
		}

		err := rowExist.Scan(&response)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return false, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
			checkUserRateExist, user.ID, filmID, errMain)
	}

	return response, nil
}

func (u *userPostgres) FilmRateUpdate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserRateFilm, user.ID, params.FilmID, params.Score)
		if err != nil {
			return err
		}

		rowFilmRating := tx.QueryRowContext(ctx, updateFilmRating, params.FilmID)
		if rowFilmRating.Err() != nil {
			return rowFilmRating.Err()
		}

		err = rowFilmRating.Scan(&resultFilm.Rating)
		if err != nil {
			return err
		}

		rowFilm := tx.QueryRowContext(ctx, getFilmRatingsCount, params.FilmID)
		if rowFilm.Err() != nil {
			return err
		}

		err = rowFilm.Scan(&resultFilm.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d, %d, %d]. Special Error [%s]",
			updateUserRateFilm, user.ID, params.FilmID, params.Score, errMain)
	}

	return resultFilm.Convert(), nil
}

func (u *userPostgres) FilmRateSet(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, setUserRateFilm, user.ID, params.FilmID, params.Score)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, updateAuthorCountRatingsUp, user.ID)
		if err != nil {
			return err
		}

		rowFilmRating := tx.QueryRowContext(ctx, updateFilmRating, params.FilmID)
		if rowFilmRating.Err() != nil {
			return rowFilmRating.Err()
		}

		err = rowFilmRating.Scan(&resultFilm.Rating)
		if err != nil {
			return err
		}

		rowFilm := tx.QueryRowContext(ctx, updateFilmCountRatingsUp, params.FilmID)
		if rowFilm.Err() != nil {
			return rowFilm.Err()
		}

		err = rowFilm.Scan(&resultFilm.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d, %d, %d]. Special Error [%s]",
			setUserRateFilm, user.ID, params.FilmID, params.Score, errMain)
	}

	return resultFilm.Convert(), nil
}

func (u *userPostgres) FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateDropParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, deleteUserRateFilm, user.ID, params.FilmID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, updateAuthorCountRatingsDown, user.ID)
		if err != nil {
			return err
		}

		rowFilmRating := tx.QueryRowContext(ctx, updateFilmRating, params.FilmID)
		if rowFilmRating.Err() != nil {
			return rowFilmRating.Err()
		}

		err = rowFilmRating.Scan(&resultFilm.Rating)
		if err != nil {
			return err
		}

		rowFilm := tx.QueryRowContext(ctx, updateFilmCountRatingsDown, params.FilmID)
		if rowFilm.Err() != nil {
			return err
		}

		err = rowFilm.Scan(&resultFilm.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Film{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
			deleteUserRateFilm, user.ID, params.FilmID, errMain)
	}

	return resultFilm.Convert(), nil
}

func (u *userPostgres) NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *innerPKG.NewFilmReviewParams) error {
	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, insertNewReview, review.Name, review.Type, review.Body)
		if row.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %+v]. Special Error [%s]",
				insertNewReview, user.ID, review, row.Err())
		}

		var reviewID int

		err := row.Scan(&reviewID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s], values - [%d, %+v]. Special Error [%s]",
				insertNewReview, user.ID, review, err)
		}

		_, err = tx.ExecContext(ctx, linkNewReviewAuthor, reviewID, user.ID, params.FilmID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s], values - [%d, %d, %d]. Special Error [%s]",
				linkNewReviewAuthor, reviewID, user.ID, params.FilmID, err)
		}

		_, err = tx.ExecContext(ctx, updateAuthorCountReviews, user.ID)
		if err != nil {
			return err
		}

		var targetCounterReviews string

		switch review.Type {
		case innerPKG.TypeReviewNegative:
			targetCounterReviews = updateFilmCountReviewNegative
		case innerPKG.TypeReviewNeutral:
			targetCounterReviews = updateFilmCountReviewNeutral
		default:
			targetCounterReviews = updateFilmCountReviewPositive
		}

		_, err = tx.ExecContext(ctx, targetCounterReviews, params.FilmID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s], values - [%d]. Special Error [%s]",
				targetCounterReviews, params.FilmID, err)
		}

		return nil
	})

	if errMain != nil {
		return errMain
	}

	return nil
}

func (u *userPostgres) GetUserActivityOnFilm(ctx context.Context, user *models.User, params *innerPKG.GetUserActivityOnFilmParams) (models.UserActivity, error) {
	response := NewUserActivitySQL()

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		// CountReviews
		rowUser := conn.QueryRowContext(ctx, getUserCountReviews, user.ID)
		if stdErrors.Is(rowUser.Err(), sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrUserNotFound,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				getUserCountReviews, user.ID, rowUser.Err())
		}

		if rowUser.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				getUserCountReviews, user.ID, rowUser.Err())
		}

		err := rowUser.Scan(&response.CountReviews)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s], values - [%d]. Special Error [%s]",
				getUserCountReviews, user.ID, rowUser.Err())
		}

		// UserRateFilm
		rowRating := conn.QueryRowContext(ctx, getUserRatingOnFilm, user.ID, params.FilmID)
		if rowRating.Err() != nil && !stdErrors.Is(rowRating.Err(), sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				getUserRatingOnFilm, user.ID, params.FilmID, err)
		}

		err = rowRating.Scan(&response.Rating, &response.DateRating)
		if err != nil && !stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				getUserRatingOnFilm, user.ID, params.FilmID, err)
		}

		// UserCollections
		rows, err := conn.QueryContext(ctx, getUserCollections, user.ID, params.FilmID)
		if err != nil && !stdErrors.Is(err, sql.ErrNoRows) {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				getUserCollections, user.ID, params.FilmID, err)
		}
		defer rows.Close()

		for rows.Next() {
			var collectionInfo NodeInUserCollectionSQL

			err = rows.Scan(&collectionInfo.ID, &collectionInfo.Name, &collectionInfo.IsUsed)
			if err != nil {
				return stdErrors.WithMessagef(errors.ErrWorkDatabase,
					"Err Scan: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
					getUserCollections, user.ID, params.FilmID, err)
			}

			response.ListCollections = append(response.ListCollections, collectionInfo)
		}

		return nil
	})

	// the main entity is not found
	if errMain != nil {
		return models.UserActivity{}, errMain
	}

	return response.Convert(), nil
}
