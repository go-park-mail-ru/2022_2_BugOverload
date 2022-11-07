package models

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

type UserPutSettingsRequest struct {
	Nickname    string `json:"nickname,omitempty" example:"StepByyyy"`
	NewPassword string `json:"new_password,omitempty" example:"Widget Adapter"`
}

func NewPutUserSettingsRequest() *UserPutSettingsRequest {
	return &UserPutSettingsRequest{}
}

func (u *UserPutSettingsRequest) Bind(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.ErrContentTypeUndefined
	}

	if r.Header.Get("Content-Type") != pkg.ContentTypeJSON {
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

	err = json.Unmarshal(body, u)
	if err != nil {
		return errors.ErrJSONUnexpectedEnd
	}

	return nil
}

func (u *UserPutSettingsRequest) GetParams() *pkg.ChangeUserSettings {
	return &pkg.ChangeUserSettings{
		NewPassword: u.NewPassword,
	}
}
