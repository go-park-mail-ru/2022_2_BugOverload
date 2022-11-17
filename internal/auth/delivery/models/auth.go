package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

type UserAuthResponse struct {
	ID       int    `json:"id,omitempty" example:"13"`
	Nickname string `json:"nickname,omitempty" example:"Bot373"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserAuthResponse(user *models.User) *UserAuthResponse {
	return &UserAuthResponse{
		ID:       user.ID,
		Email:    user.Email,
		Nickname: security.Sanitize(user.Nickname),
		Avatar:   user.Avatar,
	}
}
