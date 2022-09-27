package webserver

import (
	"github.com/gorilla/mux"
	"github.com/wonderivan/logger"

	"go-park-mail-ru/2022_2_BugOverload/project/options"
	"go-park-mail-ru/2022_2_BugOverload/project/webserver/serveroptions"
	"net/http"
	"time"
)

// Launch is used to start the server
func Launch(options options.Options, router *mux.Router) {
	serverOption, err := serveroptions.GetServerOptions(options.PathServerConfig)
	if err != nil {
		logger.Error(err)

		return
	}

	server := http.Server{
		Addr:         serverOption.Addr,
		Handler:      router,
		ReadTimeout:  time.Duration(serverOption.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(serverOption.WriteTimeout) * time.Second,
	}

	logger.Info("starting server at " + serverOption.Addr)

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err)

		return
	}
}
