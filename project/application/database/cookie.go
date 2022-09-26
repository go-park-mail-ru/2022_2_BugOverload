package database

import (
	"errors"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type CookieStorage struct {
	storage map[string]structs.User
}

func NewCookieStorage() *CookieStorage {
	return &CookieStorage{make(map[string]structs.User)}
}

func (us *CookieStorage) CheckExist(email string) error {
	_, ok := us.storage[email]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

func (us *CookieStorage) Insert(u structs.User) {
	us.storage[u.Email] = u
}
