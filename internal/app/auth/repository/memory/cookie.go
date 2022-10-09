package memory

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

const TimeoutLiveCookie = 10 * time.Hour

// cookieRepo is TMP impl database for cookie
type cookieRepo struct {
	storage map[string]http.Cookie
	mu      *sync.Mutex
}

// NewCookieRepo is constructor for cookieRepo
func NewCookieRepo() *cookieRepo {
	return &cookieRepo{
		make(map[string]http.Cookie),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a cookie in the database
func (cs *cookieRepo) CheckExist(cookie string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.storage[cookie]
	return ok
}

// CreateSession is method for creating a cookie
func (cs *cookieRepo) CreateSession(ctx context.Context, user *models.User) (string, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	expiration := time.Now().Add(TimeoutLiveCookie)

	sessionID := strconv.Itoa(len(cs.storage) + 1)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    user.Email,
		Expires:  expiration,
		HttpOnly: true,
	}

	cookieStrFullName := sessionID + "=" + user.Email

	cs.storage[cookieStrFullName] = cookie

	return cookie.String(), nil
}

type key string

const cookieKey key = "cookie"

// GetSession return user using email (primary key)
func (cs *cookieRepo) GetSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(cookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	resCookie := cs.storage[cookie]

	return resCookie.String(), nil
}

// DeleteSession delete cookie from storage
func (cs *cookieRepo) DeleteSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value(cookieKey).(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storage, cookie)

	oldCookie := cs.storage[cookie]

	oldCookie.Expires = time.Now().Add(-TimeoutLiveCookie)

	return oldCookie.String(), nil
}
