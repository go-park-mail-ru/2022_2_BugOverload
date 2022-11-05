package errors

import (
	"fmt"
)

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
		Code:   ErrCsf.GetCode(err),
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
		Code:   ErrCsf.GetCode(err),
	}
}

type ImagesError struct {
	Reason string
	Code   int
}

func (e ImagesError) Error() string {
	return fmt.Sprintf("Image: [%s]", e.Reason)
}

func NewErrImages(err error) ImagesError {
	return ImagesError{
		Reason: err.Error(),
		Code:   ErrCsf.GetCode(err),
	}
}

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
		Code:   ErrCsf.GetCode(err),
	}
}

type PersonError struct {
	Reason string
	Code   int
}

func (e PersonError) Error() string {
	return fmt.Sprintf("Person: [%s]", e.Reason)
}

func NewErrPerson(err error) PersonError {
	return PersonError{
		Reason: err.Error(),
		Code:   ErrCsf.GetCode(err),
	}
}

type ProfileError struct {
	Reason string
	Code   int
}

func (e ProfileError) Error() string {
	return fmt.Sprintf("Profile: [%s]", e.Reason)
}

func NewErrProfile(err error) ProfileError {
	return ProfileError{
		Reason: err.Error(),
		Code:   ErrCsf.GetCode(err),
	}
}

type CollectionError struct {
	Reason string
	Code   int
}

func (e CollectionError) Error() string {
	return fmt.Sprintf("Collection: [%s]", e.Reason)
}

func NewErrCollection(err error) CollectionError {
	return CollectionError{
		Reason: err.Error(),
		Code:   ErrCsf.GetCode(err),
	}
}
