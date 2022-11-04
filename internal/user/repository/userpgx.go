package repository

import (
	"context"
	"database/sql"
	"sync"

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

func (u userPostgres) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	response := NewUserSQL()

	err := sqltools.RunTx(ctx, innerPKG.TxDefaultOptions, u.database.Connection, func(tx *sql.Tx) error {
		wg := &sync.WaitGroup{}

		var err error

		wg.Add(1)
		go func() {
			defer wg.Done()

			rowUser := tx.QueryRowContext(ctx, getUser, user.ID)
			if rowUser.Err() != nil {
				err = rowUser.Err()
				return
			}

			err = rowUser.Scan(&response.Nickname)
			if err != nil {
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			rowProfile := tx.QueryRowContext(ctx, getUserProfile, user.ID)
			if rowProfile.Err() != nil {
				err = rowProfile.Err()
				return
			}

			err = rowProfile.Scan(
				&response.Profile.JoinedDate,
				&response.Profile.CountViewsFilms,
				&response.Profile.CountCollections,
				&response.Profile.CountReviews,
				&response.Profile.CountRatings)
			if err != nil {
				return
			}
		}()

		wg.Wait()

		return nil
	})

	if stdErrors.Is(err, sql.ErrNoRows) {
		return models.User{}, errors.ErrNotFoundInDB
	}

	if err != nil {
		return models.User{}, errors.ErrPostgresRequest
	}

	return response.Convert(), nil
}

func (u userPostgres) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}

func (u userPostgres) ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	return models.User{}, nil
}
