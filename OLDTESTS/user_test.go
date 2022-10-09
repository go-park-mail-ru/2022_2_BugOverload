package OLDTESTS_test

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"testing"

	stdErrors "github.com/pkg/errors"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUserStorage(t *testing.T) {
	us := memory.NewUserRepo()

	newUser := models.User{
		ID:       0,
		Nickname: "testNickname",
		Email:    "Test@corp.mail.ru",
	}

	us.Signup(newUser)

	ok := us.CheckExist(newUser.Email)
	if !ok {
		t.Errorf("Invalid result: [%t], expected: true", ok)
	}

	usr, err := us.Login(newUser.Email)
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	if !cmp.Equal(usr, newUser, cmpopts.IgnoreFields(models.User{}, "Avatar")) {
		t.Errorf("Err: [%v], expected: [%v]", usr, newUser)
	}
}

func TestUserStorageGet(t *testing.T) {
	us := memory.NewUserRepo()

	_, err := us.Login("test")
	if !stdErrors.Is(err, errors.ErrUserNotExist) {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrUserNotExist)
	}
}
