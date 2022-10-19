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
	handlers5 "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	S3Image "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	handlers2 "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func NewHandlersMap(config *pkg.Config) map[string]pkg.Handler {
	res := make(map[string]pkg.Handler)

	// Auth
	userStorage := memoryUser.NewUserCache()
	cookieStorage := memoryCookie.NewCookieCache()

	userService := serviceUser.NewUserService(userStorage)
	authService := serviceAuth.NewAuthService(cookieStorage)

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

	colStorage := memoryCollection.NewCollectionCache(pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(colStorage)

	inCinemaHandler := handlers4.NewInCinemaHandler(collectionService)
	res[InCinemaRequest] = inCinemaHandler

	popularHandler := handlers4.NewPopularFilmsHandler(collectionService)
	res[PopularRequest] = popularHandler

	// Films
	pathPreview := "test/data/preview.json"

	filmsStorage := memoryFilms.NewFilmCache(pathPreview)

	filmsService := serviceFilms.NewFilmService(filmsStorage)

	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService, authService)
	res[RecommendationRequest] = recommendationHandler

	// Images
	is, _ := S3Image.NewImageS3(config)

	imageService := serviceImage.NewImageService(is)

	getImageHandler := handlers5.NewGetImageHandler(imageService)
	res[GetImageRequest] = getImageHandler

	return res
}
