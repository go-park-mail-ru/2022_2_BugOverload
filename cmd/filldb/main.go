package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev"
)

func main() {
	var configPath string
	var dataPath string

	flag.StringVar(&configPath, "config-path", "cmd/filldb/configs/config.toml", "path to config file")
	flag.StringVar(&dataPath, "data-path", "test/newdata", "path to data files")

	config := dev.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	filler := dev.NewDBFiller(dataPath, config)

	err = filler.Action()
	if err != nil {
		logrus.Fatalf("FAILED [%s]", err)
	}

	logrus.Info("SUCCESS")
}
