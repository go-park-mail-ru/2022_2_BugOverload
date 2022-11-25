package main

import (
	"context"
	"flag"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/NYTimes/gziphandler"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factory"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/server"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/prod/configs/config.toml", "path to config file")

	config := configPKG.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	// Setup logger
	logger, closeResource := pkg.NewLogger(&config.Logger)
	defer func(closer func() error, log *logrus.Logger) {
		err = closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	// Connections
	postgres := sqltools.NewPostgresRepository()

	// Repo
	authStorage := repoAuth.NewAuthDatabase(postgres)
	sessionStorage := repoSession.NewSessionCache()

	// Service
	authService := serviceAuth.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	// Microservices
	// GRPC in dev ---------------- WARNING timeout not work. How to solve it?
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ServerGRPCImage.WorkTimeout)*time.Second)
	defer cancel()

	grpcConn, err := grpc.DialContext(ctx,
		config.ServerGRPCImage.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatal("cant connect to grpc ", err)
	}
	defer grpcConn.Close()

	res := make(map[string]*grpc.ClientConn)
	res["1"] = grpcConn

	handlers := factory.NewHandlersMap(config, postgres, sessionService, authService, res)
	// GRPC in dev ---------------- WARNING

	mw := middleware.NewHTTPMiddleware(logger, sessionService, &config.Cors)

	router := server.NewRouter(handlers, mw)

	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	routerCORS := mw.SetCORSMiddleware(router)

	logrus.Info("starting server at " + config.ServerHTTP.BindHTTPAddr + " on protocol " + config.ServerHTTP.Protocol)

	// Server
	server := server.NewServerHTTP(logger)

	err = server.Launch(config, routerCORS)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerHTTP was stopped")
}
