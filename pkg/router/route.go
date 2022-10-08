package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/pkg/application/database"
	"go-park-mail-ru/2022_2_BugOverload/pkg/application/handlers/auth"
	"go-park-mail-ru/2022_2_BugOverload/pkg/application/handlers/content"
)

// NewRouter is constructor for mux
func NewRouter(us *database.UserStorage, cs *database.CookieStorage, fs *database.FilmStorage) *mux.Router {
	router := mux.NewRouter()

	authHandler := auth.NewHandlerAuth(us, cs)
	router.HandleFunc("/v1/auth/signup", authHandler.Signup).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth", authHandler.Auth).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/logout", authHandler.Logout).Methods(http.MethodGet)

	filmHandler := content.NewHandlerFilms(fs)
	router.HandleFunc("/v1/popular_films", filmHandler.GetPopularFilms).Methods(http.MethodGet)
	router.HandleFunc("/v1/in_cinema", filmHandler.GetFilmsInCinema).Methods(http.MethodGet)
	router.HandleFunc("/v1/recommendation_film", filmHandler.GetRecommendedFilm).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
