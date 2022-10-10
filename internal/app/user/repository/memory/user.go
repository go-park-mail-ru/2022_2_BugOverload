package memory

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"sync"
)

// userRepo is TMP impl database for users, where key = User.Email
type userRepo struct {
	storage map[string]models.User
	mu      *sync.Mutex
}

// NewUserRepo is constructor for userRepo
func NewUserRepo() interfaces.UserRepository {
	return &userRepo{
		make(map[string]models.User),
		&sync.Mutex{},
	}
}

// CheckExist is method to check the existence of such a cookie in the database
func (us *userRepo) CheckExist(email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.storage[email]

	return ok
}

// Signup is method for creating a user in database
func (us *userRepo) Signup(ctx context.Context, user *models.User) (models.User, error) {
	if us.CheckExist(user.Email) {
		return models.User{}, errors.ErrSignupUserExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	user.Avatar = "asserts/img/invisibleMan.jpeg"

	us.storage[user.Email] = *user

	return *user, nil
}

// Login return user using email (primary key)
func (us *userRepo) Login(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	return us.storage[user.Email], nil
}
