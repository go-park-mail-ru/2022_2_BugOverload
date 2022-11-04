package factory

import (
	handlersAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/handlers"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	handlersFilm "go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	handlersImage "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	repoImage "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	handlersPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/delivery/handlers"
	repoPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/repository"
	servicePerson "go-park-mail-ru/2022_2_BugOverload/internal/person/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	handlersUser "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	repoUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func NewHandlersMap(config *pkg.Config) map[string]pkg.Handler {
	res := make(map[string]pkg.Handler)

	// DB
	postgres := sqltools.NewPostgresRepository()

	// Auth
	authStorage := repoAuth.NewAuthCache()
	sessionStorage := repoSession.NewSessionCache()

	authService := serviceAuth.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	authHandler := handlersAuth.NewAuthHandler(authService, sessionService)
	res[pkg.AuthRequest] = authHandler

	logoutHandler := handlersAuth.NewLogoutHandler(authService, sessionService)
	res[pkg.LogoutRequest] = logoutHandler

	loginHandler := handlersAuth.NewLoginHandler(authService, sessionService)
	res[pkg.LoginRequest] = loginHandler

	singUpHandler := handlersAuth.NewSingUpHandler(authService, sessionService)
	res[pkg.SignupRequest] = singUpHandler

	// Collections

	// Films
	pathPreview := "test/data/preview.json"

	filmsStorage := repoFilms.NewFilmCache(pathPreview)

	filmsService := serviceFilms.NewFilmService(filmsStorage)

	recommendationHandler := handlersFilm.NewRecommendationFilmHandler(filmsService, sessionService)
	res[pkg.RecommendationRequest] = recommendationHandler

	// Images
	is := repoImage.NewImageS3(config)

	imageService := serviceImage.NewImageService(is)

	downloadImageHandler := handlersImage.NewGetImageHandler(imageService)
	res[pkg.DownloadImageRequest] = downloadImageHandler

	changeImageHandler := handlersImage.NewPutImageHandler(imageService)
	res[pkg.ChangeImageRequest] = changeImageHandler

	uploadImageHandler := handlersImage.NewPostImageHandler(imageService)
	res[pkg.UploadImageRequest] = uploadImageHandler

	// Users
	userRepo := repoUser.NewUserPostgres(postgres)

	userService := serviceUser.NewUserProfileService(userRepo)

	profileHandler := handlersUser.NewUserProfileHandler(userService)
	res[pkg.GetUserProfileRequest] = profileHandler

	// Persons
	personRepo := repoPerson.NewPersonPostgres(postgres)

	personService := servicePerson.NewPersonService(personRepo)

	personHandler := handlersPerson.NewPersonHandler(personService)
	res[pkg.GetPersonRequest] = personHandler

	return res
}
