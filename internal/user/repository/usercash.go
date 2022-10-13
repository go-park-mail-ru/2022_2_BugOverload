package repository

import (
	"context"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// UserRepository provides the versatility of films repositories.
type UserRepository interface {
	GetUser(ctx context.Context, user *models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
}

// userCash is implementation repository of users in memory corresponding to the UserRepository interface.
type userCash struct {
	storage map[string]models.User
	mu      *sync.Mutex
}

// NewUserCash is constructor for userCash. Accepts only mutex.
func NewUserCash() UserRepository {
	return &userCash{
		make(map[string]models.User),
		&sync.Mutex{},
	}
}

// CheckExist is a check for the existence of such a user by email.
func (us *userCash) CheckExist(email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.storage[email]

	return ok
}

// CreateUser is creates a new user and set default avatar.
func (us *userCash) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	if us.CheckExist(user.Email) {
		return models.User{}, errors.ErrSignupUserExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	user.Avatar = "asserts/img/invisibleMan.jpeg"

	us.storage[user.Email] = *user

	return *user, nil
}

// GetUser is returns all user attributes by part user attributes.
func (us *userCash) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	return us.storage[user.Email], nil
}
