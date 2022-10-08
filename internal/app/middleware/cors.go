package middleware

import (
	"net/http"

	"github.com/rs/cors"

	"go-park-mail-ru/2022_2_BugOverload/internal"
)

func NewCorsMiddleware(config *internal.Cors) CorsMiddleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return CorsMiddleware{
		setup: cors,
	}
}

type CorsMiddleware struct {
	setup *cors.Cors
}

func (mw *CorsMiddleware) SetCors(handler http.Handler) http.Handler {
	return mw.setup.Handler(handler)
}
