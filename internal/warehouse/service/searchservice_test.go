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
	mockSearchRepository "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/search/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
)

func TestSearchService_Search_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSearchRepository.NewMockRepository(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "correct_query",
	}

	input := &constparams.SearchParams{
		Query: "%correct_query%",
	}

	outputFilms := []models.Film{}
	outputSerials := []models.Film{}
	outputPersons := []models.Person{}

	outputService := models.Search{
		Films:   outputFilms,
		Serials: outputSerials,
		Persons: outputPersons,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().SearchFilms(ctx, input).Return(outputFilms, nil)
	repository.EXPECT().SearchSeries(ctx, input).Return(outputSerials, nil)
	repository.EXPECT().SearchPersons(ctx, input).Return(outputPersons, nil)

	// Action
	searchService := service.NewSearchService(repository)

	actual, err := searchService.Search(ctx, inputService)

	// CheckNotificationSent success
	require.Nil(t, err, "Handling must be without errors")

	// CheckNotificationSent result handling
	require.Equal(t, outputService, actual)
}

func TestSearchService_Search_FilmsErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSearchRepository.NewMockRepository(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "correct_query",
	}

	input := &constparams.SearchParams{
		Query: "%correct_query%",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().SearchFilms(ctx, input).Return([]models.Film{}, expectedErr)

	// Action
	searchService := service.NewSearchService(repository)

	_, actualErr := searchService.Search(ctx, inputService)

	// CheckNotificationSent success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNotificationSent result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestSearchService_Search_SerialsErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSearchRepository.NewMockRepository(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "correct_query",
	}

	input := &constparams.SearchParams{
		Query: "%correct_query%",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().SearchFilms(ctx, input).Return([]models.Film{}, nil)
	repository.EXPECT().SearchSeries(ctx, input).Return([]models.Film{}, expectedErr)

	// Action
	searchService := service.NewSearchService(repository)

	_, actualErr := searchService.Search(ctx, inputService)

	// CheckNotificationSent success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNotificationSent result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestSearchService_Search_PersonsErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockSearchRepository.NewMockRepository(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "correct_query",
	}

	input := &constparams.SearchParams{
		Query: "%correct_query%",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().SearchFilms(ctx, input).Return([]models.Film{}, nil)
	repository.EXPECT().SearchSeries(ctx, input).Return([]models.Film{}, nil)
	repository.EXPECT().SearchPersons(ctx, input).Return([]models.Person{}, expectedErr)

	// Action
	searchService := service.NewSearchService(repository)

	_, actualErr := searchService.Search(ctx, inputService)

	// CheckNotificationSent success
	require.NotNil(t, actualErr, "Handling must be error")

	// CheckNotificationSent result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
