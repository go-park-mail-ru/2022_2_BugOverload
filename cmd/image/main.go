package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/server"
	repoImage "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
)

func main() {
	// Config
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/image/configs/debug.toml", "path to config file")

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
	metrics := monitoring.NewPrometheusMetrics(config.ServerGRPCImage.ServiceName)
	err = metrics.SetupMonitoring()
	if err != nil {
		logger.Fatal(err)
	}

	// Middleware
	md := middleware.NewGRPCMiddleware(logger, metrics)

	// Connections
	postgres := sqltools.NewPostgresRepository(&config.DatabaseParams)

	// Image repository
	is := repoImage.NewImageS3(config, postgres)

	// Image service
	imageService := serviceImage.NewImageService(is)

	// Metrics server
	go monitoring.CreateNewMonitoringServer(config.Metrics.BindHTTPAddr)

	// Server
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(md.LoggerInterceptor, md.MetricsInterceptor),
		grpc.MaxRecvMsgSize(constparams.BufSizeRequest),
		grpc.MaxSendMsgSize(constparams.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCImage.ConnectionTimeout)*time.Second),
	)

	service := server.NewImageServiceGRPCServer(grpcServer, imageService)

	logrus.Info(config.ServerGRPCImage.ServiceName + " starting server at " + config.ServerGRPCImage.BindAddr)

	err = service.StartGRPCServer(config.ServerGRPCImage.BindAddr)
	if err != nil {
		logrus.Fatal(err)
	}
}
