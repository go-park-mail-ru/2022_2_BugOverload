// Code generated by MockGen. DO NOT EDIT.
// Source: userpgx.go

// Package mockUserRepository is a generated GoMock package.
package mockUserRepository

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	constparams "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddFilmToCollection mocks base method.
func (m *MockUserRepository) AddFilmToCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilmToCollection", ctx, user, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilmToCollection indicates an expected call of AddFilmToCollection.
func (mr *MockUserRepositoryMockRecorder) AddFilmToCollection(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilmToCollection", reflect.TypeOf((*MockUserRepository)(nil).AddFilmToCollection), ctx, user, params)
}

// ChangeUserProfileNickname mocks base method.
func (m *MockUserRepository) ChangeUserProfileNickname(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserProfileNickname", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserProfileNickname indicates an expected call of ChangeUserProfileNickname.
func (mr *MockUserRepositoryMockRecorder) ChangeUserProfileNickname(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserProfileNickname", reflect.TypeOf((*MockUserRepository)(nil).ChangeUserProfileNickname), ctx, user)
}

// CheckExistFilmInCollection mocks base method.
func (m *MockUserRepository) CheckExistFilmInCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistFilmInCollection", ctx, user, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistFilmInCollection indicates an expected call of CheckExistFilmInCollection.
func (mr *MockUserRepositoryMockRecorder) CheckExistFilmInCollection(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistFilmInCollection", reflect.TypeOf((*MockUserRepository)(nil).CheckExistFilmInCollection), ctx, user, params)
}

// CheckUserAccessToUpdateCollection mocks base method.
func (m *MockUserRepository) CheckUserAccessToUpdateCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAccessToUpdateCollection", ctx, user, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserAccessToUpdateCollection indicates an expected call of CheckUserAccessToUpdateCollection.
func (mr *MockUserRepositoryMockRecorder) CheckUserAccessToUpdateCollection(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAccessToUpdateCollection", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAccessToUpdateCollection), ctx, user, params)
}

// DropFilmFromCollection mocks base method.
func (m *MockUserRepository) DropFilmFromCollection(ctx context.Context, user *models.User, params *constparams.UserCollectionFilmsUpdateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropFilmFromCollection", ctx, user, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropFilmFromCollection indicates an expected call of DropFilmFromCollection.
func (mr *MockUserRepositoryMockRecorder) DropFilmFromCollection(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropFilmFromCollection", reflect.TypeOf((*MockUserRepository)(nil).DropFilmFromCollection), ctx, user, params)
}

// FilmRateDrop mocks base method.
func (m *MockUserRepository) FilmRateDrop(ctx context.Context, user *models.User, params *constparams.FilmRateDropParams) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRateDrop", ctx, user, params)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRateDrop indicates an expected call of FilmRateDrop.
func (mr *MockUserRepositoryMockRecorder) FilmRateDrop(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRateDrop", reflect.TypeOf((*MockUserRepository)(nil).FilmRateDrop), ctx, user, params)
}

// FilmRateSet mocks base method.
func (m *MockUserRepository) FilmRateSet(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRateSet", ctx, user, params)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRateSet indicates an expected call of FilmRateSet.
func (mr *MockUserRepositoryMockRecorder) FilmRateSet(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRateSet", reflect.TypeOf((*MockUserRepository)(nil).FilmRateSet), ctx, user, params)
}

// FilmRateUpdate mocks base method.
func (m *MockUserRepository) FilmRateUpdate(ctx context.Context, user *models.User, params *constparams.FilmRateParams) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRateUpdate", ctx, user, params)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRateUpdate indicates an expected call of FilmRateUpdate.
func (mr *MockUserRepositoryMockRecorder) FilmRateUpdate(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRateUpdate", reflect.TypeOf((*MockUserRepository)(nil).FilmRateUpdate), ctx, user, params)
}

// FilmRatingExist mocks base method.
func (m *MockUserRepository) FilmRatingExist(ctx context.Context, user *models.User, filmID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmRatingExist", ctx, user, filmID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilmRatingExist indicates an expected call of FilmRatingExist.
func (mr *MockUserRepositoryMockRecorder) FilmRatingExist(ctx, user, filmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmRatingExist", reflect.TypeOf((*MockUserRepository)(nil).FilmRatingExist), ctx, user, filmID)
}

// GetUserActivityOnFilm mocks base method.
func (m *MockUserRepository) GetUserActivityOnFilm(ctx context.Context, user *models.User, params *constparams.GetUserActivityOnFilmParams) (models.UserActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserActivityOnFilm", ctx, user, params)
	ret0, _ := ret[0].(models.UserActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserActivityOnFilm indicates an expected call of GetUserActivityOnFilm.
func (mr *MockUserRepositoryMockRecorder) GetUserActivityOnFilm(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserActivityOnFilm", reflect.TypeOf((*MockUserRepository)(nil).GetUserActivityOnFilm), ctx, user, params)
}

// GetUserCollections mocks base method.
func (m *MockUserRepository) GetUserCollections(ctx context.Context, user *models.User, params *constparams.GetUserCollectionsParams) ([]models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCollections", ctx, user, params)
	ret0, _ := ret[0].([]models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCollections indicates an expected call of GetUserCollections.
func (mr *MockUserRepositoryMockRecorder) GetUserCollections(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCollections", reflect.TypeOf((*MockUserRepository)(nil).GetUserCollections), ctx, user, params)
}

// GetUserProfileByID mocks base method.
func (m *MockUserRepository) GetUserProfileByID(ctx context.Context, user *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileByID", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileByID indicates an expected call of GetUserProfileByID.
func (mr *MockUserRepositoryMockRecorder) GetUserProfileByID(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserProfileByID), ctx, user)
}

// GetUserProfileSettings mocks base method.
func (m *MockUserRepository) GetUserProfileSettings(ctx context.Context, user *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileSettings", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileSettings indicates an expected call of GetUserProfileSettings.
func (mr *MockUserRepositoryMockRecorder) GetUserProfileSettings(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileSettings", reflect.TypeOf((*MockUserRepository)(nil).GetUserProfileSettings), ctx, user)
}

// NewFilmReview mocks base method.
func (m *MockUserRepository) NewFilmReview(ctx context.Context, user *models.User, review *models.Review, params *constparams.NewFilmReviewParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFilmReview", ctx, user, review, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewFilmReview indicates an expected call of NewFilmReview.
func (mr *MockUserRepositoryMockRecorder) NewFilmReview(ctx, user, review, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFilmReview", reflect.TypeOf((*MockUserRepository)(nil).NewFilmReview), ctx, user, review, params)
}
