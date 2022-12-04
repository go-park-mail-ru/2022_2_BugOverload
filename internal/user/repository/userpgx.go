package repository

import (
	"context"
	"database/sql"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
	filmRepo "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
)

//go:generate mockgen -source userpgx.go -destination mocks/mockuserrepository.go -package mockUserRepository
type UserRepository interface {
	// Profile + settings
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)

	// FilmActivity
	GetUserActivityOnFilm(ctx context.Context, user *models.User, params *constparams.GetUserActivityOnFilmParams) (models.UserActivity, error)

	// ChangeInfo
	ChangeUserProfileNickname(ctx context.Context, user *models.User) error

	// Film
	FilmRatingExist(ctx context.Context, user *models.User, filmID int) (bool, error)
	FilmRateUpdate(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error)
	FilmRateSet(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error)
	FilmRateDrop(ctx context.Context, user *models.User, params *constparams.FilmRateDropParams) (models.Film, error)

	// Review
	NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *constparams.NewFilmReviewParams) error

	// Personal collections
	GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error)

	// Update personal collections
	CheckUserAccessToUpdateCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error)
	CheckExistFilmInCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error)
	AddFilmToCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error
	DropFilmFromCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error
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
			response.Avatar.String = constparams.DefUserAvatar
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
	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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

	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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

func (u *userPostgres) FilmRateUpdate(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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
			updateUserRateFilm, user.ID, params.FilmID, params.Score, errMain)
	}

	return resultFilm.Convert(), nil
}

func (u *userPostgres) FilmRateSet(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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

func (u *userPostgres) FilmRateDrop(ctx context.Context, user *models.User, params *constparams.FilmRateDropParams) (models.Film, error) {
	resultFilm := filmRepo.NewFilmSQL()

	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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
			"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
			deleteUserRateFilm, user.ID, params.FilmID, errMain)
	}

	return resultFilm.Convert(), nil
}

func (u *userPostgres) NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *constparams.NewFilmReviewParams) error {
	errMain := sqltools.RunTxOnConn(ctx, constparams.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
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
		case constparams.TypeReviewNegative:
			targetCounterReviews = updateFilmCountReviewNegative
		case constparams.TypeReviewNeutral:
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

func (u *userPostgres) GetUserActivityOnFilm(ctx context.Context, user *models.User, params *constparams.GetUserActivityOnFilmParams) (models.UserActivity, error) {
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

// GetUserCollections it gives away movies by genre from the repository.
func (u *userPostgres) GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error) {
	response := make([]collection.ModelSQL, 0)

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
	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
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
			collection := collection.NewCollectionSQL()

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

// CheckUserAccessToUpdateCollection check that user has author's access of params collection
func (u *userPostgres) CheckUserAccessToUpdateCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error) {
	var response bool

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkUserAccessToUpdateCollection, user.ID, params.CollectionID)
		if row.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				checkUserAccessToUpdateCollection, user.ID, params.CollectionID, row.Err())
		}

		err := row.Scan(&response)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				checkUserAccessToUpdateCollection, user.ID, params.CollectionID, err)
		}

		return nil
	})

	if errMain != nil {
		return false, errMain
	}

	return response, nil
}

// CheckExistFilmInCollection check that current film exist in collection
func (u *userPostgres) CheckExistFilmInCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error) {
	var response bool

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkFilmExistInCollection, params.CollectionID, params.FilmID)
		if row.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				checkFilmExistInCollection, user.ID, params.CollectionID, row.Err())
		}

		err := row.Scan(&response)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				checkFilmExistInCollection, user.ID, params.CollectionID, err)
		}

		return nil
	})

	if errMain != nil {
		return false, errMain
	}

	return response, nil
}

// AddFilmToCollection return nil if film added successfully, error otherwise
func (u *userPostgres) AddFilmToCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error {
	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		var filmExist bool
		var err error

		row := conn.QueryRowContext(ctx, checkFilmExist, params.FilmID)
		if row.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], value - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, row.Err())
		}
		err = row.Scan(&filmExist)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], value - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, err)
		}
		if !filmExist {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, err)
		}

		_, err = conn.ExecContext(ctx, addFilmToCollection, params.CollectionID, params.FilmID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				addFilmToCollection, params.FilmID, params.CollectionID, err)
		}

		_, err = conn.ExecContext(ctx, updateCollectionFilmsCountUp, params.CollectionID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				updateCollectionFilmsCountUp, params.CollectionID, err)
		}

		return nil
	})
	return errMain
}

// DropFilmFromCollection return nil if film removed successfully, error otherwise
func (u *userPostgres) DropFilmFromCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error {
	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		var filmExist bool
		var err error

		row := conn.QueryRowContext(ctx, checkFilmExist, params.FilmID)
		if row.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], value - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, row.Err())
		}
		err = row.Scan(&filmExist)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], value - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, err)
		}
		if !filmExist {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				checkFilmExist, params.FilmID, err)
		}

		_, err = conn.ExecContext(ctx, dropFilmFromCollection, params.CollectionID, params.FilmID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %d]. Special Error [%s]",
				addFilmToCollection, params.FilmID, params.CollectionID, err)
		}

		_, err = conn.ExecContext(ctx, updateCollectionFilmsCountDown, params.CollectionID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
				updateCollectionFilmsCountDown, params.CollectionID, err)
		}

		return nil
	})
	return errMain
}
