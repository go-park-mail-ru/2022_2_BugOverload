package view

import (
	"net/http"

	"Kinopoisk/project/application/handling"
)

func CreateMapHandling() *http.ServeMux {
	mux := http.NewServeMux()

	signupHandler := &handling.HandlerSignup{}
	mux.Handle("/v1/auth/signup", signupHandler)

	//  Дальше также сопоставляем обработчик и путь

	return mux
}
