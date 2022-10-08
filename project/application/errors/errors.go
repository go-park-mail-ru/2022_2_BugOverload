package errors

import (
	"fmt"
)

type ErrHTTP struct {
	Reason string
	Code   int
}

func (e ErrHTTP) Error() string {
	return fmt.Sprintf("HTTP: [%s]", e.Reason)
}

func NewErrHTTP(error error) ErrHTTP {
	return ErrHTTP{
		Reason: error.Error(),
		Code:   ErrCsfHTTP.GetCode(error),
	}
}

type ErrAuth struct {
	Reason string
	Code   int
}

func (e ErrAuth) Error() string {
	return fmt.Sprintf("Auth: [%s]", e.Reason)
}

func NewErrAuth(error error) ErrAuth {
	return ErrAuth{
		Reason: error.Error(),
		Code:   ErrCsfAuth.GetCode(error),
	}
}

type ErrFilms struct {
	Reason string
	Code   int
}

func (e ErrFilms) Error() string {
	return fmt.Sprintf("Films: [%s]", e.Reason)
}

func NewErrFilms(error error) ErrFilms {
	return ErrFilms{
		Reason: error.Error(),
		Code:   ErrCsfFilms.GetCode(error),
	}
}
