package database

import (
	"errors"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type UserStorage struct {
	storage map[structs.User]structs.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[structs.User]structs.User)}
}

func (us *UserStorage) CheckExist(u structs.User) error {
	_, ok := us.storage[u]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

func (us *UserStorage) Create(u structs.User) {
	us.storage[u] = u
}
