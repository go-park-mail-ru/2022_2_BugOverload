package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	mockImageRepository "go-park-mail-ru/2022_2_BugOverload/internal/image/repository/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
)

// Handler OK, NOT OK workflow in tests
func TestImageService_GetImage_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockImageRepository.NewMockImageRepository(ctrl)

	// Data
	input := &models.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	output := models.Image{
		Bytes: []byte("some image"),
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetImage(ctx, input).Return(output, nil)

	// Action
	service := service.NewImageService(repository)

	actual, err := service.GetImage(ctx, input)

	// CheckNotificationSent success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNotificationSent result handling
	require.Equal(t, output, actual)
}

func TestImageService_GetImage_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockImageRepository.NewMockImageRepository(ctrl)

	// Data
	input := &models.Image{
		Key:    "2",
		Object: "collection_poster",
	}

	errReturn := stdErrors.New("Some error")

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetImage(ctx, input).Return(models.Image{}, errReturn)

	// Action
	service := service.NewImageService(repository)

	_, errActual := service.GetImage(ctx, input)

	// CheckNotificationSent success
	require.NotNil(t, errActual, "Handling must be error")

	// CheckNotificationSent result handling
	require.Equal(t, errActual.Error(), stdErrors.Wrap(errReturn, "GetImage").Error())
}

func TestImageService_UpdateImage_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockImageRepository.NewMockImageRepository(ctrl)

	// Create required setup for handling
	expected := models.Image{
		Key:    "2",
		Object: "collection_poster",
		Bytes:  []byte("image"),
	}

	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().UpdateImage(ctx, &expected).Return(nil)

	// Action
	service := service.NewImageService(repository)

	err := service.UpdateImage(ctx, &expected)

	// CheckNotificationSent success
	require.Nil(t, err, "Handling must be without errors")
}

func TestImageService_UpdateImage_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockImageRepository.NewMockImageRepository(ctrl)

	// Data
	errReturn := stdErrors.New("Some error")

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().UpdateImage(ctx, &models.Image{}).Return(errReturn)

	// Action
	service := service.NewImageService(repository)

	errActual := service.UpdateImage(ctx, &models.Image{})

	// CheckNotificationSent success
	require.NotNil(t, errActual, "Handling must be error")

	// CheckNotificationSent result handling
	require.Equal(t, errActual.Error(), stdErrors.Wrap(errReturn, "UpdateImage").Error())
}
