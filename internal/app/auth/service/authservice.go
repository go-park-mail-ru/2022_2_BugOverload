package service

import (
	"context"
	"time"

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

func (a authService) CreateSession(ctx context.Context, user *models.User) (string, error) {
	newSession, err := a.authRepo.CreateSession(ctx, user)
	if err != nil {
		return "", err
	}

	return newSession, nil
}

func (a authService) GetSession(ctx context.Context) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authService) DeleteSession(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}
