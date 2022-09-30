package auth

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserSignupRequest is empty struct with methods for signup handler
type UserSignupRequest struct {
	user structs.User
}

// Bind is func for validation and bind request fields to User struct for signup request
func (usr *UserSignupRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	err := usr.user.Bind(w, r)
	if err != nil {
		return err
	}

	if usr.user.Nickname == "" || usr.user.Email == "" || usr.user.Password == "" {
		return errorshandlers.ErrEmptyFieldAuth
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (usr *UserSignupRequest) GetUser() *structs.User {
	return &usr.user
}

// ToPublic return fields required by API
func (usr *UserSignupRequest) ToPublic(u *structs.User) structs.User {
	return structs.User{
		Email:    u.Email,
		Nickname: u.Nickname,
	}
}

// Signup is handling request
func (ha *HandlerAuth) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest UserSignupRequest

	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	user := signupRequest.GetUser()

	suchUserExist := ha.userStorage.CheckExist(user.Email)
	if suchUserExist {
		httpwrapper.DefHandlerError(w, errorshandlers.ErrSignupUserExist)
		return
	}

	ha.userStorage.Create(*user)

	newCookie := ha.cookieStorage.Create(user.Email)

	w.Header().Set("Cookie", newCookie)

	httpwrapper.Success(w, http.StatusCreated, signupRequest.ToPublic(user))
}