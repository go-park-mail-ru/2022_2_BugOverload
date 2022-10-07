package database

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
)

const TimeoutLiveCookie = 10 * time.Hour

// CookieStorage is TMP impl database for cookie
type CookieStorage struct {
	storage map[string]http.Cookie
	mu      *sync.Mutex
}

// NewCookieStorage is constructor for CookieStorage
func NewCookieStorage() *CookieStorage {
	return &CookieStorage{
		make(map[string]http.Cookie),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a cookie in the database
func (cs *CookieStorage) CheckExist(cookie string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.storage[cookie]
	return ok
}

// Create is method for creating a cookie
func (cs *CookieStorage) Create() string {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	expiration := time.Now().Add(TimeoutLiveCookie)

	sessionID := strconv.Itoa(len(cs.storage) + 1)

	randStr := RandStringRunes(30)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    RandStringRunes(30),
		Expires:  expiration,
		HttpOnly: true,
	}

	cookieStrFullName := sessionID + "=" + randStr

	cs.storage[cookieStrFullName] = cookie

	return cookie.String()
}

// GetCookie return user using email (primary key)
func (cs *CookieStorage) GetCookie(cookie string) (http.Cookie, error) {
	if !cs.CheckExist(cookie) {
		return http.Cookie{}, errorshandlers.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	return cs.storage[cookie], nil
}

// DeleteCookie delete cookie from storage
func (cs *CookieStorage) DeleteCookie(cookie string) (string, error) {
	if !cs.CheckExist(cookie) {
		return "", errorshandlers.ErrCookieNotExist
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	defer delete(cs.storage, cookie)

	oldCookie := cs.storage[cookie]

	oldCookie.Expires = time.Now().Add(-TimeoutLiveCookie)

	return oldCookie.String(), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
