package routing

import (
	"Kinopoisk/project/application/http_handlers"
	"net/http"
)

func CreateMapHandling() *http.ServeMux {
	mux := http.NewServeMux()

	rootHandler := &http_handlers.Handler{Name: "root"}
	mux.Handle("/", rootHandler)

	return mux
}
