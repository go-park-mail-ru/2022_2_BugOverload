package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
)

func NewRouter(database *database.Database) *mux.Router {
	router := mux.NewRouter()

	signupHandler := handlers.NewHandlerSignup(database.UserStorage)
	router.HandleFunc("/v1/auth/signup", signupHandler.Signup).Methods(http.MethodPost)

	//  Дальше также сопоставляем обработчик и путь

	http.Handle("/", router)

	return router
}
