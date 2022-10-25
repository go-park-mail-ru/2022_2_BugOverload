package middleware

import (
	"net/http"

	"github.com/rs/cors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

type CORSMiddleware struct {
	Cors *cors.Cors
}

func NewCORSMiddleware(config *pkg.Cors) *CORSMiddleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return &CORSMiddleware{
		Cors: cors,
	}
}

func (cmd *CORSMiddleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return cmd.Cors.Handler(h)
}
