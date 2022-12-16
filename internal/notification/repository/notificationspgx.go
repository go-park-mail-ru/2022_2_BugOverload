package repository

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

//go:generate mockgen -source filmpgx.go -destination mocks/mockfilmrepository.go -package mockFilmRepository

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
func (i *notificationPostgres) GetFilmRelease(ctx context.Context) ([]models.Film, error) {
	return []models.Film{}, nil
}
