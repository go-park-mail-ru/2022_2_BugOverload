package router

import (
	"github.com/gorilla/mux"
	"go-park-mail-ru/2022_2_BugOverload/project/application/storages"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	//  tmp solution
	us := storages.NewUserStorage()
	//  cs := storages.NewCookieStorage()  //  может понадобиться

	signupHandler := handlers.NewHandlerSignup(us)
	router.Handle("/v1/auth/signup", signupHandler).Methods(http.MethodPost)

	//  Дальше также сопоставляем обработчик и путь

	http.Handle("/", router)

	return router
}
