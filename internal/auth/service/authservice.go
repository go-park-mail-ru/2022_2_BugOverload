package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

// AuthService provides universal service for work with users.
//
//go:generate mockgen -destination=../mocks/mock_auth_service.go -package=mock go-park-mail-ru/2022_2_BugOverload/internal/auth/service AuthService
type AuthService interface {
	Auth(ctx context.Context, user *models.User) (models.User, error)
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

func (u *authService) Auth(ctx context.Context, user *models.User) (models.User, error) {
	userRepo, err := u.authRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Auth")
	}
	return userRepo, nil
}

// Login is the service that accesses the interface AuthRepository.
// Validation: request password and password user from repository equal.
func (u *authService) Login(ctx context.Context, user *models.User) (models.User, error) {
	if err := ValidateEmail(user.Email); err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}
	if err := ValidatePassword(user.Password); err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}

	userRepo, err := u.authRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}

	if !security.IsPasswordsEqual(userRepo.Password, user.Password) {
		return models.User{}, stdErrors.Wrap(errors.ErrIncorrectPassword, "Login")
	}

	return userRepo, nil
}

// Signup is the service that accesses the interface AuthRepository
func (u *authService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	if err := ValidateNickname(user.Email); err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}
	if err := ValidateEmail(user.Email); err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}
	if err := ValidatePassword(user.Password); err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}
	user.Password = hashedPassword

	newUser, err := u.authRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	return newUser, nil
}
