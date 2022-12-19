package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	mockUserRepository "go-park-mail-ru/2022_2_BugOverload/internal/user/repository/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/service"
)

func TestUserService_GetUserProfileByID_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

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
	repository.EXPECT().GetUserProfileByID(ctx, input).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.GetUserProfileByID(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_GetUserProfileByID_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserProfileByID(ctx, input).Return(models.User{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.GetUserProfileByID(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_GetUserProfileSettings_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

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
	repository.EXPECT().GetUserProfileSettings(ctx, input).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.GetUserProfileSettings(ctx, input)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_GetUserProfileSettings_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	input := &models.User{
		ID: 0,
	}

	expectedErr := errors.ErrUserNotFound

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserProfileSettings(ctx, input).Return(models.User{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.GetUserProfileSettings(ctx, input)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_ChangeUserProfileSettings_Nickname_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID:       12,
		Nickname: "testnewnickname",
	}

	inputParams := &constparams.ChangeUserSettings{
		Nickname: "testnewnickname",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().ChangeUserProfileNickname(ctx, inputUser).Return(nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.ChangeUserProfileSettings(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, actualErr, "Handling must be without errors")
}

func TestUserService_ChangeUserProfileSettings_Password_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID:       12,
		Nickname: "testnewnickname",
	}

	inputParams := &constparams.ChangeUserSettings{
		CurPassword: "123",
		NewPassword: "12345",
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().UpdatePassword(ctx, inputUser, inputParams.CurPassword, inputParams.NewPassword).Return(nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.ChangeUserProfileSettings(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, actualErr, "Handling must be without errors")
}

func TestUserService_ChangeUserProfileSettings_Nickname_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID:       12,
		Nickname: "testnewnickname",
	}

	inputParams := &constparams.ChangeUserSettings{
		Nickname: "testnewnickname",
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().ChangeUserProfileNickname(ctx, inputUser).Return(expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.ChangeUserProfileSettings(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_ChangeUserProfileSettings_Password_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID:       12,
		Nickname: "testnewnickname",
	}

	inputParams := &constparams.ChangeUserSettings{
		CurPassword: "123",
		NewPassword: "12345",
	}

	expectedErr := errors.ErrIncorrectPassword

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	authService.EXPECT().UpdatePassword(ctx, inputUser, inputParams.CurPassword, inputParams.NewPassword).Return(expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.ChangeUserProfileSettings(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRateUpdate_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  7,
	}

	output := models.Film{
		Rating: 7.12,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(true, nil)
	repository.EXPECT().FilmRateUpdate(ctx, inputUser, inputParams).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_FilmRateSet_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  7,
	}

	output := models.Film{
		Rating: 7.12,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(false, nil)
	repository.EXPECT().FilmRateSet(ctx, inputUser, inputParams).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_FilmRate_Score_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  13,
	}

	expectedErr := errors.ErrBadRequestParams

	// Create required setup for handling
	ctx := context.TODO()

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRate_RepoExistErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  3,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(false, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRate_RepoSetErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  3,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(false, nil)
	repository.EXPECT().FilmRateSet(ctx, inputUser, inputParams).Return(models.Film{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRate_RepoUpdateErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateParams{
		FilmID: 1,
		Score:  3,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(true, nil)
	repository.EXPECT().FilmRateUpdate(ctx, inputUser, inputParams).Return(models.Film{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRate(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRateDrop_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateDropParams{
		FilmID: 1,
	}

	output := models.Film{
		Rating: 7.12,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(true, nil)
	repository.EXPECT().FilmRateDrop(ctx, inputUser, inputParams).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.FilmRateDrop(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_FilmRateDrop_RepoExistErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateDropParams{
		FilmID: 1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(false, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRateDrop(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRateDrop_NotExist(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateDropParams{
		FilmID: 1,
	}

	expectedErr := errors.ErrFilmRatingNotExist

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(false, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRateDrop(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_FilmRateDrop_RepoDropErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.FilmRateDropParams{
		FilmID: 1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().FilmRatingExist(ctx, inputUser, inputParams.FilmID).Return(true, nil)
	repository.EXPECT().FilmRateDrop(ctx, inputUser, inputParams).Return(models.Film{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.FilmRateDrop(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_NewFilmReview_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.NewFilmReviewParams{
		FilmID: 1,
	}

	inputReview := &models.Review{
		ID: 1,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().NewFilmReview(ctx, inputUser, inputReview, inputParams).Return(nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.NewFilmReview(ctx, inputUser, inputReview, inputParams)

	// check success
	require.Nil(t, actualErr, "Handling must be without errors")
}

func TestUserService_NewFilmReview_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.NewFilmReviewParams{
		FilmID: 1,
	}

	inputReview := &models.Review{
		ID: 1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().NewFilmReview(ctx, inputUser, inputReview, inputParams).Return(expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.NewFilmReview(ctx, inputUser, inputReview, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_GetUserActivityOnFilm_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.GetUserActivityOnFilmParams{
		FilmID: 1,
	}

	output := models.UserActivity{
		CountReviews: 1,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserActivityOnFilm(ctx, inputUser, inputParams).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.GetUserActivityOnFilm(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_GetUserActivityOnFilm_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.GetUserActivityOnFilmParams{
		FilmID: 1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserActivityOnFilm(ctx, inputUser, inputParams).Return(models.UserActivity{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.GetUserActivityOnFilm(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_GetUserCollections_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.GetUserCollectionsParams{
		SortParam:        "rating",
		CountCollections: 10,
		Delimiter:        "10",
	}

	output := []models.Collection{
		{
			ID:   1,
			Name: "testcollection",
		},
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserCollections(ctx, inputUser, inputParams).Return(output, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actual, err := userService.GetUserCollections(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, err, "Handling must be without errors")

	// check result handling
	require.Equal(t, output, actual)
}

func TestUserService_GetUserCollections_NOT_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.GetUserCollectionsParams{
		SortParam:        "unknown",
		CountCollections: 10,
		Delimiter:        "10",
	}

	expectedErr := errors.ErrUnsupportedSortParameter

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().GetUserCollections(ctx, inputUser, inputParams).Return([]models.Collection{}, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	_, actualErr := userService.GetUserCollections(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_AddFilmToUserCollection_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().AddFilmToCollection(ctx, inputUser, inputParams).Return(nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, actualErr, "Handling must be without errors")
}

func TestUserService_AddFilmToUserCollection_AuthorErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_AddFilmToUserCollection_NotAuthor(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrBadUserCollectionID

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(false, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_AddFilmToUserCollection_ExistErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(false, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_AddFilmToUserCollection_Exist(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrFilmExistInCollection

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(true, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_AddFilmToUserCollection_AddFilmErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(false, nil)
	repository.EXPECT().AddFilmToCollection(ctx, inputUser, inputParams).Return(expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.AddFilmToUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_DropFilmFromUserCollection_OK(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().DropFilmFromCollection(ctx, inputUser, inputParams).Return(nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.Nil(t, actualErr, "Handling must be without errors")
}

func TestUserService_DropFilmFromUserCollection_AuthorErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_DropFilmFromUserCollection_NotAuthor(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrBadUserCollectionID

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(false, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_DropFilmFromUserCollection_ExistErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(false, expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_DropFilmFromUserCollection_NotExist(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrFilmNotExistInCollection

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(false, nil)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}

func TestUserService_DropFilmFromUserCollection_DropFilmErr(t *testing.T) {
	// Work with mocks
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockUserRepository.NewMockUserRepository(ctrl)
	authService := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	inputUser := &models.User{
		ID: 12,
	}

	inputParams := &constparams.UserCollectionFilmsUpdateParams{
		CollectionID: 1,
		FilmID:       1,
	}

	expectedErr := errors.ErrWorkDatabase

	// Create required setup for handling
	ctx := context.TODO()

	// Settings mock
	repository.EXPECT().CheckUserAccessToUpdateCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().CheckExistFilmInCollection(ctx, inputUser, inputParams).Return(true, nil)
	repository.EXPECT().DropFilmFromCollection(ctx, inputUser, inputParams).Return(expectedErr)

	// Action
	userService := service.NewUserProfileService(repository, authService)

	actualErr := userService.DropFilmFromUserCollection(ctx, inputUser, inputParams)

	// check success
	require.NotNil(t, actualErr, "Handling must be error")

	// check result handling
	require.Equal(t, stdErrors.Cause(expectedErr), stdErrors.Cause(actualErr))
}
