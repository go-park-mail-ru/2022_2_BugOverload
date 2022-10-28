package repository

import (
	"context"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// SessionRepository provides the versatility of session-related repositories.
// Needed to work with stateful session pattern.
type SessionRepository interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (models.Session, error)
	DeleteSession(ctx context.Context) (models.Session, error)
}

// sessionCache is implementation repository of sessions in memory corresponding to the SessionRepository interface.
type sessionCache struct {
	storageUserSession map[string]models.Session
	mu                 *sync.RWMutex
}

// NewSessionCache is constructor for sessionCache.
func NewSessionCache() SessionRepository {
	return &sessionCache{
		make(map[string]models.Session),
		&sync.RWMutex{},
	}
}

// CheckExist is a check for the existence of such a session - cookie by name.
func (cs *sessionCache) CheckExist(sessionID string) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	_, ok := cs.storageUserSession[sessionID]
	return ok
}

// GetUserBySession is returns all user attributes by name session.
func (cs *sessionCache) GetUserBySession(ctx context.Context) (models.User, error) {
	sessionID, _ := ctx.Value(pkgInner.SessionKey).(string)

	if !cs.CheckExist(sessionID) {
		return models.User{}, errors.ErrSessionNotExist
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	session := cs.storageUserSession[sessionID]

	return *session.User, nil
}

// CreateSession is creates a new cookie and its link to the user.
func (cs *sessionCache) CreateSession(ctx context.Context, user *models.User) (models.Session, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	newSessionID, _ := pkg.CryptoRandString(pkgInner.CookieValueLength)

	newSession := models.Session{
		ID:   newSessionID,
		User: user,
	}

	cs.storageUserSession[newSessionID] = newSession

	return newSession, nil
}

// DeleteSession is takes the cookie by name, rolls back the time in it so that it becomes
// irrelevant (it is necessary that the browser deletes the cookie on its side) returns
// the cookie with the new date, and the repository deletes the cookie itself and the connection with the user.
func (cs *sessionCache) DeleteSession(ctx context.Context) (models.Session, error) {
	sessionID, _ := ctx.Value(pkgInner.SessionKey).(string)

	if !cs.CheckExist(sessionID) {
		return models.Session{}, errors.ErrSessionNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageUserSession, sessionID)

	return cs.storageUserSession[sessionID], nil
}
