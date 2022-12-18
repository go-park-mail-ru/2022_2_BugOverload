package repository

import "sync"

type NotificationHub interface {
	Update([]interface{})
	Get() []interface{}
}

// notificationPostgres is implementation repository of Postgres corresponding to the Repository interface.
type notificationCache struct {
	mu  *sync.RWMutex
	hub []interface{}
}

// NewNotificationCache is constructor for notificationPostgres.
func NewNotificationCache() NotificationHub {
	return &notificationCache{
		mu:  &sync.RWMutex{},
		hub: make([]interface{}, 0),
	}
}

func (n *notificationCache) Update(messages []interface{}) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.hub = messages
}

func (n *notificationCache) Get() []interface{} {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.hub
}
