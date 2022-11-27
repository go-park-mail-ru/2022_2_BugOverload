package main

import (
	"flag"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	handlersAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/handlers"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	handlersCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	repoCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	handlersFilm "go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client"
	handlersImage "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	handlersPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/delivery/handlers"
	repoPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/repository"
	servicePerson "go-park-mail-ru/2022_2_BugOverload/internal/person/service"
	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	handlersUser "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	repoUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// @title MovieGate
// @version 1.0
// @description ServerHTTP for MovieGate application.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host movie-gate.online
// @BasePath  /
// @termsOfService http://swagger.io/terms/
// @servers http://movie-gate.online
func main() {
	// Config
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/api/configs/debug.toml", "path to config file")

	flag.Parse()

	config := configPKG.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	// Logger
	logger, closeResource := pkg.NewLogger(&config.Logger)
	defer func(closer func() error, log *logrus.Logger) {
		err = closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	// Connections
	postgres := sqltools.NewPostgresRepository()
	// Microservices
	grpcConnImage, err := grpc.Dial(
		config.ServerGRPCImage.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnImage.Close()

	// Router
	router := mux.NewRouter()

	// Auth repository
	authStorage := repoAuth.NewAuthDatabase(postgres)
	sessionStorage := repoSession.NewSessionCache()

	// Auth service
	authService := serviceAuth.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	// Middleware
	mw := middleware.NewHTTPMiddleware(logger, sessionService, &config.Cors)

	// Auth delivery
	authHandler := handlersAuth.NewAuthHandler(authService, sessionService)
	authHandler.Configure(router, mw)

	logoutHandler := handlersAuth.NewLogoutHandler(authService, sessionService)
	logoutHandler.Configure(router, mw)

	loginHandler := handlersAuth.NewLoginHandler(authService, sessionService)
	loginHandler.Configure(router, mw)

	singUpHandler := handlersAuth.NewSingUpHandler(authService, sessionService)
	singUpHandler.Configure(router, mw)

	// Collections repository
	collectionStorage := repoCollection.NewCollectionPostgres(postgres)

	// Collections service
	collectionService := serviceCollection.NewCollectionService(collectionStorage)

	// Collections delivery
	stgCollectionHandler := handlersCollection.NewStdCollectionHandler(collectionService)
	stgCollectionHandler.Configure(router, mw)

	userCollectionsHandler := handlersCollection.NewGetUserCollectionsHandler(collectionService)
	userCollectionsHandler.Configure(router, mw)

	premieresCollectionHandler := handlersCollection.NewPremieresCollectionHandler(collectionService)
	premieresCollectionHandler.Configure(router, mw)

	addFilmToCollectionHandler := handlersCollection.NewAddFilmHandler(collectionService)
	addFilmToCollectionHandler.Configure(router, mw)

	dropFilmFromCollectionHandler := handlersCollection.NewDropFilmHandler(collectionService)
	dropFilmFromCollectionHandler.Configure(router, mw)

	// Films repository
	filmsStorage := repoFilms.NewFilmPostgres(postgres)

	// Films service
	filmsService := serviceFilms.NewFilmService(filmsStorage)

	// Films delivery
	recommendationHandler := handlersFilm.NewRecommendationFilmHandler(filmsService)
	recommendationHandler.Configure(router, mw)

	filmHandler := handlersFilm.NewFilmHandler(filmsService)
	filmHandler.Configure(router, mw)

	reviewsHandler := handlersFilm.NewReviewsHandler(filmsService)
	reviewsHandler.Configure(router, mw)

	// Images microservice
	imageService := client.NewImageServiceGRPSClient(grpcConnImage)

	downloadImageHandler := handlersImage.NewGetImageHandler(imageService)
	downloadImageHandler.Configure(router, mw)

	changeImageHandler := handlersImage.NewPutImageHandler(imageService)
	changeImageHandler.Configure(router, mw)

	uploadImageHandler := handlersImage.NewPostImageHandler(imageService)
	uploadImageHandler.Configure(router, mw)

	// User repository
	userRepo := repoUser.NewUserPostgres(postgres)

	// User service
	userService := serviceUser.NewUserProfileService(userRepo, authService)

	// User delivery
	profileHandler := handlersUser.NewUserProfileHandler(userService)
	profileHandler.Configure(router, mw)

	userSettingsHandler := handlersUser.NewGetSettingsHandler(userService)
	userSettingsHandler.Configure(router, mw)

	changeUserSettingsHandler := handlersUser.NewPutSettingsHandler(userService)
	changeUserSettingsHandler.Configure(router, mw)

	filmRateHandler := handlersUser.NewFilmRateHandler(userService)
	filmRateHandler.Configure(router, mw)

	filmRateDropHandler := handlersUser.NewFilmRateDropHandler(userService)
	filmRateDropHandler.Configure(router, mw)

	newFilmReviewHandler := handlersUser.NewFilmReviewHandler(userService)
	newFilmReviewHandler.Configure(router, mw)

	getUserActivityOnFilmHandler := handlersUser.NewGetActivityOnFilmHandler(userService)
	getUserActivityOnFilmHandler.Configure(router, mw)

	// Person repository
	personRepo := repoPerson.NewPersonPostgres(postgres)

	// Person service
	personService := servicePerson.NewPersonService(personRepo)

	// service delivery
	personHandler := handlersPerson.NewPersonHandler(personService)
	personHandler.Configure(router, mw)

	http.Handle("/", router)

	// Set middleware
	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	routerCORS := mw.SetCORSMiddleware(router)

	logrus.Info("starting server at " + config.ServerHTTP.BindHTTPAddr + " on protocol " + config.ServerHTTP.Protocol)

	// Server
	server := configPKG.NewServerHTTP(logger)

	err = server.Launch(config, routerCORS)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerHTTP was stopped")
}
