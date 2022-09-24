package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	signupHandler := &handlers.HandlerSignup{}
	router.Handle("/v1/auth/signup", signupHandler)

	//  Дальше также сопоставляем обработчик и путь

	http.Handle("/", router)

	return router
}
