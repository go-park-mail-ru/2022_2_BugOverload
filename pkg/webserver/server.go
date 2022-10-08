package webserver

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wonderivan/logger"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/pkg/options"
	"go-park-mail-ru/2022_2_BugOverload/pkg/webserver/serveroptions"
)

// Launch is used to start the server
func Launch(options options.Options, router *mux.Router, corsOptions *cors.Cors) {
	serverOption, err := serveroptions.GetServerOptions(options.PathServerConfig)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("starting server at " + options.Port)

	newRouter := corsOptions.Handler(router)

	server := http.Server{
		Addr:         options.Port,
		Handler:      newRouter,
		ReadTimeout:  serverOption.ReadTimeout,
		WriteTimeout: serverOption.ReadTimeout,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err)
		return
	}
}
