package repository

import (
	"context"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// AuthRepository provides the versatility of session-related repositories.
// Needed to work with stateful session pattern.
type AuthRepository interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}

// sessionCache is implementation repository of sessions in memory corresponding to the AuthRepository interface.
type sessionCache struct {
	storageUserSession map[string]*models.User
	mu                 *sync.RWMutex
}

// NewCookieCache is constructor for sessionCache.
func NewCookieCache() AuthRepository {
	return &sessionCache{
		make(map[string]*models.User),
		&sync.RWMutex{},
	}
}

// CheckExist is a check for the existence of such a session - cookie by name.
func (cs *sessionCache) CheckExist(cookie string) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	_, ok := cs.storageUserSession[cookie]
	return ok
}

// GetUserBySession is returns all user attributes by name session.
func (cs *sessionCache) GetUserBySession(ctx context.Context) (models.User, error) {
	session, _ := ctx.Value(pkgInner.SessionKey).(string)

	if !cs.CheckExist(session) {
		return models.User{}, errors.ErrCookieNotExist
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	user := cs.storageUserSession[session]

	return *user, nil
}

// CreateSession is creates a new cookie and its link to the user.
func (cs *sessionCache) CreateSession(ctx context.Context, user *models.User) (string, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	newCookieValue, _ := pkg.CryptoRandString(pkgInner.CookieValueLength)

	cookieKey := "session_id=" + newCookieValue

	cs.storageUserSession[cookieKey] = user

	return cookieKey, nil
}

// DeleteSession is takes the cookie by name, rolls back the time in it so that it becomes
// irrelevant (it is necessary that the browser deletes the cookie on its side) returns
// the cookie with the new date, and the repository deletes the cookie itself and the connection with the user.
func (cs *sessionCache) DeleteSession(ctx context.Context) (string, error) {
	session, _ := ctx.Value(pkgInner.SessionKey).(string)

	if !cs.CheckExist(session) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageUserSession, session)

	return session, nil
}
