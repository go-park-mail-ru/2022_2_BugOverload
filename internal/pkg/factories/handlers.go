package factories

import (
	handlers4 "go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	memoryFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	handlers5 "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	S3Image "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	handlers2 "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func NewHandlersMap(config *pkg.Config) map[string]pkg.Handler {
	res := make(map[string]pkg.Handler)

	// Auth
	userStorage := memoryUser.NewUserCache()
	sessionStorage := memoryCookie.NewSessionCache()

	userService := serviceUser.NewUserService(userStorage)
	sessionService := serviceAuth.NewSessionService(sessionStorage)

	authHandler := handlers2.NewAuthHandler(userService, sessionService)
	res[pkg.AuthRequest] = authHandler

	logoutHandler := handlers2.NewLogoutHandler(userService, sessionService)
	res[pkg.LogoutRequest] = logoutHandler

	loginHandler := handlers2.NewLoginHandler(userService, sessionService)
	res[pkg.LoginRequest] = loginHandler

	singUpHandler := handlers2.NewSingUpHandler(userService, sessionService)
	res[pkg.SignupRequest] = singUpHandler

	// Collections
	pathInCinema := "test/data/incinema.json"
	pathPopular := "test/data/popular.json"

	colStorage := memoryCollection.NewCollectionCache(pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(colStorage)

	inCinemaHandler := handlers4.NewInCinemaHandler(collectionService)
	res[pkg.InCinemaRequest] = inCinemaHandler

	popularHandler := handlers4.NewPopularFilmsHandler(collectionService)
	res[pkg.PopularRequest] = popularHandler

	// Films
	pathPreview := "test/data/preview.json"

	filmsStorage := memoryFilms.NewFilmCache(pathPreview)

	filmsService := serviceFilms.NewFilmService(filmsStorage)

	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService, sessionService)
	res[pkg.RecommendationRequest] = recommendationHandler

	// Images
	is := S3Image.NewImageS3(config)

	imageService := serviceImage.NewImageService(is)

	downloadImageHandler := handlers5.NewDownloadImageHandler(imageService)
	res[pkg.DownloadImageRequest] = downloadImageHandler

	uploadImageHandler := handlers5.NewUploadImageHandler(imageService)
	res[pkg.UploadImageRequest] = uploadImageHandler

	return res
}
