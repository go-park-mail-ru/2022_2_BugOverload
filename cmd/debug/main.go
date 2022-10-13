package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/server"
)

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
