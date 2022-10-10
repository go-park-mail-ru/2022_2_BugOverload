package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
)

type authService struct {
	userRepo       userInterface.UserRepository
	authRepo       authInterface.AuthRepository
	contextTimeout time.Duration
}

func NewAuthService(ur userInterface.UserRepository, ar authInterface.AuthRepository, timeout time.Duration) authInterface.AuthService {
	return &authService{
		userRepo:       ur,
		authRepo:       ar,
		contextTimeout: timeout,
	}
}

func (a authService) GetUserBySession(ctx context.Context) (models.User, error) {
	user, err := a.authRepo.GetUserBySession(ctx)
	if err != nil {
		return models.User{}, errors.Wrap(err, "GetUserBySession")
	}

	return user, nil
}

func (a authService) CreateSession(ctx context.Context, user *models.User) (string, error) {
	newSession, err := a.authRepo.CreateSession(ctx, user)
	if err != nil {
		return "", err
	}

	return newSession, nil
}

func (a authService) GetSession(ctx context.Context) (string, error) {
	session, err := a.authRepo.GetSession(ctx)
	if err != nil {
		return "", errors.Wrap(err, "GetSession")
	}

	return session, nil
}

func (a authService) DeleteSession(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}