// Code generated by MockGen. DO NOT EDIT.
// Source: userservice.go

// Package mocUserService is a generated GoMock package.
package mocUserService

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	constparams "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ChangeUserProfileSettings mocks base method.
func (m *MockUserService) ChangeUserProfileSettings(ctx context.Context, user *models.User, params *constparams.ChangeUserSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserProfileSettings", ctx, user, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserProfileSettings indicates an expected call of ChangeUserProfileSettings.
func (mr *MockUserServiceMockRecorder) ChangeUserProfileSettings(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserProfileSettings", reflect.TypeOf((*MockUserService)(nil).ChangeUserProfileSettings), ctx, user, params)
}

// FilmRate mocks base method.
func (m *MockUserService) FilmRate(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRate", ctx, user, params)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRate indicates an expected call of FilmRate.
func (mr *MockUserServiceMockRecorder) FilmRate(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRate", reflect.TypeOf((*MockUserService)(nil).FilmRate), ctx, user, params)
}

// FilmRateDrop mocks base method.
func (m *MockUserService) FilmRateDrop(ctx context.Context, user *models.User, params *constparams.FilmRateDropParams) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRateDrop", ctx, user, params)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRateDrop indicates an expected call of FilmRateDrop.
func (mr *MockUserServiceMockRecorder) FilmRateDrop(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRateDrop", reflect.TypeOf((*MockUserService)(nil).FilmRateDrop), ctx, user, params)
}

// GetUserActivityOnFilm mocks base method.
func (m *MockUserService) GetUserActivityOnFilm(ctx context.Context, user *models.User, params *constparams.GetUserActivityOnFilmParams) (models.UserActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserActivityOnFilm", ctx, user, params)
	ret0, _ := ret[0].(models.UserActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserActivityOnFilm indicates an expected call of GetUserActivityOnFilm.
func (mr *MockUserServiceMockRecorder) GetUserActivityOnFilm(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserActivityOnFilm", reflect.TypeOf((*MockUserService)(nil).GetUserActivityOnFilm), ctx, user, params)
}

// GetUserCollections mocks base method.
func (m *MockUserService) GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCollections", ctx, user, params)
	ret0, _ := ret[0].([]models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCollections indicates an expected call of GetUserCollections.
func (mr *MockUserServiceMockRecorder) GetUserCollections(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCollections", reflect.TypeOf((*MockUserService)(nil).GetUserCollections), ctx, user, params)
}

// GetUserProfileByID mocks base method.
func (m *MockUserService) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileByID", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileByID indicates an expected call of GetUserProfileByID.
func (mr *MockUserServiceMockRecorder) GetUserProfileByID(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileByID", reflect.TypeOf((*MockUserService)(nil).GetUserProfileByID), ctx, user)
}

// GetUserProfileSettings mocks base method.
func (m *MockUserService) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileSettings", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileSettings indicates an expected call of GetUserProfileSettings.
func (mr *MockUserServiceMockRecorder) GetUserProfileSettings(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileSettings", reflect.TypeOf((*MockUserService)(nil).GetUserProfileSettings), ctx, user)
}

// NewFilmReview mocks base method.
func (m *MockUserService) NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *constparams.NewFilmReviewParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFilmReview", ctx, user, review, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewFilmReview indicates an expected call of NewFilmReview.
func (mr *MockUserServiceMockRecorder) NewFilmReview(ctx, user, review, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFilmReview", reflect.TypeOf((*MockUserService)(nil).NewFilmReview), ctx, user, review, params)
}