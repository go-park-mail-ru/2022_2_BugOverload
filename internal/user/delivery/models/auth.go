package models

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserAuthRequest struct{}

func NewUserAuthRequest() *UserAuthRequest {
	return &UserAuthRequest{}
}

func (u *UserAuthRequest) Bind(r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors.NewErrAuth(errors.ErrNoCookie)
	}

	return nil
}

type UserAuthResponse struct {
	Nickname string `json:"nickname,omitempty" example:"Bot373"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ссылка}}"`
}

func NewUserAuthResponse() *UserAuthResponse {
	return &UserAuthResponse{}
}

func (u *UserAuthResponse) ToPublic(user *models.User) UserAuthResponse {
	return UserAuthResponse{
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
}
