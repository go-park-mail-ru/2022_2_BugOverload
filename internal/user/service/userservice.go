package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
)

// UserService provides universal service for work with users.
type UserService interface {
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
}

// userService is implementation for users service corresponding to the UserService interface.
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService is constructor for userService. Accepts UserRepository interfaces.
func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

// Login is the service that accesses the interface UserRepository.
// Validation: request password and password user from repository equal.
func (u *userService) Login(ctx context.Context, user *models.User) (models.User, error) {
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
func (u *userService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	return newUser, nil
}
