package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/factories"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/interfaces"
)

func NewRouter(handlers map[string]interfaces.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/v1/auth", handlers[factories.AuthRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/login", handlers[factories.LoginRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/signup", handlers[factories.LogoutRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", handlers[factories.SignupRequest].Action).Methods(http.MethodGet)

	// Collections
	router.HandleFunc("/v1/in_cinema", handlers[factories.InCinemaRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/v1/popular_films", handlers[factories.PopularRequest].Action).Methods(http.MethodGet)

	// Films
	router.HandleFunc("/v1/recommendation_film", handlers[factories.RecommendationRequest].Action).Methods(http.MethodGet)

	http.Handle("/", router)

	return router
}
