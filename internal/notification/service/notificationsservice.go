package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/notification/repository"
)

//go:generate mockgen -source notificationsservice.go -destination mocks/mocknotificationsservice.go -package mockNotificationsService

type NotificationsService interface {
	GetMessages(user *models.User) []interface{}
	CheckNewNotification(user *models.User) bool
}

type notificationsService struct {
	notificationHub  repository.NotificationHub
	notificationRepo repository.NotificationRepository
}

func NewNotificationsService(r repository.NotificationRepository, h repository.NotificationHub) NotificationsService {
	res := &notificationsService{
		notificationRepo: r,
		notificationHub:  h,
	}

	go res.UpdateHubDemon()

	return res
}

func (s *notificationsService) GetFilmRelease(ctx context.Context) ([]models.Notification, error) {
	films, err := s.notificationRepo.GetFilmRelease(ctx)
	if err != nil {
		return []models.Notification{}, stdErrors.Wrap(err, "GetFilmRelease")
	}

	res := make([]models.Notification, len(films))

	for idx, val := range films {
		res[idx].Action = "ANONS_FILM"
		res[idx].Payload = NewAnonsFilmNotificationPayload(val)
	}

	return res, nil
}

const (
	maxHour   = 24
	maxMinute = 59
)

func (s *notificationsService) UpdateHubDemon() {
	curTime := time.Now()

	offset := time.Duration(maxHour-curTime.Hour())*time.Hour + time.Duration(maxMinute-curTime.Minute())*time.Minute

	ticker := time.NewTicker(offset)
	for {
		notification, err := s.GetFilmRelease(context.TODO())
		if err != nil {
			break
		}

		messages := make([]interface{}, len(notification))

		for idx, val := range notification {
			messages[idx] = val
		}

		s.notificationHub.UpdateHub(messages)

		<-ticker.C
	}
}

func (s *notificationsService) GetMessages(user *models.User) []interface{} {
	return s.notificationHub.GetNotifications(user)
}

func (s *notificationsService) CheckNewNotification(user *models.User) bool {
	return s.notificationHub.CheckNewNotification(user)
}
