package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

type UserRepository interface {
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
}

// userPostgres is implementation repository of Postgres corresponding to the UserRepository interface.
type userPostgres struct {
	Database *sqltools.Database
}

// NewUserPostgres is constructor for userPostgres.
func NewUserPostgres(Database *sqltools.Database) UserRepository {
	return &userPostgres{
		Database,
	}
}

func (u userPostgres) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	response := newUserSQL()

	err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, u.Database.Connection, func(tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, getUser, user.ID)
		if rowUser.Err() != nil {
			logrus.Error("user ", rowUser.Err())
			return rowUser.Err()
		}

		err := rowUser.Scan(&response.Nickname)
		if err != nil {
			logrus.Error("user ", err)
			return err
		}

		rowProfile := tx.QueryRowContext(ctx, getUserProfile, user.ID)
		if rowProfile.Err() != nil {
			logrus.Error("profile ", rowProfile.Err())
			return rowProfile.Err()
		}

		err = rowProfile.Scan(
			&response.Profile.JoinedDate,
			&response.Profile.CountViewsFilms,
			&response.Profile.CountCollections,
			&response.Profile.CountReviews,
			&response.Profile.CountRatings)
		if err != nil {
			logrus.Error("profile ", err)
			return err
		}

		return nil
	})

	if err != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return response.convert(), nil
}

func (u userPostgres) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}

func (u userPostgres) ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}
