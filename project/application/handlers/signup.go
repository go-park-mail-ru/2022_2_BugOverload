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
func New_HandlerSignup(us *database.UserStorage) *HandlerSignup {
	return &HandlerSignup{us}
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

	//  There must be DataBase and Business logic magic
	suchUserExist := h.storage.CheckExist(user)
	if suchUserExist != nil {
		http.Error(w, "A user with such a mail already exists", http.StatusBadRequest)
		return
	}
	h.storage.Create(user)
	//  There must be DataBase and Business logic magic

	httpwrapper.Created(w, user)

	//  Логируем ответ
}

//  Подсказка для тестов
//  curl -vvv -X POST -H "Content-Type: application/json" -d '{"key": 123}' http://localhost:8086/v1/auth/signup
