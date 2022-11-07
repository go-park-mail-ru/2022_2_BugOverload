package server

import (
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factory"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/sirupsen/logrus"

	pkgInner "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
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
	// Initialize repos
	authStorage := repoAuth.NewAuthCache()
	sessionStorage := repoSession.NewSessionCache()

	// Initiaalize services
	authService := serviceAuth.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	handlers := factory.NewHandlersMap(s.config, sessionService, authService)

	mw := middleware.NewMiddleware(s.logger, sessionService, &s.config.Cors)

	router := NewRouter(handlers, mw)

	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	routerCORS := mw.SetCORSMiddleware(router)

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
