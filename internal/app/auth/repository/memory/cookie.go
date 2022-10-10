package memory

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

const TimeoutLiveCookie = 10 * time.Hour

// cookieRepo is TMP impl database for cookie
type cookieRepo struct {
	storageCookie     map[string]http.Cookie
	storageUserCookie map[string]*models.User
	mu                *sync.Mutex
}

// NewCookieRepo is constructor for cookieRepo
func NewCookieRepo() interfaces.AuthRepository {
	return &cookieRepo{
		make(map[string]http.Cookie),
		make(map[string]*models.User),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a cookie in the database
func (cs *cookieRepo) CheckExist(cookie string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.storageCookie[cookie]
	return ok
}

// GetUserBySession is method for creating a cookie
func (cs *cookieRepo) GetUserBySession(ctx context.Context) (models.User, error) {
	cookie, _ := ctx.Value("cookie").(string)

	if !cs.CheckExist(cookie) {
		return models.User{}, errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	user := cs.storageUserCookie[cookie]

	return *user, nil

}

// CreateSession is method for creating a cookie
func (cs *cookieRepo) CreateSession(ctx context.Context, user *models.User) (string, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	expiration := time.Now().Add(TimeoutLiveCookie)

	sessionID := strconv.Itoa(len(cs.storageCookie) + 1)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    user.Email,
		Expires:  expiration,
		HttpOnly: true,
	}

	cookieStrFullName := sessionID + "=" + user.Email

	cs.storageCookie[cookieStrFullName] = cookie
	cs.storageUserCookie[cookieStrFullName] = user

	return cookie.String(), nil
}

// GetSession return user using email (primary key)
func (cs *cookieRepo) GetSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value("cookie").(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	resCookie := cs.storageCookie[cookie]

	return resCookie.String(), nil
}

// DeleteSession delete cookie from storage
func (cs *cookieRepo) DeleteSession(ctx context.Context) (string, error) {
	cookie, _ := ctx.Value("cookie").(string)

	if !cs.CheckExist(cookie) {
		return "", errors.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storageCookie, cookie)
	defer delete(cs.storageUserCookie, cookie)

	oldCookie := cs.storageCookie[cookie]

	oldCookie.Expires = time.Now().Add(-TimeoutLiveCookie)

	return oldCookie.String(), nil
}
