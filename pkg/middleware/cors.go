package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

const (
	frontendDebug = "http://localhost:3001"
	frontendMain  = "http://movie-gate.online"
)

func NewCorsOptions() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowedOrigins:   []string{frontendDebug, frontendMain},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Content-length"},
		Debug:            true,
	})

	return c
}
