package factories

import (
	"sync"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/repository/memory"
	memoryFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
)

func newRepoMap() map[string]interface{} {
	res := make(map[string]interface{})

	userMutex := &sync.Mutex{}
	userS := memoryUser.NewUserRepo(userMutex)
	res[userRepo] = userS

	authMutex := &sync.Mutex{}
	cookieS := memoryCookie.NewCookieRepo(authMutex)
	res[authRepo] = cookieS

	pathInCinema := "test/testdata/incinema.json"
	pathPopular := "test/testdata/popular.json"
	collectionMutex := &sync.Mutex{}
	colS := memoryCollection.NewCollectionRepo(collectionMutex, pathPopular, pathInCinema)
	res[collectionRepo] = colS

	pathPreview := "test/testdata/preview.json"
	filmsMutex := &sync.Mutex{}
	filmsS := memoryFilms.NewFilmRepo(filmsMutex, pathPreview)
	res[filmsRepo] = filmsS

	return res
}
