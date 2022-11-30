package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
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

	// Middleware
	md := middleware.NewGRPCMiddleware(logger)

	// Connections
	postgres := sqltools.NewPostgresRepository()

	// Image repository
	is := repoImage.NewImageS3(config, postgres)

	// Image service
	imageService := serviceImage.NewImageService(is)

	// Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(md.LoggerInterceptor),
		grpc.MaxRecvMsgSize(constparams.BufSizeRequest),
		grpc.MaxSendMsgSize(constparams.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCImage.ConnectionTimeout)*time.Second),
	)

	service := server.NewImageServiceGRPCServer(grpcServer, imageService)

	logrus.Info("starting image server at " + config.ServerGRPCImage.BindHTTPAddr)

	err = service.StartGRPCServer(config.ServerGRPCImage.BindHTTPAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerGRPS - service image was stopped")
}
