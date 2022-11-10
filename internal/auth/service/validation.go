package service

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/mail"
)

const minPasswordLength = 6

func ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.ErrInvalidEmail
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < minPasswordLength {
		return errors.ErrInvalidPassword
	}
	return nil
}
