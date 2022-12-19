package repository

import (
	"sync"

	"golang.org/x/exp/maps"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

type NotificationHub interface {
	UpdateHub([]interface{})
	GetNotifications(user *models.User) []interface{}
	CheckNewNotification(user *models.User) bool
}

// notificationPostgres is implementation repository of Postgres corresponding to the Repository interface.
type notificationCache struct {
	mu   *sync.RWMutex
	hub  []interface{}
	sent map[int]bool
}

// NewNotificationCache is constructor for notificationPostgres.
func NewNotificationCache() NotificationHub {
	return &notificationCache{
		mu:   &sync.RWMutex{},
		hub:  make([]interface{}, 0),
		sent: make(map[int]bool, 0),
	}
}

func (n *notificationCache) UpdateHub(messages []interface{}) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.hub = messages

	maps.Clear(n.sent)
}

func (n *notificationCache) GetNotifications(user *models.User) []interface{} {
	n.mu.RLock()
	defer n.mu.RUnlock()

	n.sent[user.ID] = true

	return n.hub
}

func (n *notificationCache) CheckNewNotification(user *models.User) bool {
	n.mu.RLock()
	defer n.mu.RUnlock()

	sent, ok := n.sent[user.ID]
	if sent || !ok {
		return false
	}

	return true
}
