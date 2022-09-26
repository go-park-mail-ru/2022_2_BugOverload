package database

import (
	"errors"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// CookieStorage is TMP impl database for cookie
type CookieStorage struct {
	storage map[string]structs.User
}

// NewCookieStorage is constructor for CookieStorage
func NewCookieStorage() *CookieStorage {
	return &CookieStorage{make(map[string]structs.User)}
}

// CheckExist is method to check the existence of such a cookie in the database
func (us *CookieStorage) CheckExist(email string) error {
	_, ok := us.storage[email]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

// Insert is method for creating a cookie
func (us *CookieStorage) Insert(u structs.User) {
	us.storage[u.Email] = u
}
