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

	logrus.Debug(config.Methods)
	logrus.Debug(config.Origins)
	logrus.Debug(config.Credentials)
	logrus.Debug(config.Headers)
	logrus.Debug(config.Debug)

	return &CORSMiddleware{
		cors: cors,
	}
}

func (cmd *CORSMiddleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return cmd.cors.Handler(h)
}
