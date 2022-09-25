package tmp_storage

import (
	"errors"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

type UserStorage struct {
	Storage map[structs.User]string
}

func NewUserStorage() UserStorage {
	return UserStorage{make(map[structs.User]string)}
}

func (us *UserStorage) Insert(u structs.User) error {
	if len(us.Storage) == 0 {
		us.Storage[u] = "exist"

		return nil
	}

	_, ok := us.Storage[u]
	if ok {
		return errors.New("such user exist")
	}

	us.Storage[u] = "exist"

	return nil
}
