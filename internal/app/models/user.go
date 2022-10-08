package models

import (
	"encoding/json"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"io"
	"net/http"

	"github.com/wonderivan/logger"
)

// User is a carrier structure for all movie attributes and specifying them for json conversion
type User struct {
	ID       uint   `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// Bind is method for validation and create a data structure from json for processing
func (u *User) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors2.NewErrValidation(errors2.ErrContentTypeUndefined)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		return errors2.NewErrValidation(errors2.ErrUnsupportedMediaType)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	err = json.Unmarshal(body, u)
	if err != nil {
		return errors2.NewErrValidation(errors2.ErrCJSONUnexpectedEnd)
	}

	return nil
}
