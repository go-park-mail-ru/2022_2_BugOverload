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

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	protoModels "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/models"
	proto "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/protobuf"
	"go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/server"
	mockWarehouseService "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/service/mocks"
)

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetRecommendation_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	outputService := models.Film{
		ID:   12,
		Name: "testname",
	}

	input := &proto.Nothing{}

	expected := protoModels.NewFilmProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	filmService.EXPECT().GetRecommendation(ctx).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetRecommendation(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetRecommendation_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	outputServiceErr := errors.ErrWorkDatabase

	input := &proto.Nothing{}

	expectedErr := status.Error(codes.Internal, errors.ErrWorkDatabase.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	filmService.EXPECT().GetRecommendation(ctx).Return(models.Film{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetRecommendation(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetFilmByID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputFilmService := &models.Film{
		ID: 1,
	}
	inputParamsService := &constparams.GetFilmParams{
		CountActors: 1,
		CountImages: 1,
	}

	outputService := models.Film{
		ID:          1,
		Name:        "testname",
		Description: "testdescription",
	}

	input := protoModels.NewGetFilmParamsProto(inputFilmService, inputParamsService)

	expected := protoModels.NewFilmProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	filmService.EXPECT().GetFilmByID(ctx, inputFilmService, inputParamsService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetFilmByID(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetFilmByID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputFilmService := &models.Film{
		ID: 0,
	}
	inputParamsService := &constparams.GetFilmParams{
		CountActors: 0,
		CountImages: 0,
	}

	outputServiceErr := errors.ErrFilmNotFound

	input := protoModels.NewGetFilmParamsProto(inputFilmService, inputParamsService)

	expectedErr := status.Error(codes.NotFound, errors.ErrFilmNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	filmService.EXPECT().GetFilmByID(ctx, inputFilmService, inputParamsService).Return(models.Film{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetFilmByID(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetReviewsByFilmID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetFilmReviewsParams{
		FilmID:       1,
		CountReviews: 1,
		Offset:       0,
	}

	outputService := []models.Review{
		{
			ID:   1,
			Body: "testreview",
		},
	}

	input := protoModels.NewGetFilmReviewsParamsProto(inputService)

	expected := protoModels.NewReviewsProto(outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	filmService.EXPECT().GetReviewsByFilmID(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetReviewsByFilmID(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetReviewsByFilmID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetFilmReviewsParams{
		FilmID:       0,
		CountReviews: 0,
		Offset:       0,
	}

	outputServiceErr := errors.ErrNotFoundInDB

	input := protoModels.NewGetFilmReviewsParamsProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrNotFoundInDB.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	filmService.EXPECT().GetReviewsByFilmID(ctx, inputService).Return([]models.Review{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetReviewsByFilmID(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetStdCollection_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetStdCollectionParams{
		Target:     "genre",
		Key:        "comedy",
		SortParam:  "rating",
		CountFilms: 10,
		Delimiter:  "10",
	}

	outputService := models.Collection{
		ID:          1,
		Name:        "GenresCollection",
		Description: "TestDescription",
	}

	input := protoModels.NewGetStdCollectionParamsProto(inputService)

	expected := protoModels.NewCollectionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	collectionService.EXPECT().GetStdCollection(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetStdCollection(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetStdCollection_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetStdCollectionParams{
		Target: "invalid_target",
	}

	outputServiceErr := errors.ErrNotFindSuchTarget

	input := protoModels.NewGetStdCollectionParamsProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrNotFindSuchTarget.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	collectionService.EXPECT().GetStdCollection(ctx, inputService).Return(models.Collection{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetStdCollection(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetPremieresCollection_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetPremiersCollectionParams{
		CountFilms: 10,
		Delimiter:  0,
	}

	outputService := models.Collection{
		ID:          1,
		Name:        "Премьеры",
		Description: "TestDescription",
	}

	input := protoModels.NewPremiersCollectionParamsProto(inputService)

	expected := protoModels.NewCollectionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	collectionService.EXPECT().GetPremieresCollection(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetPremieresCollection(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetPremieresCollection_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.GetPremiersCollectionParams{
		CountFilms: 10,
		Delimiter:  0,
	}

	outputServiceErr := errors.ErrNotFoundInDB

	input := protoModels.NewPremiersCollectionParamsProto(inputService)

	expectedErr := status.Error(codes.NotFound, errors.ErrNotFoundInDB.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	collectionService.EXPECT().GetPremieresCollection(ctx, inputService).Return(models.Collection{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetPremieresCollection(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetPersonByID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputPersonService := &models.Person{
		ID: 1,
	}
	inputParamsService := &constparams.GetPersonParams{
		CountFilms:  10,
		CountImages: 5,
	}

	outputService := models.Person{
		ID:   1,
		Name: "TestName",
	}

	input := protoModels.NewGetPersonParamsProto(inputPersonService, inputParamsService)

	expected := protoModels.NewPersonProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	personService.EXPECT().GetPersonByID(ctx, inputPersonService, inputParamsService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetPersonByID(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetPersonByID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputPersonService := &models.Person{
		ID: 0,
	}
	inputParamsService := &constparams.GetPersonParams{
		CountFilms:  0,
		CountImages: 0,
	}

	outputServiceErr := errors.ErrPersonNotFound

	input := protoModels.NewGetPersonParamsProto(inputPersonService, inputParamsService)

	expectedErr := status.Error(codes.NotFound, errors.ErrPersonNotFound.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	personService.EXPECT().GetPersonByID(ctx, inputPersonService, inputParamsService).Return(models.Person{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetPersonByID(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetCollectionFilmsAuthorized_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputParamsService := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	outputService := models.Collection{
		ID:          1,
		Name:        "TestCollection",
		Description: "TestDescription",
	}

	input := protoModels.NewCollectionGetFilmsAuthParamsProto(inputUserService, inputParamsService)

	expected := protoModels.NewCollectionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	collectionService.EXPECT().GetCollectionAuthorized(ctx, inputUserService, inputParamsService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetCollectionFilmsAuthorized(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetCollectionFilmsAuthorized_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputUserService := &models.User{
		ID: 1,
	}
	inputParamsService := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 12,
		SortParam:    "rating",
	}

	outputServiceErr := errors.ErrCollectionIsNotPublic

	input := protoModels.NewCollectionGetFilmsAuthParamsProto(inputUserService, inputParamsService)

	expectedErr := status.Error(codes.PermissionDenied, errors.ErrCollectionIsNotPublic.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	collectionService.EXPECT().GetCollectionAuthorized(ctx, inputUserService, inputParamsService).Return(models.Collection{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetCollectionFilmsAuthorized(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_GetCollectionFilmsNotAuthorized_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputParamsService := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 1,
		SortParam:    "rating",
	}

	outputService := models.Collection{
		ID:          1,
		Name:        "TestCollection",
		Description: "TestDescription",
	}

	input := protoModels.NewCollectionGetFilmsNotAuthParamsProto(inputParamsService)

	expected := protoModels.NewCollectionProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	collectionService.EXPECT().GetCollectionNotAuthorized(ctx, inputParamsService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.GetCollectionFilmsNotAuthorized(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_GetCollectionFilmsNotAuthorized_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputParamsService := &constparams.CollectionGetFilmsRequestParams{
		CollectionID: 12,
		SortParam:    "rating",
	}

	outputServiceErr := errors.ErrCollectionIsNotPublic

	input := protoModels.NewCollectionGetFilmsNotAuthParamsProto(inputParamsService)

	expectedErr := status.Error(codes.PermissionDenied, errors.ErrCollectionIsNotPublic.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	collectionService.EXPECT().GetCollectionNotAuthorized(ctx, inputParamsService).Return(models.Collection{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.GetCollectionFilmsNotAuthorized(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}

// Handler OK, NOT OK workflow in tests
func TestWarehouseServiceGRPCServer_Search_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "correct_query",
	}

	outputService := models.Search{
		Films:   []models.Film{},
		Serials: []models.Film{},
		Persons: []models.Person{},
	}

	input := protoModels.NewSearchParamsProto(inputService)

	expected := protoModels.NewSearchResponseProto(&outputService)

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	searchService.EXPECT().Search(ctx, inputService).Return(outputService, nil)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	actual, err := warehouseServer.Search(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, expected, actual)
}

func TestWarehouseServiceGRPCServer_Search_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseService.NewMockCollectionService(ctrl)
	filmService := mockWarehouseService.NewMockFilmService(ctrl)
	personService := mockWarehouseService.NewMockPersonService(ctrl)
	searchService := mockWarehouseService.NewMockSearchService(ctrl)

	// Data
	inputService := &constparams.SearchParams{
		Query: "incorrect_query",
	}

	outputServiceErr := errors.ErrWorkDatabase

	input := protoModels.NewSearchParamsProto(inputService)

	expectedErr := status.Error(codes.Internal, errors.ErrWorkDatabase.Error())

	// Create required setup for handling
	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(context.TODO(), constparams.LoggerKey, logger)

	requestID := uuid.NewV4().String()

	ctx = context.WithValue(ctx, constparams.RequestIDKey, requestID)

	// Settings mock
	searchService.EXPECT().Search(ctx, inputService).Return(models.Search{}, outputServiceErr)

	// Init
	grpcServer := grpc.NewServer()

	warehouseServer := server.NewWarehouseServiceGRPCServer(grpcServer, collectionService, filmService, personService, searchService)

	// Action
	_, actualErr := warehouseServer.Search(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, expectedErr, actualErr)
}
