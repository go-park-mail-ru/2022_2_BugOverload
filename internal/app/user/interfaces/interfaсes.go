package interfaces

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// Example
//type ArticleUsecase interface {
//	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
//	GetByID(ctx context.Context, id int64) (Article, error)
//	Update(ctx context.Context, ar *Article) error
//	GetByTitle(ctx context.Context, title string) (Article, error)
//	Store(context.Context, *Article) error
//	Delete(ctx context.Context, id int64) error
//}

// Example
//type ArticleRepository interface {
//	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
//	GetByID(ctx context.Context, id int64) (Article, error)
//	GetByTitle(ctx context.Context, title string) (Article, error)
//	Update(ctx context.Context, ar *Article) error
//	Store(ctx context.Context, a *Article) error
//	Delete(ctx context.Context, id int64) error
//}

// AuthService represent the article's usecases
type AuthService interface {
	Auth(ctx context.Context) (models.User, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
	Logout(ctx context.Context) error
}

type AuthRepository interface {
	Auth(ctx context.Context) (models.User, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
	Signup(ctx context.Context, user *models.User) (models.User, error)
	Logout(ctx context.Context) error
}
