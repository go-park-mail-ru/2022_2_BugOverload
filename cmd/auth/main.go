package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/monitoring"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/server"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/auth"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/session"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func main() {
	// Config
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/auth/configs/debug.toml", "path to config file")

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
	metrics := monitoring.NewPrometheusMetrics(config.ServerGRPCAuth.ServiceName)
	err = metrics.SetupMonitoring()
	if err != nil {
		logger.Fatal(err)
	}

	// Middleware
	md := middleware.NewGRPCMiddleware(logger, metrics)

	// Connections
	postgres := sqltools.NewPostgresRepository()

	// Auth repository
	authStorage := repoAuth.NewAuthDatabase(postgres)
	sessionStorage := repoSession.NewSessionCache()

	// Auth service
	authService := serviceSession.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	// Metrics server
	go monitoring.CreateNewMonitoringServer(config.Metrics.BindHTTPAddr)

	// Server
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(md.LoggerInterceptor, md.MetricsInterceptor),
		grpc.MaxRecvMsgSize(constparams.BufSizeRequest),
		grpc.MaxSendMsgSize(constparams.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCAuth.ConnectionTimeout)*time.Second),
	)

	service := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	logrus.Info(config.ServerGRPCAuth.ServiceName + " starting server at " + config.ServerGRPCAuth.BindAddr)

	err = service.StartGRPCServer(config.ServerGRPCAuth.BindAddr)
	if err != nil {
		logrus.Fatal(err)
	}
}
