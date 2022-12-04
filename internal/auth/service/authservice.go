package service

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/auth"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

//go:generate mockgen -source authservice.go -destination mocks/mockauthservice.go -package mockAuthService

// AuthService provides universal service for work with users.
type AuthService interface {
	Auth(ctx context.Context, user *models.User) (models.User, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
	GetAccess(ctx context.Context, user *models.User, userPassword string) error
	UpdatePassword(ctx context.Context, user *models.User, userPassword, userNewPassword string) error
}

// authService is implementation for users service corresponding to the AuthService interface.
type authService struct {
	authRepo auth.Repository
}

// NewAuthService is constructor for authService. Accepts Repository interfaces.
func NewAuthService(ur auth.Repository) AuthService {
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

// Login is the service that accesses the interface Repository.
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

// Signup is the service that accesses the interface Repository
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

	exist, err := u.authRepo.CheckExistUserByEmail(ctx, user.Email)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}
	if exist {
		return models.User{}, errors.ErrUserExist
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}
	user.Password = hashedPassword

	userDB, err := u.authRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	user.Password = ""
	user.Avatar = innerPKG.DefUserAvatar
	user.ID = userDB.ID

	return *user, nil
}

func (u *authService) GetAccess(ctx context.Context, user *models.User, userPassword string) error {
	userRepo, err := u.authRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return stdErrors.Wrap(err, "GetAccess")
	}

	if !security.IsPasswordsEqual(userRepo.Password, userPassword) {
		return stdErrors.Wrap(errors.ErrIncorrectPassword, "GetAccess")
	}

	return nil
}

func (u *authService) UpdatePassword(ctx context.Context, user *models.User, userPassword, userNewPassword string) error {
	errAccess := u.GetAccess(ctx, user, userPassword)
	if errAccess != nil {
		return stdErrors.Wrap(errAccess, "UpdatePassword")
	}

	newUserPassword, err := security.HashPassword(userNewPassword)
	if err != nil {
		return stdErrors.Wrap(err, "UpdatePassword")
	}

	updateErr := u.authRepo.UpdatePassword(ctx, user, newUserPassword)
	if updateErr != nil {
		return stdErrors.Wrap(updateErr, "UpdatePassword")
	}

	return nil
}
