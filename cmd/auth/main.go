package main

import (
	"flag"
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

	// Middleware
	md := middleware.NewGRPCMiddleware(logger)

	// Connections
	postgres := sqltools.NewPostgresRepository()

	// Auth repository
	authStorage := repoAuth.NewAuthDatabase(postgres)
	sessionStorage := repoSession.NewSessionCache()

	// Auth service
	authService := serviceSession.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	// Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(md.LoggerInterceptor),
		grpc.MaxRecvMsgSize(constparams.BufSizeRequest),
		grpc.MaxSendMsgSize(constparams.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCAuth.ConnectionTimeout)*time.Second),
	)

	service := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	logrus.Info("starting auth server at " + config.ServerGRPCAuth.BindHTTPAddr)

	err = service.StartGRPCServer(config.ServerGRPCAuth.BindHTTPAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerGRPS - service auth was stopped")
}
