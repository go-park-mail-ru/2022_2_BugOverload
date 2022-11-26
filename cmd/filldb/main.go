package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/dev/fillerdb"
)

func main() {
	var configPath string
	var dataPath string

	flag.StringVar(&configPath, "config-path", "cmd/filldb/configs/debug.toml", "path to config file")
	flag.StringVar(&dataPath, "data-path", "test/newdata", "path to data files")

	flag.Parse()

	config := fillerdb.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	filler, err := fillerdb.NewDBFiller(dataPath, config)
	if err != nil {
		logrus.Fatalf("FAILED  [%s]", err)
	}

	err = filler.Action()
	if err != nil {
		logrus.Fatalf("FAILED [%s]", err)
	}

	logrus.Info("SUCCESS")
}
