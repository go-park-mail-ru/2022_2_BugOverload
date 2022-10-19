package repository

import (
	"context"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// UserRepository provides the versatility of users repositories.
type UserRepository interface {
	GetUser(ctx context.Context, user *models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
}

// userCache is implementation repository of users in memory corresponding to the UserRepository interface.
type userCache struct {
	storage map[string]models.User
	mu      *sync.RWMutex
}

// NewUserCache is constructor for userCache. Accepts only mutex.
func NewUserCache() UserRepository {
	return &userCache{
		make(map[string]models.User),
		&sync.RWMutex{},
	}
}

// CheckExist is a check for the existence of such a user by email.
func (us *userCache) CheckExist(email string) bool {
	us.mu.RLock()
	defer us.mu.RUnlock()

	_, ok := us.storage[email]

	return ok
}

// CreateUser is creates a new user and set default avatar.
func (us *userCache) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
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
func (us *userCache) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.RLock()
	defer us.mu.RUnlock()

	return us.storage[user.Email], nil
}
