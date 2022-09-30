package auth

import "go-park-mail-ru/2022_2_BugOverload/project/application/database"

// HandlerAuth is structure for API auth, login and signup processing
type HandlerAuth struct {
	userStorage   *database.UserStorage
	cookieStorage *database.CookieStorage
	//  Менеджер кеша
	//  Логер
	//  Менеджер моделей
}

// NewHandlerAuth is constructor for HandlerAuth
func NewHandlerAuth(us *database.UserStorage, cs *database.CookieStorage) *HandlerAuth {
	return &HandlerAuth{us, cs}
}
