package auth

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserLoginRequest is empty struct with methods for login handler
type UserLoginRequest struct {
	user structs.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (ulr *UserLoginRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := ulr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if (ulr.user.Nickname == "" && ulr.user.Email == "") || ulr.user.Password == "" {
		return errors.ErrEmptyFieldAuth
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (ulr *UserLoginRequest) GetUser() *structs.User {
	return &ulr.user
}

// ToPublic return fields required by API
func (ulr *UserLoginRequest) ToPublic(u *structs.User) structs.User {
	return structs.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

// Login is handling request
func (ha *HandlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest UserLoginRequest

	err := loginRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userFromDB, err := ha.userStorage.GetUser(user.Email)
	if err != nil {
		httpwrapper.DefaultHandlerError(w, err)

		return
	}

	if userFromDB.Password != user.Password {
		httpwrapper.DefaultHandlerError(w, errors.ErrLoginCombinationNotFound)

		return
	}

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Set-Cookie", newCookie)

	httpwrapper.Response(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}
