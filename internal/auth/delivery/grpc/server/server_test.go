package server_test

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/server"
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoModels "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/models"
	mockAuthService "go-park-mail-ru/2022_2_BugOverload/internal/auth/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_Auth_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		ID: 12,
	}

	outputService := models.User{
		ID:       12,
		Nickname: "testnickname",
		Email:    "testemail@example.com",
		Avatar:   "avatar",
	}

	input := protoModels.NewUserProto(inputService)

	expected := protoModels.NewUserProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().Auth(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.Auth(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_Auth_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		ID: 0,
	}

	outputServiceErr := errors.ErrUserNotFound

	input := protoModels.NewUserProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrUserNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	authService.EXPECT().Auth(ctx, inputService).Return(models.User{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.Auth(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_Login_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		Email:    "testemail@expample.com",
		Password: "testpassword",
	}

	outputService := models.User{
		ID:       12,
		Nickname: "testnickname",
		Email:    "testemail@example.com",
		Avatar:   "avatar",
	}

	input := protoModels.NewUserProto(inputService)

	expected := protoModels.NewUserProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().Login(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.Login(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_Login_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		Email:    "testemail@expample.com",
		Password: "testpassword123",
	}

	outputServiceErr := errors.ErrIncorrectPassword

	input := protoModels.NewUserProto(inputService)

	expectedErr := status.Error(codes.PermissionDenied, errors.ErrIncorrectPassword.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	authService.EXPECT().Login(ctx, inputService).Return(models.User{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.Login(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_Signup_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		Email:    "testemail@expample.com",
		Nickname: "testnickname",
		Password: "testpassword",
	}

	outputService := models.User{
		ID:       12,
		Nickname: "testnickname",
		Email:    "testemail@example.com",
		Avatar:   "avatar",
	}

	input := protoModels.NewUserProto(inputService)

	expected := protoModels.NewUserProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().Signup(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.Signup(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_Signup_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		Email:    "testemail@expample.com",
		Password: "testpassword123",
	}

	outputServiceErr := errors.ErrUserExist

	input := protoModels.NewUserProto(inputService)

	expectedErr := status.Error(codes.InvalidArgument, errors.ErrUserExist.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	authService.EXPECT().Signup(ctx, inputService).Return(models.User{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.Signup(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_GetAccess_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputPasswordService := "testpassword"

	input := protoModels.NewGetAccessParamsProto(inputUserService, inputPasswordService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().GetAccess(ctx, inputUserService, inputPasswordService).Return(nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, err := authServer.GetAccess(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")
}

func TestAuthServiceGRPCServer_GetAccess_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputPasswordService := "testpassword123"

	outputServiceErr := errors.ErrIncorrectPassword

	input := protoModels.NewGetAccessParamsProto(inputUserService, inputPasswordService)

	expectedErr := status.Error(codes.PermissionDenied, outputServiceErr.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	authService.EXPECT().GetAccess(ctx, inputUserService, inputPasswordService).Return(outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.GetAccess(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_UpdatePassword_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputOldPasswordService := "testpassword"
	inputNewPasswordService := "testpassword123"

	input := protoModels.NewUpdatePasswordParamsProto(inputUserService, inputOldPasswordService, inputNewPasswordService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().UpdatePassword(ctx, inputUserService, inputOldPasswordService, inputNewPasswordService).Return(nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, err := authServer.UpdatePassword(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")
}

func TestAuthServiceGRPCServer_UpdatePassword_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputOldPasswordService := "testpassword"
	inputNewPasswordService := "testpassword123"

	outputServiceErr := errors.ErrIncorrectPassword

	input := protoModels.NewUpdatePasswordParamsProto(inputUserService, inputOldPasswordService, inputNewPasswordService)

	expectedErr := status.Error(codes.PermissionDenied, outputServiceErr.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	authService.EXPECT().UpdatePassword(ctx, inputUserService, inputOldPasswordService, inputNewPasswordService).Return(outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.UpdatePassword(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_GetUserBySession_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.Session{
		ID: "CorrectSessionID",
		User: &models.User{
			ID: 12,
		},
	}

	outputService := models.User{
		ID:       12,
		Nickname: "testnickname",
		Email:    "testemail@example.com",
		Avatar:   "avatar",
	}

	input := protoModels.NewSessionProto(inputService)

	expected := protoModels.NewUserProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	sessionService.EXPECT().GetUserBySession(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.GetUserBySession(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_GetUserBySession_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.Session{
		ID: "IncorrectSessionID",
		User: &models.User{
			ID: 0,
		},
	}

	outputServiceErr := errors.ErrSessionNotFound

	input := protoModels.NewSessionProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrSessionNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	sessionService.EXPECT().GetUserBySession(ctx, inputService).Return(models.User{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.GetUserBySession(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_CreateSession_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		ID:       12,
		Email:    "testemail@expample.com",
		Nickname: "testnickname",
	}

	outputService := models.Session{
		ID: "CorrectSessionID",
		User: &models.User{
			ID: 12,
		},
	}

	input := protoModels.NewUserProto(inputService)

	expected := protoModels.NewSessionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	sessionService.EXPECT().CreateSession(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.CreateSession(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_CreateSession_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.User{
		ID: 12,
	}

	outputServiceErr := errors.ErrUserNotExist

	input := protoModels.NewUserProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrUserNotExist.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	sessionService.EXPECT().CreateSession(ctx, inputService).Return(models.Session{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.CreateSession(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestAuthServiceGRPCServer_DeleteSession_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.Session{
		ID: "CorrectSessionID",
		User: &models.User{
			ID: 12,
		},
	}

	outputService := models.Session{
		ID: "IncorrectSessionID",
		User: &models.User{
			ID: 0,
		},
	}

	input := protoModels.NewSessionProto(inputService)

	expected := protoModels.NewSessionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	sessionService.EXPECT().DeleteSession(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	actual, err := authServer.DeleteSession(ctx, input)

	// CheckNewNotification success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNewNotification result handling
	require.Equal(t, expected, actual)
}

func TestAuthServiceGRPCServer_DeleteSession_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthService.NewMockAuthService(ctrl)
	sessionService := mockAuthService.NewMockSessionService(ctrl)

	// Data
	inputService := &models.Session{
		ID: "IncorrectSessionID",
		User: &models.User{
			ID: 0,
		},
	}

	outputServiceErr := errors.ErrSessionNotFound

	input := protoModels.NewSessionProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrSessionNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	sessionService.EXPECT().DeleteSession(ctx, inputService).Return(models.Session{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	authServer := server.NewAuthServiceGRPCServer(grpcServer, authService, sessionService)

	// Action
	_, actualErr := authServer.DeleteSession(ctx, input)

	// CheckNewNotification success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNewNotification result handling
	require.Equal(t, expectedErr, actualErr)
}
