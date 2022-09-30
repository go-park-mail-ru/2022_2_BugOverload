package main

import (
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/options"
	router_ "go-park-mail-ru/2022_2_BugOverload/project/router"
	"go-park-mail-ru/2022_2_BugOverload/project/webserver"
)

func main() {
	options, err := options.GetOptions()
	if err != nil {
		logger.Error(err)

		return
	}

	us := database.NewUserStorage()
	//  cs := database.NewCookieStorage()

	fs := database.NewFilmStorage()
	fs.FillFilmStoragePartOne()
	fs.FillFilmStoragePartTwo()
	//  По аналогии кеш, логер и остальные крупные отдельные сущности

	router := router_.NewRouter(us, fs)

	webserver.Launch(options, router)
}
