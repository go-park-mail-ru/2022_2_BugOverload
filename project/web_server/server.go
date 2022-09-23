package web_server

import (
	"Kinopoisk/project/options"
	"Kinopoisk/project/web_server/routing"
	"Kinopoisk/project/web_server/server_options"
	"fmt"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

func LaunchServer(options options.Options) {
	serverOption, err := server_options.GetServerOptions(options.PathServerConfig)
	if err != nil {
		logger.Error(err)

		return
	}

	server := http.Server{
		Addr:         ":" + serverOption.Addr,
		Handler:      routing.CreateMapHandling(),
		ReadTimeout:  time.Duration(serverOption.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(serverOption.WriteTimeout) * time.Second,
	}

	fmt.Println("starting server at :" + serverOption.Addr)

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err)

		return
	}
}
