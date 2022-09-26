package webserver

import (
	"fmt"
	"github.com/wonderivan/logger"
	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"net/http"
	"time"

	"go-park-mail-ru/2022_2_BugOverload/project/options"
	"go-park-mail-ru/2022_2_BugOverload/project/router"
	"go-park-mail-ru/2022_2_BugOverload/project/webserver/server_options"
)

func Launch(options options.Options, database *database.Database) {
	serverOption, err := server_options.GetServerOptions(options.PathServerConfig)
	if err != nil {
		logger.Error(err)

		return
	}

	server := http.Server{
		Addr:         serverOption.Addr,
		Handler:      router.NewRouter(database),
		ReadTimeout:  time.Duration(serverOption.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(serverOption.WriteTimeout) * time.Second,
	}

	fmt.Println("starting server at " + serverOption.Addr)

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err)

		return
	}
}
