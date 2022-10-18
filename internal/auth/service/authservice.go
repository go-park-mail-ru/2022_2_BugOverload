package service

import (
	"context"

	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// AuthService provides universal service for authorization. Needed for stateful session pattern.
type AuthService interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}

// authService is implementation for auth service corresponding to the AuthService interface.
type authService struct {
	authRepo repository.AuthRepository
}

// NewAuthService is constructor for authService. Accepts AuthRepository interfaces.
func NewAuthService(ar repository.AuthRepository) AuthService {
	return &authService{
		authRepo: ar,
	}
}

// GetUserBySession is the service that accesses the interface AuthRepository
func (a *authService) GetUserBySession(ctx context.Context) (models.User, error) {
	user, err := a.authRepo.GetUserBySession(ctx)
	if err != nil {
		return models.User{}, errors.Wrap(err, "GetUserBySession")
	}

	return user, nil
}

// CreateSession is the service that accesses the interface AuthRepository
func (a *authService) CreateSession(ctx context.Context, user *models.User) (string, error) {
	newSession, err := a.authRepo.CreateSession(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "CreateSession")
	}

	return newSession, nil
}

// GetSession is the service that accesses the interface AuthRepository
func (a *authService) GetSession(ctx context.Context) (string, error) {
	session, err := a.authRepo.GetSession(ctx)
	if err != nil {
		return "", errors.Wrap(err, "GetSession")
	}

	return session, nil
}

// DeleteSession is the service that accesses the interface AuthRepository
func (a *authService) DeleteSession(ctx context.Context) (string, error) {
	delSession, err := a.authRepo.DeleteSession(ctx)
	if err != nil {
		return "", errors.Wrap(err, "DeleteSession")
	}

	return delSession, nil
}
