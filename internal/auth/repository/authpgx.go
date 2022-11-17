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
	UpdatePassword(ctx context.Context, user *models.User, password string) error
	CheckExistUserByEmail(ctx context.Context, email string) (bool, error)
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

// CheckExistUserByEmail is a check for the existence of such a user by email.
func (ad *AuthDatabase) CheckExistUserByEmail(ctx context.Context, email string) (bool, error) {
	response := false
	errMain := sqltools.RunQuery(ctx, ad.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
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

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return false, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Err: params input: query - [%s], email - [%s]. Special error [%s]",
			checkExist, email, errMain)
	}

	if errMain != nil {
		return false, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], email - [%s]. Special error [%s]",
			checkExist, email, errMain)
	}

	return response, nil
}

// CreateUser is creates a new user and set default avatar.
func (ad *AuthDatabase) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	exist, err := ad.CheckExistUserByEmail(ctx, user.Email)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "CreateUser falls on check exist")
	}
	if exist {
		return models.User{}, errors.ErrUserExist
	}

	var userID int

	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, ad.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, createUser, user.Email, user.Nickname, []byte(user.Password), innerPKG.DefUserAvatar)
		if rowUser.Err() != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %s, %s]. Special error: [%s]",
				createUser, user.Email, user.Nickname, user.Password, rowUser.Err())
		}

		err = rowUser.Scan(&userID)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err Scan: params input: query - [%s], values - [%s, %s, %s]. Special error: [%s]",
				createUser, user.Email, user.Nickname, user.Password, rowUser.Err())
		}

		rowsCollections, errCollections := tx.QueryContext(ctx, createDefCollections)
		if errCollections != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s]. Special error: [%s]",
				createDefCollections, errCollections)
		}
		defer rowsCollections.Close()

		ids := make([]int, 0)

		for rowsCollections.Next() {
			var colID int
			err = rowsCollections.Scan(&colID)
			if err != nil {
				return stdErrors.WithMessagef(errors.ErrWorkDatabase,
					"Err: params input: query - [%s]. Special error: [%s]",
					createDefCollections, errCollections)
			}

			ids = append(ids, colID)
		}

		_, err = tx.ExecContext(ctx, linkUserDefCollections, userID, ids[0], ids[1])
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrWorkDatabase,
				"Err: params input: query - [%s], values -[%d, %d, %d]. Special error: [%s]",
				linkUserDefCollections, userID, ids[0], ids[1], errCollections)
		}

		return nil
	})

	if errMain != nil {
		return models.User{}, errMain
	}

	return models.User{ID: userID}, nil
}

// GetUserByEmail is returns all user attributes by part user attributes.
func (ad *AuthDatabase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
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
			&userDB.password,
			&userDB.avatar)
		if err != nil {
			return err
		}

		if !userDB.avatar.Valid {
			userDB.avatar.String = innerPKG.DefPersonAvatar
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Err: params input: query - [%s], values - [%s]. Special error - [%s]",
			getUserByEmail, email, errMain)
	}

	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%s]. Special error - [%s]",
			getUserByEmail, email, errMain)
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
			&userDB.password,
			&userDB.avatar)
		if err != nil {
			return err
		}

		return nil
	})

	if stdErrors.Is(errMain, sql.ErrNoRows) {
		return models.User{}, stdErrors.WithMessagef(errors.ErrNotFoundInDB,
			"Err: params input: query - [%s], values - [%d]. Special error - [%s]",
			getUserByID, userID, errMain)
	}

	if errMain != nil {
		return models.User{}, stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%d]. Special error - [%s]",
			getUserByID, userID, errMain)
	}

	return userDB.Convert(), nil
}

func (ad *AuthDatabase) UpdatePassword(ctx context.Context, user *models.User, password string) error {
	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, ad.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserPassword, []byte(password), user.ID)
		if err != nil {
			return err
		}
		return nil
	})

	if errMain != nil {
		return stdErrors.WithMessagef(errors.ErrWorkDatabase,
			"Err: params input: query - [%s], values - [%s, %d]. Special Error [%s]",
			updateUserPassword, user.Password, user.ID, errMain)
	}

	return nil
}
