package server

import (
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/sirupsen/logrus"

	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factory"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
)

type Server struct {
	config *innerPKG.Config
	logger *logrus.Logger
}

func New(config *innerPKG.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

func (s *Server) Launch() error {
	// DB
	postgres := sqltools.NewPostgresRepository()

	// Initialize repos
	authStorage := repoAuth.NewAuthDatabase(postgres)
	sessionStorage := repoSession.NewSessionCache()

	// Initiaalize services
	authService := serviceAuth.NewAuthService(authStorage)
	sessionService := serviceSession.NewSessionService(sessionStorage)

	handlers := factory.NewHandlersMap(s.config, postgres, sessionService, authService)

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

	if s.config.Server.Protocol == innerPKG.HTTPS {
		err := server.ListenAndServeTLS(s.config.Server.FileTLSCertificate, s.config.Server.FileTLSKey)
		if err != nil {
			logrus.Error(err)
			return err
		}

		return nil
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
