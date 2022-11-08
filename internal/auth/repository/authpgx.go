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
	GetUser(ctx context.Context, user *models.User) (models.User, error)
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

// GetUser is returns all user attributes by part user attributes.
func (ad *AuthDatabase) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !ad.CheckExist(ctx, user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	userDB := NewUserSQL()
	avatar := ""

	err := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByEmail, user.Email)
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err := rowUser.Scan(
			&userDB.userID,
			&userDB.email,
			&userDB.nickname,
			&userDB.password)
		if err != nil {
			return err
		}

		rowProfile := conn.QueryRowContext(ctx, getProfileAvatar, userDB.userID)
		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err = rowProfile.Scan(&avatar)
		if err == sql.ErrNoRows {
			avatar = innerPKG.DefUserAvatar
		}

		return nil
	})

	if err != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return NewUser(userDB, avatar), nil
}
