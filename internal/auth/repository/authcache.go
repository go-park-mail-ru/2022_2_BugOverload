package repository

import (
	"context"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// AuthRepository provides the versatility of users repositories.
type AuthRepository interface {
	GetUser(ctx context.Context, user *models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
}

// AuthCache is implementation repository of users in memory corresponding to the AuthRepository interface.
type AuthCache struct {
	storage map[string]models.User
	mu      *sync.RWMutex
}

// NewAuthCache is constructor for AuthCache. Accepts only mutex.
func NewAuthCache() AuthRepository {
	return &AuthCache{
		make(map[string]models.User),
		&sync.RWMutex{},
	}
}

// CheckExist is a check for the existence of such a user by email.
func (us *AuthCache) CheckExist(email string) bool {
	us.mu.RLock()
	defer us.mu.RUnlock()

	_, ok := us.storage[email]

	return ok
}

// CreateUser is creates a new user and set default avatar.
func (us *AuthCache) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	if us.CheckExist(user.Email) {
		return models.User{}, errors.ErrSignupUserExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	user.Profile = models.Profile{
		Avatar: "avatar",
	}

	us.storage[user.Email] = *user

	return *user, nil
}

// GetUser is returns all user attributes by part user attributes.
func (us *AuthCache) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.RLock()
	defer us.mu.RUnlock()

	return us.storage[user.Email], nil
}
