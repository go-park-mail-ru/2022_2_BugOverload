package factories

import (
	authInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/interfaces"
	authService "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	collectionInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/interfaces"
	collectionService "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/service"
	filmsInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/films/interfaces"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/service"
	userInterfaces "go-park-mail-ru/2022_2_BugOverload/internal/app/user/interfaces"
	userService "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
)

func newServiceMap(repo map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	userService_ := userService.NewUserService(repo[userRepo].(userInterfaces.UserRepository), params.ContextTimeout)
	res[userS] = userService_

	authService_ := authService.NewAuthService(repo[authRepo].(authInterfaces.AuthRepository), params.ContextTimeout)
	res[authS] = authService_

	collectionService_ := collectionService.NewCollectionService(repo[collectionRepo].(collectionInterfaces.CollectionRepository), params.ContextTimeout)
	res[collectionS] = collectionService_

	filmsService_ := serviceFilms.NewFilmService(repo[filmsRepo].(filmsInterfaces.FilmsRepository), params.ContextTimeout)
	res[filmsS] = filmsService_

	return res
}
