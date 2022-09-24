package main

import (
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/project/options"
	"go-park-mail-ru/2022_2_BugOverload/project/webserver"
)

func main() {
	options, err := options.GetOptions()
	if err != nil {
		logger.Error(err)

		return
	}

	webserver.Launch(options)
}
