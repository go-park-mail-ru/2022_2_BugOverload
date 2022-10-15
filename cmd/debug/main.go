package main

import (
	"flag"
	"go-park-mail-ru/2022_2_BugOverload/pkg"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	configPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/server"
)

// @title MovieGate
// @version 1.0
// @description Server for MovieGate application.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host movie-gate.online
// @BasePath  /
// @termsOfService http://swagger.io/terms/
// @servers http://movie-gate.online
func main() {
	var configPath string

	flag.StringVar(&configPath, "config-path", "cmd/debug/configs/config.toml", "path to config file")

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

	server := server.New(config, logger)

	err = server.Launch()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Server was stopped")
}
