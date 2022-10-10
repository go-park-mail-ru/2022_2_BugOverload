package service

import (
	"context"
	stdErrors "github.com/pkg/errors"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	userInterface "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
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
	userRepo, err := u.userRepo.GetUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Login")
	}

	if userRepo.Password != user.Password {
		return models.User{}, errors.ErrLoginCombinationNotFound
	}

	return userRepo, nil
}

func (u userService) Signup(ctx context.Context, user *models.User) (models.User, error) {
	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, stdErrors.Wrap(err, "Signup")
	}

	return newUser, nil
}

//  login
//userFromDB, err := h.userStorage.Login(user.Email)
//if err != nil {
//httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(err))
//return
//}
//
//if userFromDB.Password != user.Password {
//httpwrapper.DefaultHandlerError(w, errors.NewErrAuth(errors.ErrLoginCombinationNotFound))
//return
//return
//}
//
//newCookie := h.cookieStorage.CreateSession(user.Email)

//  logout
//badCookie, err := h.cookieStorage.DeleteCookie(cookieStr)
//if err != nil {
//httpwrapper2.DefaultHandlerError(w, errors.NewErrAuth(err))
//return
//}
//}