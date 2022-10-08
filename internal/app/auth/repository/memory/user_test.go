package memory_test

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"testing"

	stdErrors "github.com/pkg/errors"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUserStorage(t *testing.T) {
	us := memory.NewUserStorage()

	newUser := models.User{
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

	if !cmp.Equal(usr, newUser, cmpopts.IgnoreFields(models.User{}, "Avatar")) {
		t.Errorf("Err: [%v], expected: [%v]", usr, newUser)
	}
}

func TestUserStorageGet(t *testing.T) {
	us := memory.NewUserStorage()

	_, err := us.GetUser("test")
	if !stdErrors.Is(err, errors.ErrUserNotExist) {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrUserNotExist)
	}
}
