package database

import (
	"errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserStorage is TMP impl database for users, where key = User.Email
type UserStorage struct {
	storage map[string]structs.User
}

// NewUserStorage is constructor for UserStorage
func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[string]structs.User)}
}

// CheckExist is method to check the existence of such a cookie in the database
func (us *UserStorage) CheckExist(email string) error {
	_, ok := us.storage[email]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

// Create is method for creating a user in database
func (us *UserStorage) Create(u structs.User) {
	us.storage[u.Email] = u
}

// Return user using email (primary key)
func (us *UserStorage) GetUser(email string) (structs.User, error) {
	if us.CheckExist(email) == nil {
		return structs.User{}, errors.New("such user doesn't exist")
	}
	return us.storage[email], nil
}
