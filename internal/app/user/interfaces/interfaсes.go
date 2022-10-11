package interfaces

import (
	"context"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type UserService interface {
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, user *models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
}
