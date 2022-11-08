package repository

import (
	"context"
	"database/sql"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type UserRepository interface {
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User) error

	// Support
	GetPassword(ctx context.Context, user *models.User) (string, error)
	CheckPassword(ctx context.Context, user *models.User) error
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
		rowUser := conn.QueryRowContext(ctx, getUser, user.ID)
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err := rowUser.Scan(&response.Nickname)
		if err != nil {
			return err
		}

		rowProfile := conn.QueryRowContext(ctx, getUserProfile, user.ID)
		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err = rowProfile.Scan(
			&response.Profile.JoinedDate,
			&response.Profile.Avatar,
			&response.Profile.CountViewsFilms,
			&response.Profile.CountCollections,
			&response.Profile.CountReviews,
			&response.Profile.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	// execution error
	if errMain != nil {
		return models.User{}, errors.ErrPostgresRequest
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
			&response.Profile.JoinedDate,
			&response.Profile.CountViewsFilms,
			&response.Profile.CountCollections,
			&response.Profile.CountReviews,
			&response.Profile.CountRatings)
		if err != nil {
			return err
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	// execution error
	if errMain != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}

func (u *userPostgres) ChangeUserProfileSettings(ctx context.Context, user *models.User) error {
	request := NewUserSQLOnUser(user)

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserSettings, request.Nickname, request.Password, request.ID)
		if err != nil {
			return err
		}

		return nil
	})

	// execution error
	if errMain != nil {
		return errors.ErrPostgresRequest
	}

	return nil
}

func (u *userPostgres) GetPassword(ctx context.Context, user *models.User) (string, error) {
	var res string

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowSalt := conn.QueryRowContext(ctx, getPass, user.ID)
		if rowSalt.Err() != nil {
			return rowSalt.Err()
		}

		err := rowSalt.Scan(&res)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return "", errors.ErrPostgresRequest
	}

	return res, nil
}

func (u *userPostgres) CheckPassword(ctx context.Context, user *models.User) error {
	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowCheck := conn.QueryRowContext(ctx, checkPassword, user.Password, user.ID)
		if rowCheck.Err() != nil {
			return rowCheck.Err()
		}

		return nil
	})

	// the main entity is not found
	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return errors.ErrWrongPassword
	}

	// execution error
	if errMain != nil {
		return errors.ErrPostgresRequest
	}

	return nil
}
