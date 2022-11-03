package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// AuthService provides universal service for work with users.
type AuthService interface {
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
}

// authService is implementation for users service corresponding to the AuthService interface.
type authService struct {
	authRepo repository.AuthRepository
}

// NewAuthService is constructor for authService. Accepts AuthRepository interfaces.
func NewAuthService(ur repository.AuthRepository) AuthService {
	return &authService{
		authRepo: ur,
	}
}

// Login is the service that accesses the interface AuthRepository.
// Validation: request password and password user from repository equal.
func (u *authService) Login(ctx context.Context, user *models.User) (models.User, error) {
	userRepo, err := u.authRepo.GetUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}

	if userRepo.Password != user.Password {
		return models.User{}, errors.ErrLoginCombinationNotFound
	}

	return userRepo, nil
}

// Signup is the service that accesses the interface AuthRepository
func (u *authService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.authRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	return newUser, nil
}
