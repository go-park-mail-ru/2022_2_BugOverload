package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers/auth"
)

// NewRouter is constructor for mux
func NewRouter(us *database.UserStorage, cs *database.CookieStorage, fs *database.FilmStorage) *mux.Router {
	router := mux.NewRouter()

	authHandler := auth.NewHandlerAuth(us, cs)

	filmHandler := handlers.NewHandlerFilms(fs)

	router.HandleFunc("/v1/auth/signup", authHandler.Signup).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth", authHandler.Auth).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/logout", authHandler.Logout).Methods(http.MethodGet)

	router.HandleFunc("/v1/popular_films", filmHandler.GetPopularFilms).Methods(http.MethodGet)
	router.HandleFunc("/v1/in_cinema", filmHandler.GetFilmsInCinema).Methods(http.MethodGet)
	router.HandleFunc("/v1/recommendation_film", filmHandler.GetFilmToPoster).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
