package handlers

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
	"net/http"
)

// HandlerLogin is structure for API login processing
type HandlerLogin struct {
	storage *database.UserStorage //  UserStorage is tmp simple impl similar DB
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

// NewHandlerSignup is constructor for HandlerSignup
func NewHandlerLogin(us *database.UserStorage) *HandlerLogin {
	return &HandlerLogin{us}
}

// Check empty fields in request
func loginRequiredFields(user *structs.User) bool {
	return user.Email != "" && user.Password != ""
}

// Login is handling request
func (h *HandlerLogin) Login(w http.ResponseWriter, r *http.Request) {
	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	// Fields validation
	if !loginRequiredFields(&user) {
		http.Error(w, "Request has empty fields (Email | Password)", http.StatusBadRequest)
		return
	}

	//  There must be DataBase and Business logic magic
	userFromDB, err := h.storage.GetUser(user.Email)
	if err != nil || user.Password != userFromDB.Password {
		http.Error(w, "No such combination of user and password", http.StatusBadRequest)
		return
	}

	// Return only required API fields
	httpwrapper.ResponseOK(w, http.StatusOK, structs.User{
		Nickname: userFromDB.Nickname,
		Email:    userFromDB.Email,
		Avatar:   userFromDB.Avatar,
	})

	//  Логируем ответ
}

//  Подсказка для тестов
//  curl -vvv -X POST -H "Content-Type: application/json" -d '{"email":"test@mail.ru", "password":"pass"}' http://localhost:8088/v1/auth/login
