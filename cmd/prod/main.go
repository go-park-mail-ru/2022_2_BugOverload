package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/pkg"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/server"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/prod/configs/config.toml", "path to config file")

	config := configPKG.NewConfig()

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

	server := server.NewServerHTTP(config, logger)

	err = server.Launch()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("ServerHTTP was stopped")
}
