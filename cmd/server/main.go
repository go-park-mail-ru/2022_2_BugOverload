package server

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"go-park-mail-ru/2022_2_BugOverload/internal"

	main_server "go-park-mail-ru/2022_2_BugOverload/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	config := internal.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	server := main_server.New(config)

	err = server.Launch()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Server was stopped")
}
