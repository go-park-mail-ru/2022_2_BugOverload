package repository

import (
	"context"
	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// AuthRepository provides the versatility of session-related repositories.
// Needed to work with stateful session pattern.
type AuthRepository interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}

// cookieCache is implementation repository of sessions in memory corresponding to the AuthRepository interface.
type cookieCache struct {
	storageCookie     map[string]http.Cookie
	storageUserCookie map[string]*models.User
	mu                *sync.RWMutex
}

// NewCookieCache is constructor for cookieCache.
func NewCookieCache() AuthRepository {
	return &cookieCache{
		make(map[string]http.Cookie),
		make(map[string]*models.User),
		&sync.RWMutex{},
	}
}

// CheckExist is a check for the existence of such a session - cookie by name.
func (cs *cookieCache) CheckExist(cookie string) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	_, ok := cs.storageCookie[cookie]
	return ok
}

// GetUserBySession is returns all user attributes by name session.
func (cs *cookieCache) GetUserBySession(ctx context.Context) (models.User, error) {
	cookie, _ := ctx.Value(pkgInner.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return models.User{}, errors.ErrCookieNotExist
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	user := cs.storageUserCookie[cookie]

	return *user, nil
}

// CreateSession is creates a new cookie and its link to the user.
func (cs *cookieCache) CreateSession(ctx context.Context, user *models.User) (string, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	expiration := time.Now().Add(pkgInner.TimeoutLiveCookie)

	sessionID := strconv.Itoa(len(cs.storageCookie) + 1)

	newCookieValue, _ := pkg.CryptoString(pkgInner.CookieValueLength)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    newCookieValue,
		Expires:  expiration,
		HttpOnly: true,
	}

	cookieKey := sessionID + "=" + newCookieValue

	cs.storageCookie[cookieKey] = cookie
	cs.storageUserCookie[cookieKey] = user

	return cookie.String(), nil
}

// GetSession is to get all attributes of a cookie in string format.
func (cs *cookieCache) GetSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(pkgInner.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	resCookie := cs.storageCookie[cookie]

	return resCookie.String(), nil
}

// DeleteSession is takes the cookie by name, rolls back the time in it so that it becomes
// irrelevant (it is necessary that the browser deletes the cookie on its side) returns
// the cookie with the new date, and the repository deletes the cookie itself and the connection with the user.
func (cs *cookieCache) DeleteSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(pkgInner.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageCookie, cookie)
	defer delete(cs.storageUserCookie, cookie)

	oldCookie := cs.storageCookie[cookie]

	oldCookie.Expires = time.Now().Add(-pkgInner.TimeoutLiveCookie)

	return oldCookie.String(), nil
}
