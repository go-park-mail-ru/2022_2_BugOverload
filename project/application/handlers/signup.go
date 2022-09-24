package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type HandlerSignup struct {
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

func (h *HandlerSignup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  Получить уникальный номер HTTP запроса
	//  requestID := GetNextRequestID()

	//  Логируем входящий HTTP запрос

	//  Достаем, валидируем параметры запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Bad Request:"+err.Error(), http.StatusBadRequest)
		return
	}

	user := &structs.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  DataBase and Business logic magic
	plug := structs.User{
		Nickname: "StepByyyy",
		Email:    "dop123@mail.ru",
		Avatar:   "*ссылка",
	}
	//  DataBase and Business logic magic

	responseJSON, err := json.Marshal(plug)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  Логируем ответ

	//  Отдаем ответ
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write(responseJSON)
}

//  Подсказка для тестов
//  curl -vvv -X POST -H "Content-Type: application/json" -d '{"key": 123}' http://localhost:8086/v1/auth/signup
