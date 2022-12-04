package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	mockFilmRepository "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/film/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
)

func TestFilmService_GetRecommendation_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	output := models.Film{
		ID:   1,
		Name: "test",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetRecommendation(ctx).Return(output, nil)

	// Action
	filmService := service.NewFilmService(repository)

	actual, err := filmService.GetRecommendation(ctx)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestFilmService_GetRecommendation_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	expectedErr := errors.ErrNotFoundInDB

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetRecommendation(ctx).Return(models.Film{}, expectedErr)

	// Action
	filmService := service.NewFilmService(repository)

	_, actualErr := filmService.GetRecommendation(ctx)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestFilmService_GetFilmByID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	inputFilm := &models.Film{
		ID: 1,
	}

	inputParams := &constparams.GetFilmParams{
		CountImages: 1,
		CountActors: 1,
	}

	output := models.Film{
		ID:   1,
		Name: "test",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetFilmByID(ctx, inputFilm, inputParams).Return(output, nil)

	// Action
	filmService := service.NewFilmService(repository)

	actual, err := filmService.GetFilmByID(ctx, inputFilm, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestFilmService_GetFilmByID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	inputFilm := &models.Film{
		ID: 0,
	}

	inputParams := &constparams.GetFilmParams{
		CountImages: 1,
		CountActors: 1,
	}

	expectedErr := errors.ErrFilmNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetFilmByID(ctx, inputFilm, inputParams).Return(models.Film{}, expectedErr)

	// Action
	filmService := service.NewFilmService(repository)

	_, actualErr := filmService.GetFilmByID(ctx, inputFilm, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestFilmService_GetReviewsByFilmID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetFilmReviewsParams{
		FilmID:       1,
		CountReviews: 10,
		Offset:       0,
	}

	output := []models.Review{
		{
			ID:   1,
			Body: "test body",
		},
		{
			ID:   2,
			Body: "another test body",
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetReviewsByFilmID(ctx, inputParams).Return(output, nil)

	// Action
	filmService := service.NewFilmService(repository)

	actual, err := filmService.GetReviewsByFilmID(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestFilmService_GetReviewsByFilmID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockFilmRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetFilmReviewsParams{
		FilmID:       0,
		CountReviews: 10,
		Offset:       0,
	}

	expectedErr := errors.ErrNotFoundInDB

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetReviewsByFilmID(ctx, inputParams).Return([]models.Review{}, expectedErr)

	// Action
	filmService := service.NewFilmService(repository)

	_, actualErr := filmService.GetReviewsByFilmID(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
