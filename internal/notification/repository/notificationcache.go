package repository

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type NotificationHub interface {
	UpdateHub([]models.Film)
}

// notificationPostgres is implementation repository of Postgres corresponding to the Repository interface.
type notificationCache struct {
	hub []models.Film
}

// NewNotificationCache is constructor for notificationPostgres.
func NewNotificationCache() NotificationHub {
	return &notificationCache{
		hub: make([]models.Film, 0),
	}
}

func (n *notificationCache) UpdateHub(films []models.Film) {
	n.hub = films
}
