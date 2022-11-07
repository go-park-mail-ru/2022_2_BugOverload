package server

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
)

func NewRouter(handlers map[string]handler.Handler, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()

	// Auth
	handlers[pkg.AuthRequest].Configure(router, mw)
	handlers[pkg.LoginRequest].Configure(router, mw)
	handlers[pkg.SignupRequest].Configure(router, mw)
	handlers[pkg.LogoutRequest].Configure(router, mw)

	// Collections
	handlers[pkg.InCinemaRequest].Configure(router, mw)
	handlers[pkg.PopularRequest].Configure(router, mw)

	// Films
	handlers[pkg.RecommendationRequest].Configure(router, mw)

	//  router.HandleFunc("/api/v1/film/{id:[0-9]+}", handlers[pkg.UploadImageRequest].Action).Methods(http.MethodPost)

	// Images
	handlers[pkg.DownloadImageRequest].Configure(router, mw)
	handlers[pkg.UploadImageRequest].Configure(router, mw)

	// User
	handlers[pkg.GetUserProfile].Configure(router, mw)

	http.Handle("/", router)

	return router
}
