package factories

import (
	authInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popularfilmshandler"
	collectionInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendationfilmhandler"
	filmsInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/interfaces"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/authhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/loginhandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/logouthandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/signuphandler"
	userInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
)

func NewHandlersMap(repo map[string]interface{}) map[string]interfaces.Handler {
	res := make(map[string]interfaces.Handler)

	authHandler := authhandler.NewHandler(res[userS].(userInterfaces.UserService), res[authS].(authInterfaces.AuthRepository))

	logoutHandler := logouthandler.NewHandler(res[userS].(userInterfaces.UserService), res[authS].(authInterfaces.AuthRepository))

	loginHandler := loginhandler.NewHandler(res[userS].(userInterfaces.UserService), res[authS].(authInterfaces.AuthRepository))

	singUpHandler := signuphandler.NewHandler(res[userS].(userInterfaces.UserService), res[authS].(authInterfaces.AuthRepository))

	inCinemaHandler := incinemafilmshandler.NewHandler(res[collectionS].(collectionInterfaces.CollectionRepository))
	
	popularHandler := popularfilmshandler.NewHandler(res[collectionS].(collectionInterfaces.CollectionRepository))

	recommendationHandler := recommendationfilmhandler.NewHandler(res[filmsS].(filmsInterfaces.FilmsRepository), res[authS].(authInterfaces.AuthRepository))

	return res
}
