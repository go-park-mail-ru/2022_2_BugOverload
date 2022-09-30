package structs

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
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
		return errorshandlers.ErrContentTypeUndefined
	}

	logger.Info(r.Header.Get("Content-Type"))

	if r.Header.Get("Content-Type") != "application/json" {
		return errorshandlers.ErrUnsupportedMediaType
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
		return err
	}

	return nil
}
