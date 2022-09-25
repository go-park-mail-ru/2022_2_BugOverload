package handlers

import (
	"encoding/json"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs/tmp_storage"
)

type HandlerSignup struct {
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

func NewHandlerSignup() *HandlerSignup {
	return &HandlerSignup{}
}

//  UserStorage is tmp simple impl similar DB

var Storage tmp_storage.UserStorage = tmp_storage.NewUserStorage()

func (h *HandlerSignup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	//  There must be DataBase and Business logic magic
	suchUserExist := Storage.Insert(user)
	if suchUserExist != nil {
		w.WriteHeader(http.StatusOK) //  A user with such a mail already exists
		return
	} //  There must be DataBase and Business logic magic

	out, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  Отдаем ответ
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	w.Write(out)

	//  Логируем ответ
}

//  Подсказка для тестов
//  curl -vvv -X POST -H "Content-Type: application/json" -d '{"key": 123}' http://localhost:8086/v1/auth/signup
