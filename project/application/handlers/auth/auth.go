package auth

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// HandlerAuth is structure for API auth, login and signup processing
type HandlerAuth struct {
	userStorage   *database.UserStorage
	cookieStorage *database.CookieStorage
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

// NewHandlerAuth is constructor for HandlerAuth
func NewHandlerAuth(us *database.UserStorage, cs *database.CookieStorage) *HandlerAuth {
	return &HandlerAuth{us, cs}
}

// UserAuthRequest is empty struct with methods for login handler
type UserAuthRequest struct {
	user structs.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (uar *UserAuthRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errorshandlers.ErrNoCookie
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (uar *UserAuthRequest) GetUser() *structs.User {
	return &uar.user
}

// ToPublic return fields required by API
func (uar *UserAuthRequest) ToPublic(u *structs.User) structs.User {
	return structs.User{
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

// Auth is handling request for check current client cookie and return user data
func (ha *HandlerAuth) Auth(w http.ResponseWriter, r *http.Request) {
	var authRequest UserAuthRequest

	err := authRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	cookie, err := ha.cookieStorage.GetCookie(cookieStr)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	userFromDB, err := ha.userStorage.GetUser(cookie.Value)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	httpwrapper.Response(w, http.StatusOK, authRequest.ToPublic(&userFromDB))
}
