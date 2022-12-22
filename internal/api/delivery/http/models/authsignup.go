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

//go:generate easyjson  -disallow_unknown_fields authsignup.go

//easyjson:json
type UserSignupRequest struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Email    string `json:"email,omitempty" example:"YasaPupkinEzji@top.world"`
	Password string `json:"password,omitempty" example:"Widget Adapter"`
}

func NewUserSignupRequest() *UserSignupRequest {
	return &UserSignupRequest{}
}

func (u *UserSignupRequest) Bind(r *http.Request) error {
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

	if u.Password == "" || u.Email == "" || u.Nickname == "" {
		return errors.ErrBadRequestParamsEmptyRequiredFields
	}

	return nil
}

func (u *UserSignupRequest) GetUser() *models.User {
	return &models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Password: u.Password,
	}
}

//easyjson:json
type UserSignupResponse struct {
	Nickname string `json:"nickname,omitempty" example:"StepByyyy"`
	Email    string `json:"email,omitempty" example:"dop123@mail.ru"`
	Avatar   string `json:"avatar,omitempty" example:"{{ключ}}"`
}

func NewUserSignUpResponse(user *models.User) *UserSignupResponse {
	return &UserSignupResponse{
		Email:    user.Email,
		Nickname: security.Sanitize(user.Nickname),
		Avatar:   user.Avatar,
	}
}
