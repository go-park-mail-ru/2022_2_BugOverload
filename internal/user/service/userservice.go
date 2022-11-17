package service

import (
	"context"

	stdErrors "github.com/pkg/errors"

	authService "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
)

// UserService provides universal service for work with users.
type UserService interface {
	GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error)
	GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error)
	ChangeUserProfileSettings(ctx context.Context, user *models.User, params *innerPKG.ChangeUserSettings) error

	FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error)
	FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateDropParams) (models.Film, error)

	NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *innerPKG.NewFilmReviewParams) error

	GetUserActivityOnFilm(ctx context.Context, user *models.User, params *innerPKG.GetUserActivityOnFilmParams) (models.UserActivity, error)
}

// userService is implementation for users service corresponding to the UserService interface.
type userService struct {
	userRepo    repository.UserRepository
	authService authService.AuthService
}

// NewUserProfileService is constructor for userService. Accepts UserService interfaces.
func NewUserProfileService(ur repository.UserRepository, as authService.AuthService) UserService {
	return &userService{
		userRepo:    ur,
		authService: as,
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

	err := u.authService.UpdatePassword(ctx, user, params.CurPassword, params.NewPassword)
	if err != nil {
		return stdErrors.Wrap(err, "ChangeUserProfileNickname")
	}

	return nil
}

// FilmRate is the service that accesses the interface UserService
func (u *userService) FilmRate(ctx context.Context, user *models.User, params *innerPKG.FilmRateParams) (models.Film, error) {
	if !(0 <= params.Score && params.Score <= 10) {
		return models.Film{}, stdErrors.Wrap(errors.ErrBadRequestParams, "FilmRateSet")
	}

	var film models.Film

	exist, err := u.userRepo.FilmRatingExist(ctx, user, params.FilmID)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "FilmRateSet")
	}

	if exist {
		film, err = u.userRepo.FilmRateUpdate(ctx, user, params)
		if err != nil {
			return models.Film{}, stdErrors.Wrap(err, "FilmRateSet")
		}
	} else {
		film, err = u.userRepo.FilmRateSet(ctx, user, params)
		if err != nil {
			return models.Film{}, stdErrors.Wrap(err, "FilmRateSet")
		}
	}

	return film, nil
}

// FilmRateDrop is the service that accesses the interface UserService
func (u *userService) FilmRateDrop(ctx context.Context, user *models.User, params *innerPKG.FilmRateDropParams) (models.Film, error) {
	exist, err := u.userRepo.FilmRatingExist(ctx, user, params.FilmID)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "FilmRateDrop")
	}
	if !exist {
		return models.Film{}, stdErrors.Wrap(errors.ErrFilmRatingNotExist, "FilmRateDrop")
	}

	film, err := u.userRepo.FilmRateDrop(ctx, user, params)
	if err != nil {
		return models.Film{}, stdErrors.Wrap(err, "FilmRateDrop")
	}

	return film, nil
}

// NewFilmReview is the service that accesses the interface UserService
func (u *userService) NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *innerPKG.NewFilmReviewParams) error {
	err := u.userRepo.NewFilmReview(ctx, user, review, params)
	if err != nil {
		return stdErrors.Wrap(err, "NewFilmReview")
	}

	return nil
}

// GetUserActivityOnFilm is the service that accesses the interface UserService
func (u *userService) GetUserActivityOnFilm(ctx context.Context, user *models.User, params *innerPKG.GetUserActivityOnFilmParams) (models.UserActivity, error) {
	userActivity, err := u.userRepo.GetUserActivityOnFilm(ctx, user, params)
	if err != nil {
		return models.UserActivity{}, stdErrors.Wrap(err, "GetUserActivityOnFilm")
	}

	return userActivity, nil
}
