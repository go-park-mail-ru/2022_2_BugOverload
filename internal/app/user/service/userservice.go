package service

import (
	"context"
	"time"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

// userService is implementation for users service corresponding to the UserService interface.
type userService struct {
	userRepo       userInterface.UserRepository
	contextTimeout time.Duration
}

// NewUserService is constructor for userService. Accepts UserRepository interfaces and context timeout.
func NewUserService(ur userInterface.UserRepository, timeout time.Duration) userInterface.UserService {
	return &userService{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

// Login is the service that accesses the interface UserRepository.
// Validation: request password and password user from repository equal.
func (u userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	userRepo, err := u.userRepo.GetUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}

	if userRepo.Password != user.Password {
		return models.User{}, errors.ErrLoginCombinationNotFound
	}

	return userRepo, nil
}

// Signup is the service that accesses the interface UserRepository
func (u userService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	return newUser, nil
}
