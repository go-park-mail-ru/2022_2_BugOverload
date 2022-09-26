package structs

import (
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	"io"
	"net/http"
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
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return errors.New("Content-Type must be application/json")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request:"+err.Error(), http.StatusBadRequest)
		return err
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	err = json.Unmarshal(body, u)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
