package OLDTESTS_test

//
//import (
//	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
//	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
//	"strings"
//	"testing"
//
//	stdErrors "github.com/pkg/errors"
//)
//
//func TestCookieStorage(t *testing.T) {
//	cs := memory.NewCookieRepo()
//
//	cookie := cs.CreateSession("test@corp.mail.ru")
//
//	if !strings.HasPrefix(cookie, "1=test@corp.mail.ru") {
//		t.Errorf("Invalid cookie, [%s]", cookie)
//	}
//
//	_, err := cs.DeleteSession("1=test@corp.mail.ru")
//	if err != nil {
//		t.Errorf("Err: [%s], expected: nil", err.Error())
//	}
//}
//
//func TestCookieStorageDelete(t *testing.T) {
//	cs := memory.NewCookieRepo()
//
//	_, err := cs.DeleteSession("")
//	if !stdErrors.Is(err, errors.ErrCookieNotExist) {
//		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist.Error())
//	}
//}
//
//func TestCookieStorageGet(t *testing.T) {
//	cs := memory.NewCookieRepo()
//
//	_ = cs.CreateSession("test@mail.ru")
//
//	_, err := cs.GetSession("1=test@mail.ru")
//
//	if err != nil {
//		t.Errorf("Err: [%s], expected: nil", err.Error())
//	}
//
//	_, err = cs.GetSession("")
//
//	if !stdErrors.Is(err, errors.ErrCookieNotExist) {
//		t.Errorf("Err: [%s], expected: [%s]", err.Error(), errors.ErrCookieNotExist)
//	}
//}
