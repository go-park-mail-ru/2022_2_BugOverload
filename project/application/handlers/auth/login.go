package auth

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
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
		return errorshandlers.ErrEmptyFieldAuth
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
		httpwrapper.DefHandlerError(w, err)
		return
	}

	user := loginRequest.GetUser()

	userFromDB, err := ha.userStorage.GetUser(user.Email)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	if userFromDB.Password != user.Password {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrLoginCombinationNotFound)
		return
	}

	ha.cookieStorage.Create(user.Email)

	newCookie, err := ha.cookieStorage.GetCookie(user.Email)
	if err != nil {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrCookieNotExist)
		return
	}

	http.SetCookie(w, &newCookie)

	httpwrapper.ResponseOK(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}
