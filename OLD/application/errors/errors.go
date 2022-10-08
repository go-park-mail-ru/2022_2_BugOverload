package errors

import (
	"fmt"
)

type DefaultValidationError struct {
	Reason string
	Code   int
}

func (e DefaultValidationError) Error() string {
	return fmt.Sprintf("Def validation: [%s]", e.Reason)
}

func NewErrValidation(err error) DefaultValidationError {
	return DefaultValidationError{
		Reason: err.Error(),
		Code:   ErrCsfValid.GetCode(err),
	}
}

type AuthError struct {
	Reason string
	Code   int
}

func (e AuthError) Error() string {
	return fmt.Sprintf("Action: [%s]", e.Reason)
}

func NewErrAuth(err error) AuthError {
	return AuthError{
		Reason: err.Error(),
		Code:   ErrCsfAuth.GetCode(err),
	}
}

type FilmsError struct {
	Reason string
	Code   int
}

func (e FilmsError) Error() string {
	return fmt.Sprintf("Films: [%s]", e.Reason)
}

func NewErrFilms(err error) FilmsError {
	return FilmsError{
		Reason: err.Error(),
		Code:   ErrCsfFilms.GetCode(err),
	}
}
