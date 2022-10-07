package database

import (
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserStorage is TMP impl database for users, where key = User.Email
type UserStorage struct {
	storage map[string]structs.User
	mu      *sync.Mutex
}

// NewUserStorage is constructor for UserStorage
func NewUserStorage() *UserStorage {
	return &UserStorage{
		make(map[string]structs.User),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a cookie in the database
func (us *UserStorage) CheckExist(email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.storage[email]
	return ok
}

// Create is method for creating a user in database
func (us *UserStorage) Create(u structs.User) {
	us.mu.Lock()
	defer us.mu.Unlock()

	u.Avatar = "asserts/img/invisibleMan.jpeg"

	us.storage[u.Email] = u
}

// GetUser return user using email (primary key)
func (us *UserStorage) GetUser(email string) (structs.User, error) {
	if !us.CheckExist(email) {
		return structs.User{}, errors.ErrUserNotExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	return us.storage[email], nil
}
