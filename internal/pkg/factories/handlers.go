package factories

import (
	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	handlers4 "go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/films/delivery/handlers"
	memoryFilms "go-park-mail-ru/2022_2_BugOverload/internal/films/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/films/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/params"
	handlers2 "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func NewHandlersMap() map[string]interfaces.Handler {
	res := make(map[string]interfaces.Handler)

	// Auth
	userStorage := memoryUser.NewUserCash()
	cookieStorage := memoryCookie.NewCookieCash()

	userService := serviceUser.NewUserService(userStorage, params.ContextTimeout)
	authService := serviceAuth.NewAuthService(cookieStorage, params.ContextTimeout)

	authHandler := handlers2.NewAuthHandler(userService, authService)
	res[AuthRequest] = authHandler

	logoutHandler := handlers2.NewLogoutHandler(userService, authService)
	res[LoginRequest] = logoutHandler

	loginHandler := handlers2.NewLoginHandler(userService, authService)
	res[LogoutRequest] = loginHandler

	singUpHandler := handlers2.NewSingUpHandler(userService, authService)
	res[SignupRequest] = singUpHandler

	// Collections
	pathInCinema := "test/data/incinema.json"
	pathPopular := "test/data/popular.json"

	colStorage := memoryCollection.NewCollectionCash(pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(colStorage, params.ContextTimeout)

	inCinemaHandler := handlers4.NewInCinemaHandler(collectionService)
	res[InCinemaRequest] = inCinemaHandler

	popularHandler := handlers4.NewPopularFilmsHandler(collectionService)
	res[PopularRequest] = popularHandler

	// Films
	pathPreview := "test/data/preview.json"

	filmsS := memoryFilms.NewFilmCash(pathPreview)

	filmsService := serviceFilms.NewFilmService(filmsS, params.ContextTimeout)

	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService, authService)
	res[RecommendationRequest] = recommendationHandler

	return res
}
