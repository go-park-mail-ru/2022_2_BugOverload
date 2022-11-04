package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
)

// UserService provides universal service for work with users.
type UserService interface {
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
}

// userService is implementation for users service corresponding to the UserService interface.
type userService struct {
	userProfileRepo repository.UserRepository
}

// NewUserProfileService is constructor for userService. Accepts UserService interfaces.
func NewUserProfileService(ur repository.UserRepository) UserService {
	return &userService{
		userProfileRepo: ur,
	}
}

// GetUserProfileByID is the service that accesses the interface UserService.
func (u *userService) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	userRepo, err := u.userProfileRepo.GetUserProfileByID(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "GetPersonByID")
	}

	return userRepo, nil
}

// GetUserProfileSettings is the service that accesses the interface UserService
func (u *userService) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userProfileRepo.GetUserProfileSettings(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "GetUserProfileSettings")
	}

	return newUser, nil
}

// ChangeUserProfileSettings is the service that accesses the interface UserService
func (u *userService) ChangeUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userProfileRepo.GetUserProfileSettings(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "ChangeUserProfileSettings")
	}

	return newUser, nil
}
