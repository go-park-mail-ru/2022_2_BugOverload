package server

import (
	memory3 "go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wonderivan/logger"

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

// Launch is used to start the server
func (s *Server) Launch() error {
	logger.Info("starting server at " + s.config.Server.BindHTTPAddr)

	fs := memory3.NewFilmStorage()
	fs.FillStorage("test/testdata/films.json")

	router := NewRouter(fs)

	cors := middleware.NewCorsMiddleware(&s.config.Cors)
	routerCors := cors.SetCors(router)

	server := http.Server{
		Addr:         s.config.Server.BindHTTPAddr,
		Handler:      routerCors,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Millisecond,
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
