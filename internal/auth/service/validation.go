package service

import (
	"net/mail"
	"unicode"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

const minPasswordLength = 8

func ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.ErrInvalidEmail
	}
	return nil
}

func ValidateNickname(username string) error {
	for _, char := range username {
		if !(unicode.IsLetter(char) || unicode.Is(unicode.Cyrillic, char)) {
			return errors.ErrInvalidNickname
		}
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < minPasswordLength {
		return errors.ErrInvalidPassword
	}
	return nil
}
