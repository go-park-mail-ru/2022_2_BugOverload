package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserAuthRequest struct {
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

func NewUserAuthRequest() *UserAuthRequest {
	return &UserAuthRequest{}
}

func (u *UserAuthRequest) Bind(r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}

func (u *UserAuthRequest) ToPublic(user *models.User) UserAuthRequest {
	return UserAuthRequest{
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
}
