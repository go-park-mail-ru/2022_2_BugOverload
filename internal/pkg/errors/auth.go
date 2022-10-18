package errors

import (
	"fmt"
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	ErrEmptyFieldAuth           = stdErrors.New("request has empty fields (nickname | email | password)")
	ErrLoginCombinationNotFound = stdErrors.New("no such combination of login and password")
	ErrUserExist                = stdErrors.New("such user exist")
	ErrUserNotExist             = stdErrors.New("such user doesn't exist")
	ErrSignupUserExist          = stdErrors.New("such a login exists")
	ErrNoCookie                 = stdErrors.New("request has no cookies")
	ErrCookieNotExist           = stdErrors.New("no such cookie")
)

type errClassifierAuth struct {
	table map[error]int
}

func NewErrClassifierAuth() errClassifier {
	res := make(map[error]int)

	res[ErrEmptyFieldAuth] = http.StatusBadRequest
	res[ErrUserExist] = http.StatusBadRequest
	res[ErrUserNotExist] = http.StatusNotFound
	res[ErrSignupUserExist] = http.StatusBadRequest

	res[ErrLoginCombinationNotFound] = http.StatusUnauthorized
	res[ErrNoCookie] = http.StatusUnauthorized
	res[ErrCookieNotExist] = http.StatusNotFound

	return &errClassifierAuth{
		table: res,
	}
}

func (ec *errClassifierAuth) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfAuth = NewErrClassifierAuth()

type AuthError struct {
	Reason string
	Code   int
}

func (e AuthError) Error() string {
	return fmt.Sprintf("Auth: [%s]", e.Reason)
}

func NewErrAuth(err error) AuthError {
	return AuthError{
		Reason: err.Error(),
		Code:   ErrCsfAuth.GetCode(err),
	}
}
