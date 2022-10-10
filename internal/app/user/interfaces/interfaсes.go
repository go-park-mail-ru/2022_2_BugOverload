package interfaces

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// UserService represent the article's usecases
type UserService interface {
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
}

type UserRepository interface {
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
}
