package handlers

import (
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/httpwrapper"
	"go-park-mail-ru/2022_2_BugOverload/project/application/storages"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type HandlerSignup struct {
	storage *storages.UserStorage //  UserStorage is tmp simple impl similar DB
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

func NewHandlerSignup(us *storages.UserStorage) *HandlerSignup {
	return &HandlerSignup{us}
}

func (h *HandlerSignup) Some(w http.ResponseWriter, r *http.Request) {
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
