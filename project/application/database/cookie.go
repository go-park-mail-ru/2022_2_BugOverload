package database

import (
	"net/http"
	"strconv"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
)

const TimeoutLiveCookie = 10

// CookieStorage is TMP impl database for cookie
type CookieStorage struct {
	storage map[string]http.Cookie
}

// NewCookieStorage is constructor for CookieStorage
func NewCookieStorage() *CookieStorage {
	return &CookieStorage{make(map[string]http.Cookie)}
}

// CheckExist is method to check the existence of such a cookie in the database
func (cs *CookieStorage) CheckExist(cookie string) bool {
	_, ok := cs.storage[cookie]
	return ok
}

// Create is method for creating a cookie
func (cs *CookieStorage) Create(email string) string {
	expiration := time.Now().Add(TimeoutLiveCookie * time.Hour)

	sessionID := strconv.Itoa(len(cs.storage) + 1)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    email,
		Expires:  expiration,
		HttpOnly: true,
	}

	cookieStr := cookie.String()

	cs.storage[cookieStr] = cookie

	return cookieStr
}

// GetCookie return user using email (primary key)
func (cs *CookieStorage) GetCookie(cookie string) (http.Cookie, error) {
	if !cs.CheckExist(cookie) {
		return http.Cookie{}, errorshandlers.ErrCookieNotExist
	}

	return cs.storage[cookie], nil
}

// DeleteCookie delete cookie from storage
func (cs *CookieStorage) DeleteCookie(cookie string) error {
	if !cs.CheckExist(cookie) {
		return errorshandlers.ErrCookieNotExist
	}

	return nil
}
