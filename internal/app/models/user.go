package models

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type User struct {
	ID       uint   `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

func (u *User) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.NewErrValidation(errors.ErrContentTypeUndefined)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		return errors.NewErrValidation(errors.ErrUnsupportedMediaType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	err = json.Unmarshal(body, u)
	if err != nil {
		return errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd)
	}

	return nil
}
