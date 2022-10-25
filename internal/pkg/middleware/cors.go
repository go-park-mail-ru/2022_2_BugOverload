package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/rs/cors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type CORSMiddleware struct {
	cors *cors.Cors
}

func NewCORSMiddleware(config *pkg.Cors) *CORSMiddleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	logrus.Info(config.Methods)
	logrus.Info(config.Origins)
	logrus.Info(config.Credentials)
	logrus.Info(config.Headers)
	logrus.Info(config.Debug)

	return &CORSMiddleware{
		cors: cors,
	}
}

func (cmd *CORSMiddleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return cmd.cors.Handler(h)
}
