package models

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

//go:generate easyjson  -disallow_unknown_fields authlogin.go

//easyjson:json
type UserLoginRequest struct {
	Email    string `json:"email,omitempty" example:"YasaPupkinEzji@top.world"`
	Password string `json:"password,omitempty" example:"Widget Adapter"`
}

func NewUserLoginRequest() *UserLoginRequest {
	return &UserLoginRequest{}
}

func (u *UserLoginRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != constparams.ContentTypeJSON {
		return errors.ErrUnsupportedMediaType
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
		return errors.ErrEmptyBody
	}

	err = easyjson.Unmarshal(body, u)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	if u.Password == "" || u.Email == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (u *UserLoginRequest) GetUser() *models.User {
	return &models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

//easyjson:json
type UserLoginResponse struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserLoginResponse(user *models.User) *UserLoginResponse {
	return &UserLoginResponse{
		Email:    user.Email,
		Nickname: security.Sanitize(user.Nickname),
		Avatar:   user.Avatar,
	}
}
