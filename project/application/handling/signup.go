package handling

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"Kinopoisk/project/application/structs"
)

type HandlerSignup struct {
	user structs.User
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user := &structs.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//  Формируем ответ
	var responseBody io.ReadCloser
	responseBody.Read(responseJSON)

	responseHeaders := http.Header{}
	responseHeaders.Add("Content-Type", "application/json")

	var HTTPResponse = http.Response{
		StatusCode: http.StatusCreated,
		Header:     responseHeaders,
		Body:       responseBody,
	}

	//  Логируем ответ

	//  Отдаем ответ
}
