package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func NewRouter(handlers map[string]pkg.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/api/v1/auth", handlers[pkg.AuthRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/auth/login", handlers[pkg.LoginRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/signup", handlers[pkg.SignupRequest].Action).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/logout", handlers[pkg.LogoutRequest].Action).Methods(http.MethodGet)

	// Collections
	router.HandleFunc("/api/v1/collection/{tag}", handlers[pkg.TagCollectionRequest].Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}", "delimiter", "{delimiter}")

	// Films
	router.HandleFunc("/api/v1/film/recommendation", handlers[pkg.RecommendationRequest].Action).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/film/{id:[0-9]+}", handlers[pkg.FilmRequest].Action).
		Methods(http.MethodGet).
		Queries("count_images", "{count_images}")

	//  router.HandleFunc("/api/v1/film/{id:[0-9]+}", handlers[pkg.UploadImageRequest].Action).Methods(http.MethodPost)

	// Images
	router.HandleFunc("/api/v1/image", handlers[pkg.DownloadImageRequest].Action).
		Methods(http.MethodGet).
		Queries("object", "{object}", "key", "{key}")

	router.HandleFunc("/api/v1/image", handlers[pkg.UploadImageRequest].Action).
		Methods(http.MethodPost).
		Queries("object", "{object}", "key", "{key}")

	// User
	router.HandleFunc("/api/v1/user/profile/{id:[0-9]+}", handlers[pkg.GetUserProfileRequest].Action).Methods(http.MethodGet)

	// Person
	router.HandleFunc("/api/v1/person/{id:[0-9]+}", handlers[pkg.PersonRequest].Action).
		Methods(http.MethodGet).
		Queries("count_films", "{count_films}", "count_images", "{count_images}")

	// Reviews
	router.HandleFunc("/api/v1/film/{id:[0-9]+}/reviews", handlers[pkg.ReviewsFilmRequest].Action).
		Methods(http.MethodGet).
		Queries("count_reviews", "{count_reviews}", "offset", "{offset}")

	http.Handle("/", router)

	return router
}
