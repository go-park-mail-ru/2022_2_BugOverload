package models

import (
	"context"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserAuthRequest struct{}

func NewUserAuthRequest() *UserAuthRequest {
	return &UserAuthRequest{}
}

func (u *UserAuthRequest) Bind(r *http.Request) (context.Context, error) {
	if r.Header.Get("Cookie") == "" {
		return nil, errors.NewErrAuth(errors.ErrNoCookie)
	}

	cookie := r.Cookies()[0]

	ctx := context.WithValue(r.Context(), pkg.SessionKey, cookie.Value)

	return ctx, nil
}

type UserAuthResponse struct {
	Nickname string `json:"nickname,omitempty" example:"Bot373"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserAuthResponse(user *models.User) *UserAuthResponse {
	return &UserAuthResponse{
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Profile.Avatar,
	}
}

func (u *UserAuthResponse) ToPublic() *UserAuthResponse {
	return u
}
