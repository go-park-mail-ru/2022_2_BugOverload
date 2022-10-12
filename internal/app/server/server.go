package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"go-park-mail-ru/2022_2_BugOverload/internal"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/middleware"
)

type Server struct {
	config *internal.Config
}

func New(config *internal.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Launch() error {
	router := NewRouter()

	cors := middleware.NewCorsMiddleware(&s.config.Cors)
	routerCors := cors.SetCors(router)

	logrus.Info("starting server at " + s.config.Server.BindHTTPAddr)

	server := http.Server{
		Addr:         s.config.Server.BindHTTPAddr,
		Handler:      routerCors,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
