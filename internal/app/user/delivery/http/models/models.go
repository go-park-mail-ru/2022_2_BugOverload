package models

import (
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

type UserLoginRequest struct {
	user models.User
}

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

func (ulr *UserLoginRequest) GetUser() *models.User {
	return &ulr.user
}

func (ulr *UserLoginRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

type UserAuthRequest struct {
	user models.User
}

func (uar *UserAuthRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors2.NewErrAuth(errors2.ErrNoCookie)
	}

	return nil
}

func (uar *UserAuthRequest) GetUser() *models.User {
	return &uar.user
}

func (uar *UserAuthRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

type UserLogoutRequest struct {
	user models.User
}

func (ulr *UserLogoutRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errors2.NewErrAuth(errors2.ErrNoCookie)
	}

	return nil
}

func (ulr *UserLogoutRequest) GetUser() *models.User {
	return &ulr.user
}

type UserSignupRequest struct {
	user models.User
}

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

func (usr *UserSignupRequest) GetUser() *models.User {
	return &usr.user
}

func (usr *UserSignupRequest) ToPublic(u *models.User) models.User {
	return models.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}
