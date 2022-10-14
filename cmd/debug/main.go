package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
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

	config := pkg.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	server := server.New(config)

	err = server.Launch()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Server was stopped")
}
