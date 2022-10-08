package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/auth_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/login_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/logout_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/signup_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinema_films_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popular_films_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendation"
)

// NewRouter is constructor for mux
func NewRouter(us *database.UserStorage, cs *database.CookieStorage, fs *database.FilmStorage) *mux.Router {
	router := mux.NewRouter()

	authHandler := auth_handler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth", authHandler.Action).Methods(http.MethodGet)
	loginHandler := login_handler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/login", loginHandler.Action).Methods(http.MethodPost)
	singUpHandler := signup_handler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/signup", singUpHandler.Action).Methods(http.MethodPost)
	logoutHandler := logout_handler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/logout", logoutHandler.Action).Methods(http.MethodGet)

	inCinemaHandler := incinema_films_handler.NewCollectionInCinemaHandler(fs)
	router.HandleFunc("/v1/in_cinema", inCinemaHandler.Action).Methods(http.MethodGet)
	popularHandler := popular_films_handler.NewCollectionPopularHandler(fs)
	router.HandleFunc("/v1/popular_films", popularHandler.Action).Methods(http.MethodGet)
	recommendationHandler := recommendation.NewHandlerRecommendationFilm(fs)
	router.HandleFunc("/v1/recommendation_film", recommendationHandler.GetRecommendedFilm).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
