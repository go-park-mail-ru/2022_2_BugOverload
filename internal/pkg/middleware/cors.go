package middleware

import (
	"net/http"

	"github.com/rs/cors"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func SetCors(config *pkg.Cors, handler http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return cors.Handler(handler)
}
