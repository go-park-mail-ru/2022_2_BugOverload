package database

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"net/http"
	"strconv"
	"time"
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
func (cs *CookieStorage) CheckExist(email string) error {
	_, ok := cs.storage[email]
	if ok {
		return errorshandlers.ErrUserExist
	}

	return nil
}

// Create is method for creating a cookie
func (cs *CookieStorage) Create(email string) {
	expiration := time.Now().Add(TimeoutLiveCookie * time.Hour)

	sessionID := strconv.Itoa(len(cs.storage) + 1)

	cookie := http.Cookie{
		Name:     sessionID,
		Value:    email,
		Expires:  expiration,
		HttpOnly: true,
	}

	cs.storage[email] = cookie
}

// GetCookie return exist cookie
func (cs *CookieStorage) GetCookie(email string) (http.Cookie, error) {
	if cs.CheckExist(email) == nil {
		return http.Cookie{}, errorshandlers.ErrUserNotExist
	}

	return cs.storage[email], nil
}
