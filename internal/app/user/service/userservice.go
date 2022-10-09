package service

import (
	"context"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
)

type userService struct {
	userRepo       userInterface.UserRepository
	contextTimeout time.Duration
}

func NewUserService(ur userInterface.UserRepository, timeout time.Duration) userInterface.UserService {
	return &userService{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

func (u userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userRepo.Signup(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

//suchUserExist := h.userStorage.CheckExist(user.Email)
//if suchUserExist {
//	httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrSignupUserExist))
//	return
//}
//
//h.userStorage.CreateSession(*user)
//
//newCookie := h.cookieStorage.CreateSession(user.Email)
