package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func NewRouter(handlers map[string]pkg.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/v1/auth", handlers[pkg.AuthRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/login", handlers[pkg.LoginRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/signup", handlers[pkg.SignupRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", handlers[pkg.LogoutRequest].Action).Methods(http.MethodGet)

	// Collections
	router.HandleFunc("/v1/in_cinema", handlers[pkg.InCinemaRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/popular_films", handlers[pkg.PopularRequest].Action).Methods(http.MethodGet)

	// Films
	router.HandleFunc("/v1/recommendation_film", handlers[pkg.RecommendationRequest].Action).Methods(http.MethodGet)

	// Images
	router.HandleFunc("/v1/image", handlers[pkg.DownloadImageRequest].Action).
		Methods(http.MethodGet).
		Queries("object", "{object}", "key", "{key}")

	router.HandleFunc("/v1/image", handlers[pkg.UploadImageRequest].Action).
		Methods(http.MethodPost).
		Queries("object", "{object}", "key", "{key}")

	http.Handle("/", router)

	return router
}
