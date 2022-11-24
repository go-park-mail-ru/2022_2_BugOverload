package server

import (
	"log"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

func NewServerHTTP(config *innerPKG.Config, logger *logrus.Logger) *Server {
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

	grpcConn, err := grpc.Dial(
		s.config.URls.ImageServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grpcConn.Close()

	res := make(map[string]*grpc.ClientConn)
	res["1"] = grpcConn

	handlers := factory.NewHandlersMap(s.config, postgres, sessionService, authService, res)

	mw := middleware.NewMiddleware(s.logger, sessionService, &s.config.Cors)

	router := NewRouter(handlers, mw)

	router.Use(
		mw.SetDefaultLoggerMiddleware,
		mw.UpdateDefaultLoggerMiddleware,
		mw.SetSizeRequest,
		gziphandler.GzipHandler,
	)

	routerCORS := mw.SetCORSMiddleware(router)

	logrus.Info("starting server at " + s.config.ServerHTTP.BindHTTPAddr)

	server := http.Server{
		Addr:         s.config.ServerHTTP.BindHTTPAddr,
		Handler:      routerCORS,
		ReadTimeout:  time.Duration(s.config.ServerHTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.ServerHTTP.WriteTimeout) * time.Second,
	}

	if s.config.ServerHTTP.Protocol == innerPKG.HTTPS {
		err = server.ListenAndServeTLS(s.config.ServerHTTP.FileTLSCertificate, s.config.ServerHTTP.FileTLSKey)
		if err != nil {
			logrus.Error(err)
			return err
		}

		return nil
	}

	err = server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
