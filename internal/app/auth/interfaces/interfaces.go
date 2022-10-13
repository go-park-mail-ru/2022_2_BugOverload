package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// AuthService provides universal service for authorization. Needed for stateful session pattern.
type AuthService interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}

// AuthRepository provides the versatility of session-related repositories.
// Needed to work with stateful session pattern.
type AuthRepository interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}
