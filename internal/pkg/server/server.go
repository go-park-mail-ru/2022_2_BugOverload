package server

import (
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
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
	corsMW := middleware.NewCORSMiddleware(&s.config.Cors)
	loggerMW := middleware.NewLoggerMiddleware(s.logger)
	requestParamsMW := middleware.NewRequestMiddleware()

	router.Use(
		loggerMW.SetDefaultLoggerMiddleware,
		loggerMW.UpdateDefaultLoggerMiddleware,
		requestParamsMW.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	routerCORS := corsMW.SetCORSMiddleware(router)

	logrus.Info("starting server at " + s.config.Server.BindHTTPAddr)

	server := http.Server{
		Addr:         s.config.Server.BindHTTPAddr,
		Handler:      routerCORS,
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
