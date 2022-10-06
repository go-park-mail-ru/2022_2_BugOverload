package database_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errorshandlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

func TestUserStorage(t *testing.T) {
	us := database.NewUserStorage()

	newUser := structs.User{
		ID:       0,
		Nickname: "testNickname",
		Email:    "Test@corp.mail.ru",
	}

	us.Create(newUser)

	ok := us.CheckExist(newUser.Email)
	if !ok {
		t.Errorf("Invalid result: [%t], expected: true", ok)
	}

	usr, err := us.GetUser(newUser.Email)
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	if !cmp.Equal(usr, newUser, cmpopts.IgnoreFields(structs.User{}, "Avatar")) {
		t.Errorf("Err: [%v], expected: [%v]", usr, newUser)
	}
}

func TestUserStorageGet(t *testing.T) {
	us := database.NewUserStorage()

	_, err := us.GetUser("test")
	if err != errorshandlers.ErrUserNotExist {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errorshandlers.ErrUserNotExist)
	}
}
