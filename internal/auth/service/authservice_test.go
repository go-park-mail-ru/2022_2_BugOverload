package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	mockAuthRepository "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/auth/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/security"
)

func TestAuthService_Auth_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 12,
	}

	output := models.User{
		ID:       12,
		Nickname: "test",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(output, nil)

	// Action
	authService := service.NewAuthService(repository)

	actual, err := authService.Auth(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, output, actual)
}

func TestAuthService_Auth_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(models.User{}, expectedErr)

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Auth(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Login_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       12,
		Email:    "testemail@gmail.com",
		Password: "testpassword123",
	}

	outputPassword, _ := security.HashPassword(input.Password)

	output := models.User{
		ID:       12,
		Nickname: "test",
		Password: outputPassword,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByEmail(ctx, input.Email).Return(output, nil)

	// Action
	authService := service.NewAuthService(repository)

	actual, err := authService.Login(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, output, actual)
}

func TestAuthService_Login_EmailValidate_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrInvalidEmail

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Login(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Login_PasswordValidate_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:    0,
		Email: "testemail123@mail.ru",
	}

	expectedErr := errors.ErrInvalidPassword

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Login(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Login_Repo_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Email:    "invalidemail@gmail.com",
		Password: "testpassword123",
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByEmail(ctx, input.Email).Return(models.User{}, expectedErr)

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Login(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Login_PasswordCheck_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Email:    "correctemail@mail.ru",
		Password: "incorrect_password",
	}
	output := models.User{
		Password: "correct_password",
	}

	expectedErr := errors.ErrIncorrectPassword

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByEmail(ctx, input.Email).Return(output, nil)

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Login(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_GetAccess_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	inputPassword := "testpassword123"
	input := &models.User{
		ID:       12,
		Password: inputPassword,
	}

	hashedPassword, _ := security.HashPassword(inputPassword)
	output := models.User{
		ID:       12,
		Password: hashedPassword,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(output, nil)

	// Action
	authService := service.NewAuthService(repository)

	err := authService.GetAccess(ctx, input, inputPassword)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")
}

func TestAuthService_GetAccess_Repo_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	inputPassword := "testpassword123"
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(models.User{}, expectedErr)

	// Action
	authService := service.NewAuthService(repository)

	actualErr := authService.GetAccess(ctx, input, inputPassword)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_GetAccess_Password_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	inputPassword := "testpassword123"
	input := &models.User{
		ID:       0,
		Password: "incorrectpassword",
	}

	hashedPassword, _ := security.HashPassword(input.Password)
	output := models.User{
		ID:       12,
		Password: hashedPassword,
	}

	expectedErr := errors.ErrIncorrectPassword

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(output, nil)

	// Action
	authService := service.NewAuthService(repository)

	actualErr := authService.GetAccess(ctx, input, inputPassword)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Signup_NicknameValidate_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrInvalidNickname

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Signup_EmailValidate_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Nickname: "testnickname123",
		Email:    "invalidmail@@liast.ruz",
	}

	expectedErr := errors.ErrInvalidEmail

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Signup_PasswordValidate_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Nickname: "testnickname123",
		Email:    "validmail@mail.ru",
		Password: "123",
	}

	expectedErr := errors.ErrInvalidPassword

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Signup_CheckExistErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Nickname: "testnickname123",
		Email:    "validmail@mail.ru",
		Password: "correctpassword123",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckExistUserByEmail(ctx, input.Email).Return(false, expectedErr)

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_Signup_UserExist(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID:       0,
		Nickname: "testnickname123",
		Email:    "validmail@mail.ru",
		Password: "correctpassword123",
	}

	expectedErr := errors.ErrUserExist

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckExistUserByEmail(ctx, input.Email).Return(true, nil)

	// Action
	authService := service.NewAuthService(repository)

	_, actualErr := authService.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestAuthService_UpdatePassword_AccessErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockAuthRepository.NewMockRepository(ctrl)

	// Data
	inputOldPassword := "testpassword123"
	inputNewPassword := "testpassword"
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserByID(ctx, input.ID).Return(models.User{}, expectedErr)

	// Action
	authService := service.NewAuthService(repository)

	actualErr := authService.UpdatePassword(ctx, input, inputOldPassword, inputNewPassword)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
