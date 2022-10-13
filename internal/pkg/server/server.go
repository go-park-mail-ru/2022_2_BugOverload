package server

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factories"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config *pkg.Config
}

func New(config *pkg.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Launch() error {
	handlers := factories.NewHandlersMap()

	router := NewRouter(handlers)

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
