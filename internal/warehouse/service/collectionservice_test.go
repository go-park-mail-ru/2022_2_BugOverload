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
	mockCollectionRepository "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/repository/collection/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service"
)

func TestCollectionService_GetCollectionByTag_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularFrom,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularIn,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	output := models.Collection{
		ID:   1,
		Name: "test collection name",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByTag(ctx, inputParamsRepository).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetCollectionByTag(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetCollectionByTag_TagErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        "unknown",
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrTagNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionByTag(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionByTag_RepoErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularFrom,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularIn,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrFilmsNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByTag(ctx, inputParamsRepository).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionByTag(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionByGenre_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyFrom,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyIn,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	output := models.Collection{
		ID:   1,
		Name: "test collection name",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByGenre(ctx, inputParamsRepository).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetCollectionByGenre(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetCollectionByGenre_GenreErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        "unknown",
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrGenreNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionByGenre(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionByGenre_RepoErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyFrom,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyIn,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrFilmsNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByGenre(ctx, inputParamsRepository).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionByGenre(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetStdCollection_Tag_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularFrom,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.PopularIn,
		Target:     "tag",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	output := models.Collection{
		ID:   1,
		Name: "test collection name",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByTag(ctx, inputParamsRepository).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetStdCollection(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetStdCollection_Genre_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyFrom,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyIn,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	output := models.Collection{
		ID:   1,
		Name: "test collection name",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByGenre(ctx, inputParamsRepository).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetStdCollection(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetStdCollection_TargetErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyFrom,
		Target:     "unknown",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrNotFindSuchTarget

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetStdCollection(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetStdCollection_RepoErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyFrom,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	inputParamsRepository := &constparams.GetStdCollectionParams{
		Key:        constparams.ComedyIn,
		Target:     "genre",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	expectedErr := errors.ErrFilmsNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetCollectionByGenre(ctx, inputParamsRepository).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetStdCollection(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetPremieresCollection_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetPremiersCollectionParams{
		CountFilms: 10,
		Delimiter:  0,
	}

	output := models.Collection{
		ID:   1,
		Name: "Премьеры",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetPremieresCollection(ctx, inputParams).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetPremieresCollection(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetPremieresCollection_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.GetPremiersCollectionParams{
		CountFilms: 10,
		Delimiter:  0,
	}

	expectedErr := errors.ErrNotFoundInDB

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetPremieresCollection(ctx, inputParams).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetPremieresCollection(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionAuthorized_Author_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	output := models.Collection{
		ID:   1,
		Name: "Премьеры",
		Author: models.User{
			ID: 1,
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().GetCollection(ctx, inputParams).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetCollectionAuthorized_NotAuthor_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	output := models.Collection{
		ID:   1,
		Name: "Премьеры",
		Author: models.User{
			ID: 0,
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(true, nil)
	repository.EXPECT().GetCollection(ctx, inputParams).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetCollectionAuthorized_AuthorErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(false, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionAuthorized_IsPublicErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrCollectionNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(false, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionAuthorized_NotPublic(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrCollectionIsNotPublic

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(false, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionAuthorized_RepoErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputUser := &models.User{
		ID: 1,
	}
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserIsAuthor(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(true, nil)
	repository.EXPECT().GetCollection(ctx, inputParams).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionAuthorized(ctx, inputUser, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionNotAuthorized_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	output := models.Collection{
		ID:   1,
		Name: "Премьеры",
		Author: models.User{
			ID: 0,
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(true, nil)
	repository.EXPECT().GetCollection(ctx, inputParams).Return(output, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	actual, err := collectionService.GetCollectionNotAuthorized(ctx, inputParams)

	// Check success
	require.Nil(t, err, "Handling must be without errors")

	// Check result handling
	require.Equal(t, output, actual)
}

func TestCollectionService_GetCollectionNotAuthorized_IsPublicErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrCollectionNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(false, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionNotAuthorized(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionNotAuthorized_NotPublic(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrCollectionIsNotPublic

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(false, nil)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionNotAuthorized(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestCollectionService_GetCollectionNotAuthorized_RepoErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockCollectionRepository.NewMockRepository(ctrl)

	// Data
	inputParams := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckCollectionIsPublic(ctx, inputParams).Return(true, nil)
	repository.EXPECT().GetCollection(ctx, inputParams).Return(models.Collection{}, expectedErr)

	// Action
	collectionService := service.NewCollectionService(repository)

	_, actualErr := collectionService.GetCollectionNotAuthorized(ctx, inputParams)

	// Check success
	require.NotNil(t, actualErr, "Handling must be error")

	// Check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
