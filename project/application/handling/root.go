package handling

import (
	"fmt"
	"net/http"
)

type HandlerRoot struct {
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

func (h *HandlerRoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL.String()) //  Заглушка
	w.Write([]byte("booooooooo"))       //  Заглушка
}
