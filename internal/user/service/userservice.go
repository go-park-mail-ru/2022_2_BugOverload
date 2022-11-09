package service

import (
	"context"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
)

// UserService provides universal service for work with users.
type UserService interface {
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User, params *innerPKG.ChangeUserSettings) error

	FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error
	FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error
}

// userService is implementation for users service corresponding to the UserService interface.
type userService struct {
	userRepo repository.UserRepository
}

// NewUserProfileService is constructor for userService. Accepts UserService interfaces.
func NewUserProfileService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

// GetUserProfileByID is the service that accesses the interface UserService.
func (u *userService) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	userRepo, err := u.userRepo.GetUserProfileByID(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "GetPersonByID")
	}

	return userRepo, nil
}

// GetUserProfileSettings is the service that accesses the interface UserService
func (u *userService) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userRepo.GetUserProfileSettings(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "GetUserProfileSettings")
	}

	return newUser, nil
}

// ChangeUserProfileSettings is the service that accesses the interface UserService
func (u *userService) ChangeUserProfileSettings(ctx context.Context, user *models.User, params *innerPKG.ChangeUserSettings) error {
	user.Nickname = params.Nickname

	if params.NewPassword == "" {
		err := u.userRepo.ChangeUserProfileNickname(ctx, user)
		if err != nil {
			return stdErrors.Wrap(err, "ChangeUserProfileNickname")
		}

		return nil
	}

	passwordDB, err := u.userRepo.GetPassword(ctx, user)
	if err != nil {
		return stdErrors.Wrap(err, "ChangeUserProfileNickname")
	}

	ok := security.IsPasswordsEqual(passwordDB, params.CurPassword)
	if !ok {
		return stdErrors.Wrap(err, "ChangeUserProfileNickname")
	}

	user.Password, err = security.HashPassword(params.NewPassword)
	if !ok {
		return stdErrors.Wrap(err, "ChangeUserProfileNickname")
	}

	err = u.userRepo.ChangeUserProfilePassword(ctx, user)
	if err != nil {
		return stdErrors.Wrap(err, "ChangeUserProfileNickname")
	}

	return nil
}

// FilmRate is the service that accesses the interface UserService
func (u *userService) FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error {
	err := u.userRepo.FilmRate(ctx, user, params)
	if err != nil {
		return stdErrors.Wrap(err, "FilmRate")
	}

	return nil
}

// FilmRateDrop is the service that accesses the interface UserService
func (u *userService) FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) error {
	err := u.userRepo.FilmRateDrop(ctx, user, params)
	if err != nil {
		return stdErrors.Wrap(err, "FilmRateDrop")
	}

	return nil
}
