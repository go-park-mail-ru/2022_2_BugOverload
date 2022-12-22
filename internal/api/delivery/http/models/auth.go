package models

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

//go:generate easyjson -disallow_unknown_fields auth.go

//easyjson:json
type UserAuthResponse struct {
	Nickname string `json:"nickname,omitempty" example:"Bot373"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserAuthResponse(user *models.User) *UserAuthResponse {
	return &UserAuthResponse{
		Email:    user.Email,
		Nickname: security.Sanitize(user.Nickname),
		Avatar:   user.Avatar,
	}
}
