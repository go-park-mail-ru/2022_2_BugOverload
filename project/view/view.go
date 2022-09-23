package view

import (
	"Kinopoisk/project/application/handling"
	"net/http"
)

func CreateMapHandling() *http.ServeMux {
	mux := http.NewServeMux()

	rootHandler := &handling.HandlerRoot{}
	mux.Handle("/", rootHandler)

	//  Дальше также сопоставляем обработчик и путь

	return mux
}
