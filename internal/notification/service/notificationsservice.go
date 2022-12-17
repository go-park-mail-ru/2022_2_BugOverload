package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/notification/repository"
)

//go:generate mockgen -source imageservice.go -destination mocks/mockimageservice.go -package mockImageService

// NotificationsService provides universal service for work with images.
type NotificationsService interface {
	GetFilmRelease(ctx context.Context) ([]models.Film, error)
}

// notificationsService is implementation for users service corresponding to the NotificationsService interface.
type notificationsService struct {
	notificationHub  repository.NotificationHub
	notificationRepo repository.NotificationRepository
}

// NewImageService is constructor for imageService. Accepts ImageService interfaces.
func NewImageService(r repository.NotificationRepository) NotificationsService {
	return &notificationsService{
		notificationRepo: r,
	}
}

// GetFilmRelease is the service that accesses the interface ImageRepository
func (i *notificationsService) GetFilmRelease(ctx context.Context) ([]models.Film, error) {
	films, err := i.notificationRepo.GetFilmRelease(ctx)
	if err != nil {
		return []models.Film{}, stdErrors.Wrap(err, "GetFilmRelease")
	}

	return films, nil
}
