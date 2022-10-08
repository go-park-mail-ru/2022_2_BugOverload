package server

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/authhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/loginhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/logouthandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/signuphandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popularfilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendation"
)

// NewRouter is constructor for mux
func NewRouter(us *memory.UserStorage, cs *memory.CookieStorage, fs *memory.FilmStorage) *mux.Router {
	router := mux.NewRouter()

	authHandler := authhandler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth", authHandler.Action).Methods(http.MethodGet)
	loginHandler := loginhandler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/login", loginHandler.Action).Methods(http.MethodPost)
	singUpHandler := signuphandler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/signup", singUpHandler.Action).Methods(http.MethodPost)
	logoutHandler := logouthandler.NewHandler(us, cs)
	router.HandleFunc("/v1/auth/logout", logoutHandler.Action).Methods(http.MethodGet)

	inCinemaHandler := incinemafilmshandler.NewCollectionInCinemaHandler(fs)
	router.HandleFunc("/v1/in_cinema", inCinemaHandler.Action).Methods(http.MethodGet)
	popularHandler := popularfilmshandler.NewCollectionPopularHandler(fs)
	router.HandleFunc("/v1/popular_films", popularHandler.Action).Methods(http.MethodGet)
	recommendationHandler := recommendation.NewHandlerRecommendationFilm(fs)
	router.HandleFunc("/v1/recommendation_film", recommendationHandler.GetRecommendedFilm).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
