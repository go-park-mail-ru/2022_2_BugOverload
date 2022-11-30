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

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/handlers"
	clientAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client"
	clientImage "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client"
	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	repoUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	clientWarehouse "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// @title MovieGate
// @version 1.0
// @description ServerHTTPApi for MovieGate application.

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
	grpcConnWarehouse, err := grpc.Dial(
		config.ServerGRPCWarehouse.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnWarehouse.Close()
	grpcConnAuth, err := grpc.Dial(
		config.ServerGRPCAuth.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatal("cant connect to grpc ", err)
	}
	defer grpcConnAuth.Close()

	// Auth
	authService := clientAuth.NewAuthServiceGRPSClient(grpcConnAuth)

	// Metrics
	metrics := monitoring.NewPrometheusMetrics(config.Metrics.BindHTTPAddr)
	err = metrics.SetupMonitoring()
	if err != nil {
		logger.Fatal(err)
	}

	// Middleware
	mw := middleware.NewHTTPMiddleware(logger, authService, &config.Cors, metrics)

	// Router
	router := mux.NewRouter()

	// Auth delivery
	authHandler := handlers.NewAuthHandler(authService)
	authHandler.Configure(router, mw)

	logoutHandler := handlers.NewLogoutHandler(authService)
	logoutHandler.Configure(router, mw)

	loginHandler := handlers.NewLoginHandler(authService)
	loginHandler.Configure(router, mw)

	singUpHandler := handlers.NewSingUpHandler(authService)
	singUpHandler.Configure(router, mw)

	// Warehouse microservice
	warehouseService := clientWarehouse.NewWarehouseServiceGRPSClient(grpcConnWarehouse)

	// Collections delivery
	stgCollectionHandler := handlers.NewStdCollectionHandler(warehouseService)
	stgCollectionHandler.Configure(router, mw)

	premieresCollectionHandler := handlers.NewPremieresCollectionHandler(warehouseService)
	premieresCollectionHandler.Configure(router, mw)

	getCollectionFilmsHandler := handlers.NewGetCollectionHandler(warehouseService)
	getCollectionFilmsHandler.Configure(router, mw)

	// Films delivery
	recommendationHandler := handlers.NewRecommendationFilmHandler(warehouseService)
	recommendationHandler.Configure(router, mw)

	filmHandler := handlers.NewFilmHandler(warehouseService)
	filmHandler.Configure(router, mw)

	reviewsHandler := handlers.NewReviewsHandler(warehouseService)
	reviewsHandler.Configure(router, mw)

	// Person delivery
	personHandler := handlers.NewPersonHandler(warehouseService)
	personHandler.Configure(router, mw)

	// Images microservice
	imageService := clientImage.NewImageServiceGRPSClient(grpcConnImage)

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

	addFilmToUserCollectionHandler := handlers.NewAddFilmToUserCollectionHandler(userService)
	addFilmToUserCollectionHandler.Configure(router, mw)

	dropFilmFromUserCollectionHandler := handlers.NewDropFilmFromUserCollectionHandler(userService)
	dropFilmFromUserCollectionHandler.Configure(router, mw)

	// Search delivery
	searchHandler := handlers.NewSearchHandler(warehouseService)
	searchHandler.Configure(router, mw)

	http.Handle("/", router)

	// Set middleware
	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		mw.SetDefaultMetrics,
		gziphandler.GzipHandler,
	)

	routerCORS := mw.SetCORSMiddleware(router)

	logrus.Info(config.ServerHTTPApi.ServiceName + "starting server at " + config.ServerHTTPApi.BindHTTPAddr + " on protocol " + config.ServerHTTPApi.Protocol)

	// Server
	server := configPKG.NewServerHTTP(logger)

	err = server.Launch(config, routerCORS)
	if err != nil {
		logrus.Fatal(err)
	}
}
