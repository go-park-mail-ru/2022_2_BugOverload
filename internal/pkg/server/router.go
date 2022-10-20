package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factories"
)

func NewRouter(handlers map[string]pkg.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/v1/auth", handlers[factories.AuthRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/login", handlers[factories.LoginRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/signup", handlers[factories.SignupRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", handlers[factories.LogoutRequest].Action).Methods(http.MethodGet)

	// Collections
	router.HandleFunc("/v1/in_cinema", handlers[factories.InCinemaRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/popular_films", handlers[factories.PopularRequest].Action).Methods(http.MethodGet)

	// Films
	router.HandleFunc("/v1/recommendation_film", handlers[factories.RecommendationRequest].Action).Methods(http.MethodGet)

	// Images
	router.HandleFunc("/v1/get_static", handlers[factories.GetImageRequest].Action).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
