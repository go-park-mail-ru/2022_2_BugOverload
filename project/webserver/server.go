package webserver

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/project/options"
	"go-park-mail-ru/2022_2_BugOverload/project/webserver/serveroptions"
)

// Launch is used to start the server
func Launch(options options.Options, router *mux.Router, corsOptions *cors.Cors) {
	serverOption, err := serveroptions.GetServerOptions(options.PathServerConfig)
	if err != nil {
		logger.Error(err)

		return
	}

	logger.Info("starting server at " + serverOption.Port)

	routerCORS := corsOptions.Handler(router)

	server := http.Server{
		Addr:         serverOption.Port,
		Handler:      routerCORS,
		ReadTimeout:  time.Duration(serverOption.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(serverOption.WriteTimeout) * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err)

		return
	}
}
