package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
)

type UserLoginRequest struct {
	user models.User
}

func (ulr *UserLoginRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := ulr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if (ulr.user.Nickname == "" && ulr.user.Email == "") || ulr.user.Password == "" {
		return errors.NewErrAuth(errors.ErrEmptyFieldAuth)
	}

	return nil
}

func (ulr *UserLoginRequest) GetUser() *models.User {
	return &ulr.user
}

func (ulr *UserLoginRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}
