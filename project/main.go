package main

import (
	"Kinopoisk/project/options"
	"Kinopoisk/project/web_server"
	"github.com/wonderivan/logger"
)

func main() {
	options, err := options.GetOptions()
	if err != nil {
		logger.Error(err)

		return
	}

	web_server.LaunchServer(options)
}
