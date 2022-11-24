package main

import (
	"flag"
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

	postgres := sqltools.NewPostgresRepository()

	is := repoImage.NewImageS3(config, postgres)

	imageService := serviceImage.NewImageService(is)

	grpc := grpc.NewServer(
		grpc.MaxRecvMsgSize(innerPKG.BufSizeRequest),
		grpc.MaxSendMsgSize(innerPKG.BufSizeRequest),
		grpc.ConnectionTimeout(time.Duration(config.ServerGRPC.ConnectionTimeout)*time.Second),
	)

	service := server.NewImageServiceGRPCServer(grpc, imageService)

	logrus.Info("starting server at " + config.ServerGRPC.BindHTTPAddr)

	err = service.StartGRPCServer(config.ServerGRPC.BindHTTPAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerGRPS - service image was stopped")
}
