package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/handlers"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/auth/service"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/auth/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/auth/session/service"
	repoUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	repoCollection "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	repoPerson "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/person"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client"
	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
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
	authHandler := handlers.NewAuthHandler(authService, sessionService)
	authHandler.Configure(router, mw)

	logoutHandler := handlers.NewLogoutHandler(authService, sessionService)
	logoutHandler.Configure(router, mw)

	loginHandler := handlers.NewLoginHandler(authService, sessionService)
	loginHandler.Configure(router, mw)

	singUpHandler := handlers.NewSingUpHandler(authService, sessionService)
	singUpHandler.Configure(router, mw)

	// Collections repository
	collectionStorage := repoCollection.NewCollectionPostgres(postgres)

	// Collections service
	collectionService := serviceCollection.NewCollectionService(collectionStorage)

	// Collections delivery
	stgCollectionHandler := handlers.NewStdCollectionHandler(collectionService)
	stgCollectionHandler.Configure(router, mw)

	premieresCollectionHandler := handlers.NewPremieresCollectionHandler(collectionService)
	premieresCollectionHandler.Configure(router, mw)

	// Films repository
	filmsStorage := repoFilms.NewFilmPostgres(postgres)

	// Films service
	filmsService := serviceCollection.NewFilmService(filmsStorage)

	// Films delivery
	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService)
	recommendationHandler.Configure(router, mw)

	filmHandler := handlers.NewFilmHandler(filmsService)
	filmHandler.Configure(router, mw)

	reviewsHandler := handlers.NewReviewsHandler(filmsService)
	reviewsHandler.Configure(router, mw)

	// Images microservice
	imageService := client.NewImageServiceGRPSClient(grpcConnImage)

	downloadImageHandler := handlers.NewGetImageHandler(imageService)
	downloadImageHandler.Configure(router, mw)

	changeImageHandler := handlers.NewPutImageHandler(imageService)
	changeImageHandler.Configure(router, mw)

	uploadImageHandler := handlers.NewPostImageHandler(imageService)
	uploadImageHandler.Configure(router, mw)

	// User repository
	userRepo := repoUser.NewUserPostgres(postgres)

	// User service
	userService := serviceUser.NewUserProfileService(userRepo, authService)

	// User delivery
	profileHandler := handlers.NewUserProfileHandler(userService)
	profileHandler.Configure(router, mw)

	userSettingsHandler := handlers.NewGetSettingsHandler(userService)
	userSettingsHandler.Configure(router, mw)

	changeUserSettingsHandler := handlers.NewPutSettingsHandler(userService)
	changeUserSettingsHandler.Configure(router, mw)

	filmRateHandler := handlers.NewFilmRateHandler(userService)
	filmRateHandler.Configure(router, mw)

	filmRateDropHandler := handlers.NewFilmRateDropHandler(userService)
	filmRateDropHandler.Configure(router, mw)

	newFilmReviewHandler := handlers.NewFilmReviewHandler(userService)
	newFilmReviewHandler.Configure(router, mw)

	getUserActivityOnFilmHandler := handlers.NewGetActivityOnFilmHandler(userService)
	getUserActivityOnFilmHandler.Configure(router, mw)

	userCollectionsHandler := handlers.NewGetUserCollectionsHandler(userService)
	userCollectionsHandler.Configure(router, mw)

	// Person repository
	personRepo := repoPerson.NewPersonPostgres(postgres)

	// Person service
	personService := serviceCollection.NewPersonService(personRepo)

	// service delivery
	personHandler := handlers.NewPersonHandler(personService)
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
