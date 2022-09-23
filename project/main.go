package main

import (
	"github.com/wonderivan/logger"

	"Kinopoisk/project/options"
	"Kinopoisk/project/webserver"
)

func main() {
	options, err := options.GetOptions()
	if err != nil {
		logger.Error(err)

		return
	}

	webserver.Launch(options)
}
