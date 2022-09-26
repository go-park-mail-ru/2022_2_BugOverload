package handlers

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// HandlerSignup is structure for API registration processing
type HandlerSignup struct {
	storage *database.UserStorage //  UserStorage is tmp simple impl similar DB
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

// NewHandlerSignup is constructor for HandlerSignup
func NewHandlerSignup(us *database.UserStorage) *HandlerSignup {
	return &HandlerSignup{us}
}

// Check empty fields in request
func signupRequiredFields(user *structs.User) bool {
	return user.Nickname != "" && user.Email != "" && user.Password != ""
}

// Signup is handling request
func (h *HandlerSignup) Signup(w http.ResponseWriter, r *http.Request) {
	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	// Fields validation
	if !signupRequiredFields(&user) {
		http.Error(w, "Request has empty fields (Nickname | Email | Password)", http.StatusBadRequest)
		return
	}

	//  There must be DataBase and Business logic magic
	suchUserExist := h.storage.CheckExist(user.Email)
	if suchUserExist != nil {
		http.Error(w, "A user with such a mail already exists", http.StatusBadRequest)
		return
	}
	h.storage.Create(user)
	//  There must be DataBase and Business logic magic

	httpwrapper.ResponseOK(w, http.StatusCreated, user)

	//  Логируем ответ
}

//  Подсказка для тестов
//  curl -vvv -X POST -H "Content-Type: application/json" -d '{"nickname":"nick", "email":"mail@mail.ru", "password":"pass"}' http://localhost:8088/v1/auth/signup
