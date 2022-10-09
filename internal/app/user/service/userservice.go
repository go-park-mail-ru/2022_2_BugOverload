package service

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"time"
)

type authService struct {
	articleRepo    interfaces.UserRepository
	authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}
