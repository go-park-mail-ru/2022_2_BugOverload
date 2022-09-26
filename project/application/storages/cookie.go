package storages

import (
	"errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type CookieStorage struct {
	storage map[structs.User]string
}

func NewCookieStorage() *CookieStorage {
	return &CookieStorage{make(map[structs.User]string)}
}

func (us *CookieStorage) CheckExist(u structs.User) error {
	_, ok := us.storage[u]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

func (us *CookieStorage) Insert(u structs.User) {
	us.storage[u] = "exist" + u.Email
}
