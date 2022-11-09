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

// AuthDatabase is implementation repository of users to the AuthRepository interface.
type AuthDatabase struct {
	database *sqltools.Database
}

// NewAuthDatabase is constructor for AuthDatabase. Accepts only sqltools.Database.
func NewAuthDatabase(database *sqltools.Database) AuthRepository {
	return &AuthDatabase{
		database,
	}
}

// CheckExist is a check for the existence of such a user by email.
func (ad *AuthDatabase) CheckExist(ctx context.Context, email string) bool {
	response := false
	err := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkExist, email)
		if row.Err() != nil {
			return row.Err()
		}

		err := row.Scan(&response)
		if err != nil {
			return err
		}

		return nil
	})

	if stdErrors.Is(err, sql.ErrNoRows) {
		return false
	}

	if err != nil {
		return false
	}

	return response
}

// CreateUser is creates a new user and set default avatar.
func (ad *AuthDatabase) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	if ad.CheckExist(ctx, user.Email) {
		return models.User{}, errors.ErrSignupUserExist
	}

	err := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, ad.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, createUser, user.Email, user.Nickname, user.Password)
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err := rowUser.Scan(&user.ID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, createUserProfile, user.ID)
		if err != nil {
			return err
		}

		rowsCollections, err := tx.QueryContext(ctx, createDefCollections)
		if err != nil {
			return err
		}

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

	if err != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	user.Profile = models.Profile{
		Avatar: innerPKG.DefUserAvatar,
	}

	return *user, nil
}

// GetUserByEmail is returns all user attributes by part user attributes.
func (ad *AuthDatabase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	if !ad.CheckExist(ctx, email) {
		return models.User{}, errors.ErrUserNotExist
	}

	userDB := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByEmail, email)
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

	if errMain != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return userDB.Convert(), nil
}

// GetUserByID is returns all user attributes by part user attributes.
func (ad *AuthDatabase) GetUserByID(ctx context.Context, userID int) (models.User, error) {
	userDB := NewUserSQL()

	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByID, userID)
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

	if errMain != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return userDB.Convert(), nil
}
