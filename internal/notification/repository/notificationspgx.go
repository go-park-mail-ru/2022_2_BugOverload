package repository

import (
	"context"
	"database/sql"
	stdErrors "github.com/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

//go:generate mockgen -source notificationspgx.go -destination mocks/mocknotificationsrepository.go -package mockNotificationsRepository

type NotificationRepository interface {
	GetFilmRelease(ctx context.Context) ([]models.Film, error)
}

// notificationPostgres is implementation repository of Postgres corresponding to the Repository interface.
type notificationPostgres struct {
	database *sqltools.Database
}

// NewFilmPostgres is constructor for notificationPostgres.
func NewFilmPostgres(database *sqltools.Database) NotificationRepository {
	return &notificationPostgres{
		database,
	}
}

// GetFilmRelease is the service that accesses the interface ImageRepository
func (n *notificationPostgres) GetFilmRelease(ctx context.Context) ([]models.Film, error) {
	response := make([]film.ModelSQL, 0)

	var err error

	// Films - Main
	errMain := sqltools.RunQuery(ctx, n.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		response, err = film.GetFilmsRealisesBatch(ctx, conn)
		if err != nil {
			return stdErrors.WithMessagef(errors.ErrNotFoundInDB,
				"Film main info: Special Error [%s]", err)
		}

		return nil
	})

	if errMain != nil {
		return []models.Film{}, errMain
	}

	res := make([]models.Film, len(response))

	for idx, value := range response {
		res[idx] = value.Convert()
	}

	return res, nil
}
