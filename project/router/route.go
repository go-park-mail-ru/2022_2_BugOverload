package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
)

// NewRouter is constructor for mux
func NewRouter(us *database.UserStorage) *mux.Router {
	router := mux.NewRouter()

	signupHandler := handlers.NewHandlerSignup(us)
	router.HandleFunc("/v1/auth/signup", signupHandler.Signup).Methods(http.MethodPost)

	//  Дальше также сопоставляем обработчик и путь

	http.Handle("/", router)

	return router
}
