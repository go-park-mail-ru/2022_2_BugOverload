package main

import (
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/cors"
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
	cs := database.NewCookieStorage()

	router := router_.NewRouter(us, cs)
	corsOptions := cors.NewCorsOptions()

	webserver.Launch(options, router, corsOptions)
}
