package database

import (
	"errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserStorage is TMP impl database for users
type UserStorage struct {
	storage map[structs.User]structs.User
}

// NewUserStorage is constructor for UserStorage
func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[structs.User]structs.User)}
}

// CheckExist is method to check the existence of such a cookie in the database
func (us *UserStorage) CheckExist(u structs.User) error {
	_, ok := us.storage[u]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

// Create is method for creating a user in database
func (us *UserStorage) Create(u structs.User) {
	us.storage[u] = u
}
