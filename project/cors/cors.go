package cors

import (
	"net/http"

	"github.com/rs/cors"
)

const frontend = "http://localhost:3000"

func NewCorsOptions() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowedOrigins:   []string{frontend},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Content-length"},
		Debug:            true,
	})

	return c
}
