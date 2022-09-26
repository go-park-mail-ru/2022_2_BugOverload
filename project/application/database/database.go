package database

type Database struct {
	UserStorage   *UserStorage
	CookieStorage *CookieStorage
}

func NewDatabase() *Database {
	return &Database{
		NewUserStorage(),
		NewCookieStorage(),
	}
}
