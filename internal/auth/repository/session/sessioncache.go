package session

import (
	"context"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// Repository provides the versatility of session-related repositories.
// Needed to work with stateful session pattern.
type Repository interface {
	GetUserBySession(ctx context.Context, session models.Session) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (models.Session, error)
	DeleteSession(ctx context.Context, session models.Session) (models.Session, error)
}

// sessionCache is implementation repository of sessions in memory corresponding to the Repository interface.
type sessionCache struct {
	storageUserSession map[string]models.Session
	mu                 *sync.RWMutex
}

// NewSessionCache is constructor for sessionCache.
func NewSessionCache() Repository {
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
func (cs *sessionCache) GetUserBySession(ctx context.Context, session models.Session) (models.User, error) {
	if !cs.CheckExist(session.ID) {
		return models.User{}, errors.ErrSessionNotFound
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	sessionUser := cs.storageUserSession[session.ID]

	return *sessionUser.User, nil
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
func (cs *sessionCache) DeleteSession(ctx context.Context, session models.Session) (models.Session, error) {
	if !cs.CheckExist(session.ID) {
		return models.Session{}, errors.ErrSessionNotFound
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageUserSession, session.ID)

	return cs.storageUserSession[session.ID], nil
}
