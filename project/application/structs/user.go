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

// Bind is method for validation and create a data structure from json for processing
func (u *User) Bind(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return http.StatusUnsupportedMediaType, errors.New("Unsupported Media Type")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	err = json.Unmarshal(body, u)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return 0, nil
}

// UserLoginRequest is empty struct with methods for login handler
type UserLoginRequest struct {
	user User
}

// Bind is func for validation and bind request fields to User struct for login request
func (loginRequest *UserLoginRequest) Bind(w http.ResponseWriter, r *http.Request) (int, error) {
	code, err := loginRequest.user.Bind(w, r)
	if err != nil {
		return code, err
	}

	if loginRequest.user.Email == "" && loginRequest.user.Password == "" {
		return http.StatusBadRequest, errors.New("request has empty fields (email | password)")
	}

	return 0, nil
}

// GetUser is func for parse user fields and create struct User
func (loginRequest *UserLoginRequest) GetUser() *User {
	return &loginRequest.user
}

// ToPublic return fields required by API
func (loginRequest *UserLoginRequest) ToPublic(u *User) User {
	return User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

// UserSignupRequest is empty struct with methods for signup handler
type UserSignupRequest struct {
	user User
}

// Bind is func for validation and bind request fields to User struct for signup request
func (signupRequest *UserSignupRequest) Bind(w http.ResponseWriter, r *http.Request) (int, error) {
	code, err := signupRequest.user.Bind(w, r)
	if err != nil {
		return code, err
	}

	if signupRequest.user.Nickname != "" && signupRequest.user.Email != "" && signupRequest.user.Password != "" {
		return http.StatusBadRequest, errors.New("request has empty fields (nickname | email | password)")
	}

	return 0, nil
}

// GetUser is func for parse user fields and create struct User
func (signupRequest *UserSignupRequest) GetUser() *User {
	return &signupRequest.user
}

// ToPublic return fields required by API
func (signupRequest *UserSignupRequest) ToPublic(u *User) User {
	return User{
		Email:    u.Email,
		Nickname: u.Nickname,
	}
}
