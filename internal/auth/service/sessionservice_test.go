package service_test

import (
	"context"
	stdErrors "github.com/pkg/errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockSessionRepository "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository/session/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

func TestSessionService_CreateSession_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSessionRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 123,
	}

	output := models.Session{
		ID: "session_id",
		User: &models.User{
			ID: 123,
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CreateSession(ctx, input).Return(output, nil)

	// Action
	sessionService := service.NewSessionService(repository)

	actual, err := sessionService.CreateSession(ctx, input)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestSessionService_CreateSession_NOT_OK_ID(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSessionRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotExist

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	sessionService := service.NewSessionService(repository)

	_, actualErr := sessionService.CreateSession(ctx, input)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestSessionService_CreateSession_NOT_OK_Repository(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSessionRepository.NewMockRepository(ctrl)

	// Data
	input := &models.User{
		ID: 1,
	}

	expectedErr := errors.ErrCreateSession

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CreateSession(ctx, input).Return(models.Session{}, expectedErr)

	// Action
	sessionService := service.NewSessionService(repository)

	_, actualErr := sessionService.CreateSession(ctx, input)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
