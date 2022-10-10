package server

import (
	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	memory3 "go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popularfilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendation"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/authhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/loginhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/logouthandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/signuphandler"
)

// NewRouter is constructor for mux
func NewRouter(fs *memory3.FilmStorage) *mux.Router {
	router := mux.NewRouter()

	userMutex := &sync.Mutex{}
	authMutex := &sync.Mutex{}

	us := memoryUser.NewUserRepo(userMutex)
	cs := memoryCookie.NewCookieRepo(authMutex)

	userService := serviceUser.NewUserService(us, 2)
	authService := serviceAuth.NewAuthService(cs, 2)

	authHandler := authhandler.NewHandler(userService, authService)
	logoutHandler := logouthandler.NewHandler(userService, authService)
	loginHandler := loginhandler.NewHandler(userService, authService)
	singUpHandler := signuphandler.NewHandler(userService, authService)

	router.HandleFunc("/v1/auth", authHandler.Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/login", loginHandler.Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/signup", singUpHandler.Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", logoutHandler.Action).Methods(http.MethodGet)

	inCinemaHandler := incinemafilmshandler.NewHandler(fs)
	router.HandleFunc("/v1/in_cinema", inCinemaHandler.Action).Methods(http.MethodGet)
	popularHandler := popularfilmshandler.NewHandler(fs)
	router.HandleFunc("/v1/popular_films", popularHandler.Action).Methods(http.MethodGet)
	recommendationHandler := recommendation.NewHandlerRecommendationFilm(fs)
	router.HandleFunc("/v1/recommendation_film", recommendationHandler.GetRecommendedFilm).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
