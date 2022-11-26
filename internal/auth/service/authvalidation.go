package service

import (
	"net/mail"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

const minNicknameLength = 4
const minPasswordLength = 6

func ValidateNickname(nickname string) error {
	if len(nickname) < minNicknameLength {
		return errors.ErrInvalidNickname
	}
	return nil
}

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
