package memory

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
)

// cookieRepo is implementation repository of sessions in memory corresponding to the AuthRepository interface.
type cookieRepo struct {
	storageCookie     map[string]http.Cookie
	storageUserCookie map[string]*models.User
	mu                *sync.Mutex
}

// NewCookieRepo is constructor for cookieRepo. Accepts only mutex.
func NewCookieRepo(mu *sync.Mutex) interfaces.AuthRepository {
	return &cookieRepo{
		make(map[string]http.Cookie),
		make(map[string]*models.User),
		mu,
	}
}

// CheckExist is a check for the existence of such a session - cookie by name.
func (cs *cookieRepo) CheckExist(cookie string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.storageCookie[cookie]
	return ok
}

// GetUserBySession is returns all user attributes by name session.
func (cs *cookieRepo) GetUserBySession(ctx context.Context) (models.User, error) {
	cookie, _ := ctx.Value(params.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return models.User{}, errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	user := cs.storageUserCookie[cookie]

	return *user, nil
}

// CreateSession is creates a new cookie and its link to the user.
func (cs *cookieRepo) CreateSession(ctx context.Context, user *models.User) (string, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	expiration := time.Now().Add(params.TimeoutLiveCookie)

	sessionID := strconv.Itoa(len(cs.storageCookie) + 1)

	newCookieValue, _ := utils.CryptoString(params.CookieValueLength)

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
func (cs *cookieRepo) GetSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(params.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	resCookie := cs.storageCookie[cookie]

	return resCookie.String(), nil
}

// DeleteSession is takes the cookie by name, rolls back the time in it so that it becomes
// irrelevant (it is necessary that the browser deletes the cookie on its side) returns
// the cookie with the new date, and the repository deletes the cookie itself and the connection with the user.
func (cs *cookieRepo) DeleteSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(params.CookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageCookie, cookie)
	defer delete(cs.storageUserCookie, cookie)

	oldCookie := cs.storageCookie[cookie]

	oldCookie.Expires = time.Now().Add(-params.TimeoutLiveCookie)

	return oldCookie.String(), nil
}
