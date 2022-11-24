package main

import (
	"flag"
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
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/image/configs/config.toml", "path to config file")

	config := innerPKG.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	logger, closeResource := pkg.NewLogger(&config.Logger)
	defer func(closer func() error, log *logrus.Logger) {
		err = closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	md := middleware.NewGPRCMiddleware(logger)

	postgres := sqltools.NewPostgresRepository()

	is := repoImage.NewImageS3(config, postgres)

	imageService := serviceImage.NewImageService(is)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(md.LoggerInterceptor),
		grpc.MaxRecvMsgSize(innerPKG.BufSizeRequest),
		grpc.MaxSendMsgSize(innerPKG.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPCImage.ConnectionTimeout)*time.Second),
	)

	service := server.NewImageServiceGRPCServer(grpcServer, imageService)

	logrus.Info("starting server at " + config.ServerGRPCImage.BindHTTPAddr)

	err = service.StartGRPCServer(config.ServerGRPCImage.BindHTTPAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerGRPS - service image was stopped")
}
