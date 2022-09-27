package handlers

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// HandlerAuth is structure for API auth, login and signup processing
type HandlerAuth struct {
	storage *database.UserStorage //  UserStorage is tmp simple impl similar DB
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

// NewHandlerAuth is constructor for HandlerAuth
func NewHandlerAuth(us *database.UserStorage) *HandlerAuth {
	return &HandlerAuth{us}
}

// Login is handling request
func (ha *HandlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	//  There must be DataBase and Business logic magic
	userFromDB, err := ha.storage.GetUser(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userFromDB.Password != user.Password {
		http.Error(w, "No such combination of user and password", http.StatusBadRequest)
		return
	}

	// Return only required API fields
	httpwrapper.ResponseOK(w, http.StatusOK, userFromDB.ToPublic(r))

	//  Логируем ответ
}

// Signup is handling request
func (ha *HandlerAuth) Signup(w http.ResponseWriter, r *http.Request) {
	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	//  There must be DataBase and Business logic magic
	suchUserExist := ha.storage.CheckExist(user.Email)
	if suchUserExist != nil {
		http.Error(w, "A user with such a mail already exists", http.StatusBadRequest)
		return
	}
	ha.storage.Create(user)
	//  There must be DataBase and Business logic magic

	httpwrapper.ResponseOK(w, http.StatusCreated, user.ToPublic(r))

	//  Логируем ответ
}
