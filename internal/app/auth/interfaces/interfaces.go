package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// AuthService represent the auth service
type AuthService interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}

// AuthRepository represent the article's repository
type AuthRepository interface {
	GetUserBySession(ctx context.Context) (models.User, error)
	CreateSession(ctx context.Context, user *models.User) (string, error)
	GetSession(ctx context.Context) (string, error)
	DeleteSession(ctx context.Context) (string, error)
}