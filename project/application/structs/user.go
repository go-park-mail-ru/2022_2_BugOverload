package structs

import (
	"encoding/json"
	"errors"
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

const (
	AUTH_REQUEST_URI   = "/v1/auth"
	SIGNUP_REQUEST_URI = AUTH_REQUEST_URI + "/signup"
	LOGIN_REQUEST_URI  = AUTH_REQUEST_URI + "/login"
)

func validate(r *http.Request, user *User) error {
	requestURI := r.RequestURI

	switch requestURI {
	case AUTH_REQUEST_URI:
		return nil
	case LOGIN_REQUEST_URI:
		if user.Email != "" && user.Password != "" {
			return nil
		} else {
			return errors.New("request has empty fields (email | password)")
		}
	case SIGNUP_REQUEST_URI:
		if user.Email != "" {
			return nil
		} else {
			return errors.New("request has empty fields (nickname | email | password)")
		}
	default:
		return errors.New("invalid uri")
	}
}

func (u *User) ToPublic(r *http.Request) User {
	requestURI := r.RequestURI

	switch requestURI {
	case AUTH_REQUEST_URI:
		return User{
			Email:    u.Email,
			Nickname: u.Nickname,
			Password: u.Password,
		}
	case LOGIN_REQUEST_URI:
		return User{
			Nickname: u.Nickname,
			Email:    u.Email,
			Avatar:   u.Avatar,
		}
	case SIGNUP_REQUEST_URI:
		return User{
			Nickname: u.Nickname,
			Email:    u.Email,
		}
	default:
		return User{}
	}
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

	err = validate(r, u)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
