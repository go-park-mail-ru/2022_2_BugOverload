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

type UserRepository interface {
	// Profile + settings
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)

	// ChangeInfo
	ChangeUserProfileNickname(ctx context.Context, user *models.User) error
	ChangeUserProfilePassword(ctx context.Context, user *models.User) error

	// Support
	GetPassword(ctx context.Context, user *models.User) (string, error)

	// Film
	FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error
	FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error
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
		return models.User{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getUser, user.ID, errMain)
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
		return models.User{}, stdErrors.WithMessagef(errors.ErrPostgresRequest,
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
		return stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%s, %d]. Special Error [%s]",
			updateUserSettingsNickname, user.Nickname, user.ID, errMain)
	}

	return nil
}

func (u *userPostgres) ChangeUserProfilePassword(ctx context.Context, user *models.User) error {
	errMain := sqltools.RunTxOnConn(ctx, innerPKG.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, updateUserSettingsPassword, []byte(user.Password), user.ID)
		if err != nil {
			return err
		}
		return nil
	})

	// execution error
	if errMain != nil {
		return stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%s, %d]. Special Error [%s]",
			updateUserSettingsPassword, user.Password, user.ID, errMain)
	}

	return nil
}

func (u *userPostgres) GetPassword(ctx context.Context, user *models.User) (string, error) {
	var res []byte

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, getPass, user.ID)
		if row.Err() != nil {
			return row.Err()
		}

		err := row.Scan(&res)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return "", stdErrors.WithMessagef(errors.ErrPostgresRequest,
			"Err: params input: query - [%s], values - [%d]. Special Error [%s]",
			getPass, user.ID, errMain)
	}

	return string(res), nil
}

func (u *userPostgres) FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error {
	return nil
}

func (u *userPostgres) FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error {
	return nil
}
