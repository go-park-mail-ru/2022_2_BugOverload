package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	authInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type authService struct {
	authRepo       authInterface.AuthRepository
	contextTimeout time.Duration
}

func NewAuthService(ar authInterface.AuthRepository, timeout time.Duration) authInterface.AuthService {
	return &authService{
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
		return "", errors.Wrap(err, "CreateSession")
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
	delSession, err := a.authRepo.DeleteSession(ctx)
	if err != nil {
		return "", errors.Wrap(err, "DeleteSession")
	}

	return delSession, nil
}
