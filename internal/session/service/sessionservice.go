package service

import (
	"context"

	"github.com/pkg/errors"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
)

// SessionService provides universal service for authorization. Needed for stateful session pattern.
type SessionService interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (models.Session, error)
	DeleteSession(ctx context.Context) (models.Session, error)
}

// sessionService is implementation for auth service corresponding to the SessionService interface.
type sessionService struct {
	sessionRepo repository.SessionRepository
}

// NewSessionService is constructor for sessionService. Accepts SessionRepository interfaces.
func NewSessionService(ar repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepo: ar,
	}
}

// GetUserBySession is the service that accesses the interface SessionRepository
func (a *sessionService) GetUserBySession(ctx context.Context) (models.User, error) {
	user, err := a.sessionRepo.GetUserBySession(ctx)
	if err != nil {
		return models.User{}, errors.Wrap(err, "GetUserBySession")
	}

	return user, nil
}

// CreateSession is the service that accesses the interface SessionRepository
func (a *sessionService) CreateSession(ctx context.Context, user *models.User) (models.Session, error) {
	newSession, err := a.sessionRepo.CreateSession(ctx, user)
	if err != nil {
		return models.Session{}, errors.Wrap(err, "CreateSession")
	}

	return newSession, nil
}

// DeleteSession is the service that accesses the interface SessionRepository
func (a *sessionService) DeleteSession(ctx context.Context) (models.Session, error) {
	delSession, err := a.sessionRepo.DeleteSession(ctx)
	if err != nil {
		return models.Session{}, errors.Wrap(err, "DeleteSession")
	}

	return delSession, nil
}
