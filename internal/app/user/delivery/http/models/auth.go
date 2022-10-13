package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type UserAuthRequest struct {
	user models.User
}

func (uar *UserAuthRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}

func (uar *UserAuthRequest) GetUser() *models.User {
	return &uar.user
}

func (uar *UserAuthRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}
