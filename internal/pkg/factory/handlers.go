package factory

import (
	handlersAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/handlers"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	handlersCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	repoCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	handlersFilm "go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	handlersImage "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	repoImage "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	handlersPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/delivery/handlers"
	repoPerson "go-park-mail-ru/2022_2_BugOverload/internal/person/repository"
	servicePerson "go-park-mail-ru/2022_2_BugOverload/internal/person/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	handlersUser "go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	repoUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func NewHandlersMap(config *pkg.Config, postgres *sqltools.Database, sessionService serviceSession.SessionService, authService serviceAuth.AuthService) map[string]handler.Handler {
	res := make(map[string]handler.Handler)

	authHandler := handlersAuth.NewAuthHandler(authService, sessionService)
	res[pkg.AuthRequest] = authHandler

	logoutHandler := handlersAuth.NewLogoutHandler(authService, sessionService)
	res[pkg.LogoutRequest] = logoutHandler

	loginHandler := handlersAuth.NewLoginHandler(authService, sessionService)
	res[pkg.LoginRequest] = loginHandler

	singUpHandler := handlersAuth.NewSingUpHandler(authService, sessionService)
	res[pkg.SignupRequest] = singUpHandler

	// Collections
	collectionStorage := repoCollection.NewCollectionCache(postgres)

	collectionService := serviceCollection.NewCollectionService(collectionStorage)

	tagCollectionHandler := handlersCollection.NewStdCollectionHandler(collectionService)
	res[pkg.TagCollectionRequest] = tagCollectionHandler

	userCollectionsHandler := handlersCollection.NewGetUserCollectionsHandler(collectionService)
	res[pkg.UserCollectionsRequest] = userCollectionsHandler

	premieresCollectionHandler := handlersCollection.NewPremieresCollectionHandler(collectionService)
	res[pkg.PremieresCollectionRequest] = premieresCollectionHandler

	// Films
	filmsStorage := repoFilms.NewFilmPostgres(postgres)

	filmsService := serviceFilms.NewFilmService(filmsStorage)

	recommendationHandler := handlersFilm.NewRecommendationFilmHandler(filmsService)
	res[pkg.RecommendationRequest] = recommendationHandler

	filmHandler := handlersFilm.NewFilmHandler(filmsService)
	res[pkg.FilmRequest] = filmHandler

	// Reviews
	reviewsHandler := handlersFilm.NewReviewsHandler(filmsService)
	res[pkg.ReviewsFilmRequest] = reviewsHandler

	// Images
	is := repoImage.NewImageS3(config, postgres)

	imageService := serviceImage.NewImageService(is)

	downloadImageHandler := handlersImage.NewGetImageHandler(imageService)
	res[pkg.DownloadImageRequest] = downloadImageHandler

	changeImageHandler := handlersImage.NewPutImageHandler(imageService)
	res[pkg.ChangeImageRequest] = changeImageHandler

	uploadImageHandler := handlersImage.NewPostImageHandler(imageService)
	res[pkg.UploadImageRequest] = uploadImageHandler

	// Users
	userRepo := repoUser.NewUserPostgres(postgres)

	userService := serviceUser.NewUserProfileService(userRepo, authService)

	profileHandler := handlersUser.NewUserProfileHandler(userService)
	res[pkg.GetUserProfileRequest] = profileHandler

	userSettingsHandler := handlersUser.NewGetSettingsHandler(userService)
	res[pkg.GetUserSettingsRequest] = userSettingsHandler

	changeUserSettingsHandler := handlersUser.NewPutSettingsHandler(userService)
	res[pkg.PutUserSettingsRequest] = changeUserSettingsHandler

	filmRateHandler := handlersUser.NewFilmRateHandler(userService)
	res[pkg.FilmRateRequest] = filmRateHandler

	filmRateDropHandler := handlersUser.NewFilmRateDropHandler(userService)
	res[pkg.FilmRateDropRequest] = filmRateDropHandler

	newFilmReviewHandler := handlersUser.NewFilmReviewHandler(userService)
	res[pkg.NewFilmReviewRequest] = newFilmReviewHandler

	getUserActivityOnFilmHandler := handlersUser.NewGetActivityOnFilmHandler(userService)
	res[pkg.GetUserActivityOnFilmRequest] = getUserActivityOnFilmHandler

	// Persons
	personRepo := repoPerson.NewPersonPostgres(postgres)

	personService := servicePerson.NewPersonService(personRepo)

	personHandler := handlersPerson.NewPersonHandler(personService)
	res[pkg.PersonRequest] = personHandler

	return res
}
