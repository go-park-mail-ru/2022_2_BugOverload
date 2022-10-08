package errors

import (
	"fmt"
)

type ErrDefaultValidation struct {
	Reason string
	Code   int
}

func (e ErrDefaultValidation) Error() string {
	return fmt.Sprintf("Def validation: [%s]", e.Reason)
}

func NewErrValidation(error error) ErrDefaultValidation {
	return ErrDefaultValidation{
		Reason: error.Error(),
		Code:   ErrCsfValid.GetCode(error),
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
