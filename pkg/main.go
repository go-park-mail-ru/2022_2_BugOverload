package main

import (
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/pkg/application/database"
	"go-park-mail-ru/2022_2_BugOverload/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/pkg/options"
	router_ "go-park-mail-ru/2022_2_BugOverload/pkg/router"
	"go-park-mail-ru/2022_2_BugOverload/pkg/webserver"
)

func main() {
	options, err := options.GetOptions()
	if err != nil {
		logger.Error(err)
		return
	}

	us := database.NewUserStorage()
	cs := database.NewCookieStorage()
	fs := database.NewFilmStorage()

	router := router_.NewRouter(us, cs, fs)

	corsOptions := middleware.NewCorsOptions()

	webserver.Launch(options, router, corsOptions)
}
