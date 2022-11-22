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
	handlers[pkg.TagCollectionRequest].Configure(router, mw)
	handlers[pkg.UserCollectionsRequest].Configure(router, mw)
	handlers[pkg.PremiersCollectionRequest].Configure(router, mw)

	// Films
	handlers[pkg.FilmRequest].Configure(router, mw)
	handlers[pkg.RecommendationRequest].Configure(router, mw)

	// Images
	handlers[pkg.DownloadImageRequest].Configure(router, mw)
	handlers[pkg.UploadImageRequest].Configure(router, mw)
	handlers[pkg.ChangeImageRequest].Configure(router, mw)

	// User
	handlers[pkg.GetUserProfileRequest].Configure(router, mw)
	handlers[pkg.GetUserSettingsRequest].Configure(router, mw)
	handlers[pkg.PutUserSettingsRequest].Configure(router, mw)
	handlers[pkg.FilmRateRequest].Configure(router, mw)
	handlers[pkg.FilmRateDropRequest].Configure(router, mw)
	handlers[pkg.NewFilmReviewRequest].Configure(router, mw)
	handlers[pkg.GetUserActivityOnFilmRequest].Configure(router, mw)

	// Person
	handlers[pkg.PersonRequest].Configure(router, mw)

	// Reviews
	handlers[pkg.ReviewsFilmRequest].Configure(router, mw)

	http.Handle("/", router)

	return router
}
