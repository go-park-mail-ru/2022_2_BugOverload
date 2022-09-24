package handlers

import (
	"encoding/json"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type HandlerSignup struct {
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

func NewHandlerSignup() *HandlerSignup {
	return &HandlerSignup{}
}

func (h *HandlerSignup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  Получить уникальный номер HTTP запроса
	//  requestID := GetNextRequestID()

	//  Логируем входящий HTTP запрос

	// Достаем, валидируем и конвертруем параметры в объект
	var user structs.User
	err := user.Bind(w, r)
	if err != nil {
		return
	}

	//  DataBase and Business logic magic
	//  user -> handler
	//  Эхо сервер
	plug := user

	//suchUserExist := true
	//
	//if suchUserExist {
	//	w.WriteHeader(http.StatusOK)
	//  return
	//}

	//  handler -> plug
	//  DataBase and Business logic magic

	out, err := json.Marshal(plug)
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
