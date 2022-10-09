package models

import (
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// UserLoginRequest is empty struct with methods for login handler
type UserLoginRequest struct {
	user models.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (ulr *UserLoginRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := ulr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if (ulr.user.Nickname == "" && ulr.user.Email == "") || ulr.user.Password == "" {
		return errors2.NewErrAuth(errors2.ErrEmptyFieldAuth)
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (ulr *UserLoginRequest) GetUser() *models.User {
	return &ulr.user
}

// ToPublic return fields required by API
func (ulr *UserLoginRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

// UserAuthRequest is empty struct with methods for login handler
type UserAuthRequest struct {
	user models.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (uar *UserAuthRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors2.NewErrAuth(errors2.ErrNoCookie)
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (uar *UserAuthRequest) GetUser() *models.User {
	return &uar.user
}

// ToPublic return fields required by API
func (uar *UserAuthRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

// UserLogoutRequest is empty struct with methods for login handler
type UserLogoutRequest struct {
	user models.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (ulr *UserLogoutRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors2.NewErrAuth(errors2.ErrNoCookie)
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (ulr *UserLogoutRequest) GetUser() *models.User {
	return &ulr.user
}

// UserSignupRequest is empty struct with methods for signup handler
type UserSignupRequest struct {
	user models.User
}

// Bind is func for validation and bind request fields to User struct for signup request
func (usr *UserSignupRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := usr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if usr.user.Nickname == "" || usr.user.Email == "" || usr.user.Password == "" {
		return errors2.NewErrAuth(errors2.ErrEmptyFieldAuth)
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (usr *UserSignupRequest) GetUser() *models.User {
	return &usr.user
}

// ToPublic return fields required by API
func (usr *UserSignupRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}
