package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/server"
	repoCollection "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection"
	repoFilm "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film"
	repoPerson "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/person"
	repoSearch "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/search"
	serviceWarehouse "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func main() {
	// Config
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/warehouse/configs/debug.toml", "path to config file")

	flag.Parse()

	config := innerPKG.NewConfig()

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

	// Metrics
	metrics := monitoring.NewPrometheusMetrics(config.Metrics.BindHTTPAddr)
	err = metrics.SetupMonitoring()
	if err != nil {
		logger.Fatal(err)
	}

	// Middleware
	md := middleware.NewGRPCMiddleware(logger, metrics)

	// Connections
	postgres := sqltools.NewPostgresRepository()

	// Film repository
	fs := repoFilm.NewFilmPostgres(postgres)

	// Film service
	filmService := serviceWarehouse.NewFilmService(fs)

	// Collection repository
	cs := repoCollection.NewCollectionPostgres(postgres)

	// Collection service
	collectionService := serviceWarehouse.NewCollectionService(cs)

	// Person repository
	ps := repoPerson.NewPersonPostgres(postgres)

	// Person service
	personService := serviceWarehouse.NewPersonService(ps)

	// Search repository
	sr := repoSearch.NewSearchPostgres(postgres)
	// Search service
	searchService := serviceWarehouse.NewSearchService(sr)

	// Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(md.LoggerInterceptor),
		grpc.MaxRecvMsgSize(constparams.BufSizeRequest),
		grpc.MaxSendMsgSize(constparams.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCWarehouse.ConnectionTimeout)*time.Second),
	)

	service := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	logrus.Info(config.ServerGRPCWarehouse.ServiceName + " starting server at " + config.ServerGRPCWarehouse.BindHTTPAddr)

	err = service.StartGRPCServer(config.ServerGRPCWarehouse.BindHTTPAddr)
	if err != nil {
		logrus.Fatal(err)
	}
}
