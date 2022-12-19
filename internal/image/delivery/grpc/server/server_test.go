package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/server"
	mockImageService "go-park-mail-ru/2022_2_BugOverload/internal/image/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// Handler OK, NOT OK workflow in tests
func TestImageServiceServer_GetImage_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageService.NewMockImageService(ctrl)

	// Data
	inputService := &models.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	outputService := models.Image{
		Bytes: []byte("some image"),
	}

	input := &proto.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	expected := &proto.Image{
		Bytes: []byte("some image"),
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	service.EXPECT().GetImage(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	server := server.NewImageServiceGRPCServer(grpcServer, service)

	// Action
	actual, err := server.GetImage(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestImageServiceServer_GetImage_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageService.NewMockImageService(ctrl)

	// Data
	inputService := &models.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	outputServiceErr := errors.ErrImageNotFound

	input := &proto.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	expectedErr := status.Error(codes.NotFound, errors.ErrImageNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	service.EXPECT().GetImage(ctx, inputService).Return(models.Image{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	server := server.NewImageServiceGRPCServer(grpcServer, service)

	// Action
	_, actualErr := server.GetImage(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

func TestImageServiceServer_UpdateImage_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageService.NewMockImageService(ctrl)

	// Data
	inputService := &models.Image{
		Key:    "2",
		Object: "collection_poster",
		Bytes:  []byte("some image"),
	}

	input := &proto.Image{
		Key:    "2",
		Object: "collection_poster",
		Bytes:  []byte("some image"),
	}

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	service.EXPECT().UpdateImage(ctx, inputService).Return(nil)

	// Init
	grpcServer := grpc.NewServer()

	server := server.NewImageServiceGRPCServer(grpcServer, service)

	// Action
	_, err := server.UpdateImage(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")
}

func TestImageServiceServer_UpdateImage_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageService.NewMockImageService(ctrl)

	// Data
	inputService := &models.Image{
		Key:    "2",
		Object: "collection_poster",
		Bytes:  []byte("some image"),
	}

	outputServiceErr := errors.ErrWorkDatabase

	input := &proto.Image{
		Key:    "2",
		Object: "collection_poster",
		Bytes:  []byte("some image"),
	}

	expectedErr := status.Error(codes.Internal, errors.ErrWorkDatabase.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	service.EXPECT().UpdateImage(ctx, inputService).Return(outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	server := server.NewImageServiceGRPCServer(grpcServer, service)

	// Action
	_, actualErr := server.UpdateImage(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}
