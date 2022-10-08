package memory_test

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"strings"
	"testing"

	stdErrors "github.com/pkg/errors"
)

func TestCookieStorage(t *testing.T) {
	cs := memory.NewCookieStorage()

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
	cs := memory.NewCookieStorage()

	_, err := cs.DeleteCookie("")
	if !stdErrors.Is(err, errors.ErrCookieNotExist) {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist.Error())
	}
}

func TestCookieStorageGet(t *testing.T) {
	cs := memory.NewCookieStorage()

	_ = cs.Create("test@mail.ru")

	_, err := cs.GetCookie("1=test@mail.ru")

	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err.Error())
	}

	_, err = cs.GetCookie("")

	if !stdErrors.Is(err, errors.ErrCookieNotExist) {
		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist)
	}
}
