package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
)

// NewRouter is constructor for mux
func NewRouter(us *database.UserStorage) *mux.Router {
	router := mux.NewRouter()

	authHandler := handlers.NewHandlerAuth(us)
	router.HandleFunc("/v1/auth/signup", authHandler.Signup).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/login", authHandler.Login).Methods(http.MethodPost)

	http.Handle("/", router)

	return router
}
