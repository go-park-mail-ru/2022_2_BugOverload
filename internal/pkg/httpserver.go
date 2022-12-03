package pkg

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
}

func NewServerHTTP(logger *logrus.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Launch(config *Config, router http.Handler) error {
	server := http.Server{
		Addr:         config.ServerHTTPApi.BindAddr,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ServerHTTPApi.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServerHTTPApi.WriteTimeout) * time.Second,
	}

	if config.ServerHTTPApi.Protocol == constparams.HTTPS {
		err := server.ListenAndServeTLS(config.ServerHTTPApi.FileTLSCertificate, config.ServerHTTPApi.FileTLSKey)
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
