package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

func TestAuthValidation_ValidateNickname_OK(t *testing.T) {
	input := "CorrectNickname123"

	actualErr := service.ValidateNickname(input)

	require.Nil(t, actualErr)
}

func TestAuthValidation_ValidateNickname_NOT_OK(t *testing.T) {
	input := ""
	expectedErr := errors.ErrInvalidNickname

	actualErr := service.ValidateNickname(input)

	require.NotNil(t, actualErr)

	require.Equal(t, expectedErr, actualErr)
}

func TestAuthValidation_ValidateEmail_OK(t *testing.T) {
	input := "correctmail@mail.ru"

	actualErr := service.ValidateEmail(input)

	require.Nil(t, actualErr)
}

func TestAuthValidation_ValidateEmail_NOT_OK(t *testing.T) {
	input := "asda@@lafs.qq"
	expectedErr := errors.ErrInvalidEmail

	actualErr := service.ValidateEmail(input)

	require.NotNil(t, actualErr)

	require.Equal(t, expectedErr, actualErr)
}

func TestAuthValidation_ValidatePassword_OK(t *testing.T) {
	input := "correct_password_123"

	actualErr := service.ValidatePassword(input)

	require.Nil(t, actualErr)
}

func TestAuthValidation_ValidatePassword_NOT_OK(t *testing.T) {
	input := "1234"
	expectedErr := errors.ErrInvalidPassword

	actualErr := service.ValidatePassword(input)

	require.NotNil(t, actualErr)

	require.Equal(t, expectedErr, actualErr)
}
