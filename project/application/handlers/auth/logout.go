package auth

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// UserLogoutRequest is empty struct with methods for login handler
type UserLogoutRequest struct {
	user structs.User
}

// Bind is func for validation and bind request fields to User struct for login request
func (ulr *UserLogoutRequest) Bind(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("Cookie") == "" {
		return errorshandlers.ErrNoCookie
	}

	return nil
}

// GetUser is func for parse user fields and create struct User
func (ulr *UserLogoutRequest) GetUser() *structs.User {
	return &ulr.user
}

// Logout is handling request for check current client cookie and return user data
func (ha *HandlerAuth) Logout(w http.ResponseWriter, r *http.Request) {
	var authRequest UserAuthRequest

	err := authRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	cookieStr := r.Header.Get("Cookie")

	badCookie, err := ha.cookieStorage.DeleteCookie(cookieStr)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	w.Header().Set("Set-Cookie", badCookie)

	httpwrapper.Response(w, http.StatusOK, "")
}
