package models

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserLoginRequest struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Email    string `json:"email,omitempty" example:"YasaPupkinEzji@top.world"`
	Password string `json:"password,omitempty" example:"Widget Adapter"`
}

func NewUserLoginRequest() *UserLoginRequest {
	return &UserLoginRequest{}
}

func (u *UserLoginRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.NewErrValidation(errors.ErrContentTypeUndefined)
	}

	if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrBadBodyRequest
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if len(body) == 0 {
		return errors.NewErrValidation(errors.ErrEmptyBody)
	}

	err = json.Unmarshal(body, u)
	if err != nil {
		return errors.NewErrValidation(errors.ErrJSONUnexpectedEnd)
	}

	if (u.Nickname == "" && u.Email == "") || u.Password == "" {
		return errors.NewErrAuth(errors.ErrEmptyFieldAuth)
	}

	return nil
}

func (u *UserLoginRequest) GetUser() *models.User {
	return &models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Password: u.Password,
	}
}

type UserLoginResponse struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserLoginResponse(user *models.User) *UserLoginResponse {
	return &UserLoginResponse{
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Profile.Avatar,
	}
}
