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

// AuthRepository provides the versatility of users repositories.
type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, userID int) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
}

// AuthPostgres is implementation repository of users to the AuthRepository interface.
type AuthPostgres struct {
	database *sqltools.Database
}

// NewAuthDatabase is constructor for AuthPostgres. Accepts only sqltools.Database.
func NewAuthDatabase(database *sqltools.Database) AuthRepository {
	return &AuthPostgres{
		database,
	}
}

// CheckExist is a check for the existence of such a user by email.
func (ad *AuthPostgres) CheckExist(ctx context.Context, email string) (bool, error) {
	response := false
	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkExist, email)
		if stdErrors.Is(row.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}
		if row.Err() != nil {
			return row.Err()
		}

		err := row.Scan(&response)
		if err != nil {
			return err
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return false, errors.ErrNotFoundInDB
	}

	if errMain != nil {
		return false, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], email - [%s]. Special error [%s]",
			checkExist, email, errMain)
	}

	return response, nil
}

// CreateUser is creates a new user and set default avatar.
func (ad *AuthPostgres) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	exist, err := ad.CheckExist(ctx, user.Email)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "CreateUser falls on check exist")
	}
	if exist {
		return models.User{}, errors.ErrSignupUserExist
	}

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, ad.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, createUser, user.Email, user.Nickname, []byte(user.Password))
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err = rowUser.Scan(&user.ID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, createUserProfile, user.ID)
		if err != nil {
			return err
		}

		rowsCollections, errCollections := tx.QueryContext(ctx, createDefCollections)
		if errCollections != nil {
			return errCollections
		}
		defer rowsCollections.Close()

		ids := make([]int, 0)

		for rowsCollections.Next() {
			var colID int
			err = rowsCollections.Scan(&colID)
			if err != nil {
				return err
			}

			ids = append(ids, colID)
		}
		err = rowsCollections.Close()
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, linkUserProfileDefCollections, user.ID, ids[0], ids[1])
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: user - [%v]. Special error: [%s]",
			user, err)
	}

	user.Profile = models.Profile{
		Avatar: innerPKG.DefUserAvatar,
	}

	return *user, nil
}

// GetUserByEmail is returns all user attributes by part user attributes.
func (ad *AuthPostgres) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	exist, err := ad.CheckExist(ctx, email)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "GetUserByEmail falls on check exist")
	}
	if !exist {
		return models.User{}, errors.ErrUserNotExist
	}

	userDB := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByEmail, email)
		if stdErrors.Is(rowUser.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err = rowUser.Scan(
			&userDB.ID,
			&userDB.email,
			&userDB.nickname,
			&userDB.password)
		if err != nil {
			return err
		}

		rowProfile := conn.QueryRowContext(ctx, getProfileAvatar, userDB.ID)
		if stdErrors.Is(rowProfile.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}
		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err = rowProfile.Scan(&userDB.avatar)
		if err != nil {
			return err
		}

		if !userDB.avatar.Valid {
			userDB.avatar.String = innerPKG.DefPersonAvatar
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: email - [%s]. Special error - [%s]",
			email, errMain)
	}

	return userDB.Convert(), nil
}

// GetUserByID is returns all user attributes by part user attributes.
func (ad *AuthPostgres) GetUserByID(ctx context.Context, userID int) (models.User, error) {
	userDB := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByID, userID)
		if stdErrors.Is(rowUser.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err := rowUser.Scan(
			&userDB.ID,
			&userDB.email,
			&userDB.nickname,
			&userDB.password)
		if err != nil {
			return err
		}

		rowProfile := conn.QueryRowContext(ctx, getProfileAvatar, userDB.ID)
		if stdErrors.Is(rowProfile.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}
		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err = rowProfile.Scan(&userDB.avatar)
		if err != nil {
			return err
		}

		if !userDB.avatar.Valid {
			userDB.avatar.String = innerPKG.DefPersonAvatar
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: userID - [%d]. Special error - [%s]",
			userID, errMain)
	}

	return userDB.Convert(), nil
}
