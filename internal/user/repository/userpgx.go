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
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
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

	errTX := sqltools.RunTxOnConn(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, getUser, user.ID)
		if rowUser.Err() != nil {
			return rowUser.Err()
		}

		err := rowUser.Scan(&response.Nickname)
		if err != nil {
			return err
		}

		rowProfile := tx.QueryRowContext(ctx, getUserProfile, user.ID)
		if stdErrors.Is(rowProfile.Err(), sql.ErrNoRows) {
			return errors.ErrNotFoundInDB
		}

		if rowProfile.Err() != nil {
			return rowProfile.Err()
		}

		err = rowProfile.Scan(
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
	if stdErrors.Is(errTX, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	// execution error
	if errTX != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}

func (u *userPostgres) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}

func (u *userPostgres) ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}
