package memory

import (
	"context"
	"sync"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

type userRepo struct {
	storage map[string]models.User
	mu      *sync.Mutex
}

func NewUserRepo(mu *sync.Mutex) interfaces.UserRepository {
	return &userRepo{
		make(map[string]models.User),
		mu,
	}
}

func (us *userRepo) CheckExist(email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.storage[email]

	return ok
}

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

func (us *userRepo) GetUser(ctx context.Context, user *models.User) (models.User, error) {
	if !us.CheckExist(user.Email) {
		return models.User{}, errors.ErrUserNotExist
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	return us.storage[user.Email], nil
}
