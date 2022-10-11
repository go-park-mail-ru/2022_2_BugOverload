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

type cookieRepo struct {
	storageCookie     map[string]http.Cookie
	storageUserCookie map[string]*models.User
	mu                *sync.Mutex
}

func NewCookieRepo(mu *sync.Mutex) interfaces.AuthRepository {
	return &cookieRepo{
		make(map[string]http.Cookie),
		make(map[string]*models.User),
		mu,
	}
}

func (cs *cookieRepo) CheckExist(cookie string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.storageCookie[cookie]
	return ok
}

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
