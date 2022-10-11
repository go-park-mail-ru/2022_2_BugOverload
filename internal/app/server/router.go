package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popularfilmshandler"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/repository/memory"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendationfilmhandler"
	memoryFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/authhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/loginhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/logouthandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/signuphandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/contextparams"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Auth
	userMutex := &sync.Mutex{}
	authMutex := &sync.Mutex{}

	us := memoryUser.NewUserRepo(userMutex)
	cs := memoryCookie.NewCookieRepo(authMutex)

	userService := serviceUser.NewUserService(us, contextparams.ContextTimeout)
	authService := serviceAuth.NewAuthService(cs, contextparams.ContextTimeout)

	authHandler := authhandler.NewHandler(userService, authService)
	logoutHandler := logouthandler.NewHandler(userService, authService)
	loginHandler := loginhandler.NewHandler(userService, authService)
	singUpHandler := signuphandler.NewHandler(userService, authService)

	router.HandleFunc("/v1/auth", authHandler.Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/login", loginHandler.Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/signup", singUpHandler.Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", logoutHandler.Action).Methods(http.MethodGet)

	// Collections
	pathInCinema := "test/testdata/incinema.json"
	pathPopular := "test/testdata/popular.json"

	collectionMutex := &sync.Mutex{}
	colS := memoryCollection.NewCollectionRepo(collectionMutex, pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(colS, contextparams.ContextTimeout)

	inCinemaHandler := incinemafilmshandler.NewHandler(collectionService)
	popularHandler := popularfilmshandler.NewHandler(collectionService)

	router.HandleFunc("/v1/in_cinema", inCinemaHandler.Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/popular_films", popularHandler.Action).Methods(http.MethodGet)

	// Films
	pathPreview := "test/testdata/preview.json"

	filmsMutex := &sync.Mutex{}

	fs := memoryFilms.NewFilmRepo(filmsMutex, pathPreview)

	filmsService := serviceFilms.NewFilmService(fs, contextparams.ContextTimeout)

	recommendationHandler := recommendationfilmhandler.NewHandler(filmsService, authService)

	router.HandleFunc("/v1/recommendation_film", recommendationHandler.Action).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
