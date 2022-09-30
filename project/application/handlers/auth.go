package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
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

// Login is handling request
func (ha *HandlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest structs.UserLoginRequest
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

	httpwrapper.ResponseOK(w, http.StatusOK, loginRequest.ToPublic(&userFromDB))
}

// Signup is handling request
func (ha *HandlerAuth) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest structs.UserSignupRequest
	err := signupRequest.Bind(w, r)
	if err != nil {
		httpwrapper.DefHandlerError(w, err)
		return
	}

	user := signupRequest.GetUser()

	suchUserExist := ha.userStorage.CheckExist(user.Email)
	if suchUserExist != nil {
		http.Error(w, "A user with such a mail already exists", http.StatusBadRequest)
		return
	}
	ha.userStorage.Create(*user)

	httpwrapper.ResponseOK(w, http.StatusCreated, signupRequest.ToPublic(user))
}
