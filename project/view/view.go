package view

import (
	"net/http"

	"Kinopoisk/project/application/handling"
)

func CreateMapHandling() *http.ServeMux {
	mux := http.NewServeMux()

	rootHandler := &handling.HandlerRoot{}
	mux.Handle("/", rootHandler)

	//  Дальше также сопоставляем обработчик и путь

	return mux
}
