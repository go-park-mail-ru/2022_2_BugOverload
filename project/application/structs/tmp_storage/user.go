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

func (us *UserStorage) Contains(u structs.User) error {
	_, ok := us.Storage[u]
	if ok {
		return errors.New("such user exist")
	}

	return nil
}

func (us *UserStorage) Insert(u structs.User) {
	us.Storage[u] = "exist"

}
