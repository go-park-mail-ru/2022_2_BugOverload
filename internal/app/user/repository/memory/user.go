package memory

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
)

// userRepo is implementation repository of users in memory corresponding to the UserRepository interface.
type userRepo struct {
	storage map[string]models.User
	mu      *sync.Mutex
}

// NewUserRepo is constructor for userRepo. Accepts only mutex.
func NewUserRepo(mu *sync.Mutex) interfaces.UserRepository {
	return &userRepo{
		make(map[string]models.User),
		mu,
	}
}

// CheckExist is a check for the existence of such a user by email.
func (us *userRepo) CheckExist(email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.storage[email]

	return ok
}

// CreateUser is creates a new user and set default avatar.
func (us *userRepo) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
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
func (us *userRepo) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	return us.storage[user.Email], nil
}
