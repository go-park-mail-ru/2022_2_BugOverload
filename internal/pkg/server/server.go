package server

import (
	"github.com/NYTimes/gziphandler"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factories"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
)

type Server struct {
	config *pkgInner.Config
	logger *logrus.Logger
}

func New(config *pkgInner.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Launch() error {
	handlers := factories.NewHandlersMap(s.config)

	router := NewRouter(handlers)
	routerCors := middleware.SetCors(&s.config.Cors, router)

	utilsMiddleware := middleware.NewLoggerMiddleware(s.logger)
	router.Use(utilsMiddleware.SetDefaultLoggerMiddleware, utilsMiddleware.UpdateDefaultLoggerMiddleware)

	routerCorsWithGz := gziphandler.GzipHandler(routerCors)

	logrus.Info("starting server at " + s.config.Server.BindHTTPAddr)

	server := http.Server{
		Addr:         s.config.Server.BindHTTPAddr,
		Handler:      routerCorsWithGz,
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
