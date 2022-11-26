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
		Addr:         config.ServerHTTP.BindHTTPAddr,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ServerHTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServerHTTP.WriteTimeout) * time.Second,
	}

	if config.ServerHTTP.Protocol == constparams.HTTPS {
		err := server.ListenAndServeTLS(config.ServerHTTP.FileTLSCertificate, config.ServerHTTP.FileTLSKey)
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
