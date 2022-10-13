package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type UserSignupRequest struct {
	user models.User
}

func (usr *UserSignupRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := usr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if usr.user.Nickname == "" || usr.user.Email == "" || usr.user.Password == "" {
		return errors.NewErrAuth(errors.ErrEmptyFieldAuth)
	}

	return nil
}

func (usr *UserSignupRequest) GetUser() *models.User {
	return &usr.user
}

func (usr *UserSignupRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}
