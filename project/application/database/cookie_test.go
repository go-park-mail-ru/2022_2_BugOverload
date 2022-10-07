package database_test

import (
	"strings"
	"testing"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/errors"
)

func TestCookieStorage(t *testing.T) {
	cs := database.NewCookieStorage()

	cookie := cs.Create("test@corp.mail.ru")

	if !strings.HasPrefix(cookie, "1=test@corp.mail.ru") {
		t.Errorf("Invalid cookie, [%s]", cookie)
	}

	_, err := cs.DeleteCookie("1=test@corp.mail.ru")
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}
}

func TestCookieStorageDelete(t *testing.T) {
	cs := database.NewCookieStorage()

	_, err := cs.DeleteCookie("")
	if err != errors.ErrCookieNotExist {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist)
	}
}

func TestCookieStorageGet(t *testing.T) {
	cs := database.NewCookieStorage()

	_ = cs.Create("test@mail.ru")

	_, err := cs.GetCookie("1=test@mail.ru")

	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	_, err = cs.GetCookie("")

	if err != errors.NewErrAuth(errors.ErrCookieNotExist) {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist)
	}
}
